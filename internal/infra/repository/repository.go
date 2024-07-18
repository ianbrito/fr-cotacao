package repository

import (
	"database/sql"
	"github.com/ianbrito/fr-cotacao/internal/infra/db"
)

var DB *sql.DB

func InitDB() {
	DB = db.GetConnection()
}

func CloseDB() error {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
