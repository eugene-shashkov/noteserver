package app

import (
	"database/sql"
	"log"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Db return connection to the mysql
func Db() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionParams := os.Getenv("DB_USER_NAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", connectionParams)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
