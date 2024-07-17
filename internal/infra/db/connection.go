package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance *sql.DB
	once       sync.Once
	closeOnce  sync.Once
)

func GetConnection() *sql.DB {
	once.Do(func() {
		var (
			host     = os.Getenv("DB_HOST")
			port     = os.Getenv("DB_PORT")
			user     = os.Getenv("DB_USERNAME")
			password = os.Getenv("DB_PASSWORD")
			database = os.Getenv("DB_DATABASE")
		)

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

		conn, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err := conn.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		fmt.Println("Connected to database")
		dbInstance = conn
	})

	return dbInstance
}

func CloseConnection() {
	closeOnce.Do(func() {
		if dbInstance != nil {
			err := dbInstance.Close()
			if err != nil {
				log.Fatalf("Failed to close database connection: %v", err)
			}
			fmt.Println("Database connection closed")
		}
	})
}
