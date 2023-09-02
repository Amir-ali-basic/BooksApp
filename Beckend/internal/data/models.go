package data

import "database/sql"

type Models struct {
	User  User
	Token Token
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User:  User{},
		Token: Token{},
	}
}
