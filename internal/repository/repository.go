package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var DB = getDB()

func getDB() *sql.DB {
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

func getAllSyntax(tableName string, columns string, wheres string) string {
	return generateSyntax("SELECT %s FROM `%s` %s", columns, tableName, wheres)
}

func getFirstSyntax(tableName string, columns string, wheres string) string {
	return generateSyntax("SELECT %s FROM `%s` %s LIMIT 1", columns, tableName, wheres)
}

func insertSyntax(tableName string, columns string, values string) string {
	return generateSyntax("INSERT INTO `%s` (%s) VALUES %v", tableName, columns, values)
}

func getLastInsertID(tableName string) string {
	return generateSyntax("SELECT LAST_INSERT_ID() FROM %s", tableName)
}

func debug(result interface{}) {
	if os.Getenv("DB_DEBUG") == "true" {
		log.Printf("repository debug: %v", result)
	}
}
