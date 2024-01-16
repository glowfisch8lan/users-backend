package repository

import (
	"database/sql"
	"idapp/users/cmd/users/types"
)

var db *sql.DB

func CreateUser() any {
	user := types.User{
		ID:    1,
		Login: "hello",
	}

	return user
}
