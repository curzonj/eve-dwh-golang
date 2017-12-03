package model

import (
	"database/sql"
)

type DB interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Get(dest interface{}, query string, args ...interface{}) error
	Begin() (*sql.Tx, error)
}

type DAO struct {
	DB
}
