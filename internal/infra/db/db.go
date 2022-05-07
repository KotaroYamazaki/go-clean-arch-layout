package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

type DBCon struct {
	DB *sql.DB
}

func NewDBCon(dsn string) (*DBCon, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(3 * time.Minute)

	boil.SetLocation(time.Local)
	boil.SetDB(db)

	return &DBCon{
		DB: db,
	}, nil
}
