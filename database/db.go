package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/joho/godotenv"
	"os"
)

var DbConn *sql.DB

func getDns() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("dbUserName"),
		os.Getenv("dbPassword"),
		os.Getenv("dbUrl"),
		os.Getenv("dbPort"),
		os.Getenv("dbName"))
}

func SetupDatabase()  {
	var err error
	DbConn, err = sql.Open("mysql",getDns())
	if err != nil{
		log.Fatalln(err)
	}
}
