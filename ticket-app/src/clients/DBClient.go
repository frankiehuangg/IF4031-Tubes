package clients

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"microservices/ticket/src/utils"
)

var (
	DbUser     = utils.GoDotEnvVariable("TICKET_POSTGRES_USER")
	DbPassword = utils.GoDotEnvVariable("TICKET_POSTGRES_PASSWORD")
	DbName     = utils.GoDotEnvVariable("TICKET_POSTGRES_DB")
	DbPort     = utils.GoDotEnvVariable("TICKET_POSTGRES_PORT")
	DbHost     = utils.GoDotEnvVariable("TICKET_POSTGRES_HOST")
)

func GetDBInstance() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		DbUser,
		DbPassword,
		DbHost,
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
