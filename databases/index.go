package databases

import (
	"database/sql"
	"fmt"
	"quiz/config"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	var port int
	port, err = strconv.Atoi(config.VarConfig.DB_PORT)
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.VarConfig.DB_HOST, port, config.VarConfig.DB_USERNAME, config.VarConfig.DB_PASSWORD, config.VarConfig.DB_NAME)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
