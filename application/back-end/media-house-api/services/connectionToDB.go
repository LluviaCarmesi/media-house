package services

import (
	"back-end/secrets"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	connection, err := sql.Open("mysql", secrets.CONNECTION_STRING)
	if err != nil {
		fmt.Println("Couldn't connect")
		panic(err)
	}

	connection.SetConnMaxLifetime(60)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)

	return connection
}
