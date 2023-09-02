package data

import (
	"context"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type TokenStore interface {
	GetByToken(plainText string) (*Token, error)
	GetUserForToken(token Token) (*User, error)
	GenerateToken(userID int, ttl time.Duration) (*Token, error)
	AuthenticateToken(r *http.Request) (*User, error)
	InsertToken(token Token, u User) error
	DeleteByToken(plainText string) error
	ValidToken(plainText string) (bool, error)
	GetAllTokens() ([]Token, error)
}

type Token struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	TokenHash []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Expiry    time.Time `json:"expiry"`
}

// GetByToken retrieves a token from the database based on its plaintext value.
func (t *Token) GetByToken(plainText string) (*Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, user_id, email, token, token_hash,created_at,updated_at,expiry from tokens where token = $1`

	var token Token
	row := db.QueryRowContext(ctx, query, plainText)
	err := row.Scan(
		&token.ID,
		&token.UserID,
		&token.Email,
		&token.Token,
		&token.TokenHash,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.Expiry,
	)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// GetUserForToken retrieves a user associated with the provided token from the database.
func (t *Token) GetUserForToken(token Token) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := "SELECT id, email, first_name, last_name, password, created_at, updated_at FROM users WHERE id = $1"

	var user User
	row := db.QueryRowContext(ctx, query, token.UserID)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GenerateToken generates a new token for the given user ID with a specified time-to-live duration.
func (t *Token) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Token = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Token)) //this can make error
	token.TokenHash = hash[:]

	return token, nil
}

// AuthenticateToken authenticates a user based on the token provided in the HTTP request.
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("no authorization header received")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no valid authorization header received")
	}

	token := headerParts[1]

	if len(token) != 26 {
		return nil, errors.New("Token wrong size")
	}

	tkn, err := t.GetByToken(token)
	if err != nil {
		return nil, errors.New("no matching token found")
	}

	if tkn.Expiry.Before(time.Now()) {
		return nil, errors.New("Token expired")
	}

	user, err := t.GetUserForToken(*tkn)
	if err != nil {
		return nil, errors.New("no matching user found")
	}

	return user, nil
}

func (t *Token) InsertToken(token Token, u User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	//delete any exsisting tokens
	stmt := `delete from tokens where user_id = $1`
	_, err := db.ExecContext(ctx, stmt, token.UserID)
	if err != nil {
		return nil
	}

	token.Email = u.Email

	stmt = `insert into tokens (user_id, email, token, token_hash, created_at, updated_at, expiry)
	 values($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.ExecContext(ctx, stmt,
		token.UserID,
		token.Email,
		token.Token,
		token.TokenHash,
		time.Now(),
		time.Now(),
		token.Expiry,
	)

	if err != nil {
		return err
	}
	return nil
}

func (t *Token) DeleteByToken(plainText string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `delete from tokens where token = $1`
	_, err := db.ExecContext(ctx, stmt, plainText)
	if err != nil {
		return err
	}

	return nil
}

func (t *Token) ValidToken(plainText string) (bool, error) {
	if t == nil {
		return false, errors.New("token instance is nil")
	}

	// Check if the token exists in the database
	token, err := t.GetByToken(plainText)
	fmt.Println("token", token)
	if err != nil {
		return false, fmt.Errorf("error getting token: %v", err)
	}

	// Check if the token has expired
	if token.Expiry.Before(time.Now()) {
		return false, errors.New("expired token")
	}

	// Check if the associated user exists in the database
	_, err = t.GetUserForToken(*token)
	if err != nil {
		return false, fmt.Errorf("error getting user for token: %v", err)
	}

	// The token is valid
	return true, nil
}

// GetAllTokens retrieves all tokens from the "tokens" table.
func (t *Token) GetAllTokens() ([]Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, user_id, email, token, token_hash, created_at, updated_at, expiry FROM tokens`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []Token

	for rows.Next() {
		var token Token
		err := rows.Scan(
			&token.ID,
			&token.UserID,
			&token.Email,
			&token.Token,
			&token.TokenHash,
			&token.CreatedAt,
			&token.UpdatedAt,
			&token.Expiry,
		)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}
