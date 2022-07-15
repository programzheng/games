package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func GetDB() *sql.DB {
	conn := os.Getenv("MYSQL_DBSTRING")

	DB, err := sql.Open("mysql", conn)
	if err != nil {
		panic(fmt.Sprintf("connection to mysql failed:%v", err))
	}
	return DB
}

func generateSyntax(syntax string, args ...any) string {
	s := fmt.Sprintf(syntax, args...)

	debug(s)
	return s
}

func debug(result interface{}) {
	if os.Getenv("DEBUG") == "true" {
		log.Printf("repository debug: %v", result)
	}
}
