package clients

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"microservices/ticket/src/utils"
)

var (
	DbUser     = utils.GoDotEnvVariable("DB_USER")
	DbPassword = utils.GoDotEnvVariable("DB_PASSWORD")
	DbName     = utils.GoDotEnvVariable("DB_NAME")
	DbPort     = utils.GoDotEnvVariable("DB_PORT")
)

func GetDBInstance() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s host=ticket-db port=%s dbname=%s sslmode=disable",
		DbUser,
		DbPassword,
		DbPort,
		DbName,
	)
	//dbinfo := fmt.Sprintf("postgresql://%s:%s@ticket-db:%s/%s?schema=public",
	//	DbUser,
	//	DbPassword,
	//	DbPort,
	//	DbName,
	//)

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	}

	return db
}
