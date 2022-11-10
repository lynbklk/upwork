package mysql

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config struct {
	DBMySQLHost     string
	DBMySQLPort     int
	DBMySQLUser     string
	DBMySQLPassword string
	DBMySQLDatabase string
}

func New(ctx context.Context, conf *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.DBMySQLUser,
		conf.DBMySQLPassword,
		conf.DBMySQLHost,
		conf.DBMySQLPort,
		conf.DBMySQLDatabase))
	if err != nil {
		return nil, err
	}
	setupDB(db)
	return db, nil
}

func setupDB(db *sqlx.DB) {
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(10 * time.Minute)
}
