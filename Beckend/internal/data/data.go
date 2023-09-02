// data.go
package data

import (
	"net/http"
	"time"
)

type Database interface {
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) (int, error)
	UpdateUser(user User) error
	DeleteUser(id int) error
	GetAllUsers() ([]*User, error)

	GetTokenByToken(token string) (*Token, error)
	GetUserForToken(token Token) (*User, error)
	GenerateToken(userID int, ttl time.Duration) (*Token, error)
	AuthenticateToken(r *http.Request) (*User, error)
	InsertToken(token Token, u User) error
	DeleteTokenByToken(token string) error
	ValidToken(token string) (bool, error)
	GetAllTokens() ([]Token, error)
}
