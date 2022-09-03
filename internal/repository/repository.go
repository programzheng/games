package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DB = getDB()

func getDB() *sql.DB {
	driver := os.Getenv("DB_DRIVER")
	connString := os.Getenv("DB_CONNECTION_STRING")

	switch driver {
	case "mysql":
		DB, err := sql.Open("mysql", connString)
		if err != nil {
			panic(fmt.Sprintf("connection to mysql failed:%v", err))
		}
		return DB
	case "postgres":
		DB, err := sql.Open("postgres", connString)
		if err != nil {
			panic(fmt.Sprintf("connection to postgres failed:%v", err))
		}
		return DB
	}

	return nil
}

func generateSyntax(syntax string, args ...any) string {
	s := fmt.Sprintf(syntax, args...)

	debug(s)
	return s
}

func getAllSyntax(tableName string, columns string, wheres string) string {
	return generateSyntax("SELECT %s FROM `%s` %s", columns, tableName, wheres)
}

func getAllWithJoinSyntax(tableName string, columns string, wheres string, join string) string {
	return generateSyntax("SELECT %s FROM `%s` %s %s", columns, tableName, join, wheres)
}

func getFirstSyntax(tableName string, columns string, wheres string) string {
	return generateSyntax("SELECT %s FROM `%s` %s LIMIT 1", columns, tableName, wheres)
}

func insertSyntax(tableName string, columns string, values string) string {
	return generateSyntax("INSERT INTO `%s` (%s) VALUES %v", tableName, columns, values)
}

func updateSyntax(tableName string, updates string, wheres string) string {
	return generateSyntax("UPDATE `%s` SET %s %s", tableName, updates, wheres)
}

func getLastInsertID(tableName string) string {
	return generateSyntax("SELECT LAST_INSERT_ID() FROM %s", tableName)
}

func debug(result interface{}) {
	if os.Getenv("DB_DEBUG") == "true" {
		log.Printf("repository debug: %v", result)
	}
}
