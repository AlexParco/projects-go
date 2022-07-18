package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alexparco/rest-api-todo/config"
	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	*sql.DB
}

func NewDbClient(config *config.Mysql) *DBClient {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBname,
	))

	if err != nil {
		log.Fatalf("Error connecting database %e", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting database %e", err)
	}

	return &DBClient{db}
}
