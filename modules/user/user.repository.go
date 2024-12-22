package user

import (
	"errors"
	"quiz/config"
	"quiz/databases"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func GetAll() (result []User, err error) {
	getAllQuery := `SELECT * FROM users`
	rows, err := databases.DB.Query(getAllQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Username, &user.Created_at, &user.Created_by, &user.Modified_at, &user.Modified_by)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

func GetOne(id int) (result User, err error) {
	getOneQuery := `SELECT * FROM users WHERE id=$1`
	err = databases.DB.QueryRow(getOneQuery, id).Scan(&result.Id, &result.Username, &result.Created_at, &result.Created_by, &result.Modified_at, &result.Modified_by)

	if err != nil {
		return result, errors.New("data not found")
	}

	return result, nil
}

func Create(user User) (err error) {
	createQuery := `INSERT INTO users (username,password,  created_by,  modified_by) VALUES ($1, $2, $3, $4)`
	_, err = databases.DB.Exec(createQuery, user.Username, user.Password, user.Created_by, user.Modified_by)

	if err != nil {
		return err
	}

	return nil
}

func Update(id int, user User) (err error) {
	_, err = GetOne(id)
	if err != nil {
		return err
	}

	updateQuery := `UPDATE users SET username=$1, password=$2, created_at=$3, created_by=$4, modified_at=$5, modified_by=$6 WHERE id=$7`
	_, err = databases.DB.Exec(updateQuery, user.Username, user.Password, user.Token, user.Created_at, user.Created_by, user.Modified_at, user.Modified_by, id)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id int) (err error) {
	_, err = GetOne(id)
	if err != nil {
		return err
	}
	deleteQuery := `DELETE FROM users WHERE id=$1`
	_, err = databases.DB.Exec(deleteQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func GetToken(username string, password string) (result User, status int, err error) {
	getOneQuery := `SELECT username,password,id FROM users WHERE username=$1`
	err = databases.DB.QueryRow(getOneQuery, username).Scan(&result.Username, &result.Password, &result.Id)

	if err != nil {
		return result, 404, errors.New("username or password is invalid")
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return result, 404, errors.New("username or password is invalid")
	}

	jwtClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
		"id":       result.Id,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := jwtClaim.SignedString([]byte(config.VarConfig.SecretJwt))

	if err != nil {
		return result, 500, err
	}

	insertToken := `UPDATE users SET token=array_append(token, $1) WHERE id=$2`
	_, err = databases.DB.Exec(insertToken, token, result.Id)

	if err != nil {
		return result, 500, err
	}

	result.Token = token

	return result, 200, nil
}

func RemoveToken(token string, id int) (status int, err error) {
	removeToken := `UPDATE users SET token=array_remove(token, $1) WHERE id=$2`
	_, err = databases.DB.Exec(removeToken, token, id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func CheckToken(token string, id int) bool {
	checkToken := `SELECT token FROM users WHERE id=$1`
	var tokens []string
	err := databases.DB.QueryRow(checkToken, id).Scan(pq.Array(&tokens))

	if err != nil {
		return false
	}
	for _, t := range tokens {
		if t == token {
			return true
		}
	}

	return false
}
