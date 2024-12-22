package user

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id          int
	Username    string
	Password    string
	Token       string
	Created_at  string
	Created_by  string
	Modified_at string
	Modified_by string
}

type Claims struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	jwt.RegisteredClaims
}

func (u *User) Scan(rows *sql.Rows) error {
	return rows.Scan(&u.Id, &u.Username, &u.Password, &u.Created_at, &u.Created_by, &u.Modified_at, &u.Modified_by)
}
