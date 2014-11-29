package app

import (
	"database/sql"
)

type DbProvider struct {
	Db *sql.DB
	GetSql func(name string) (string, error)
}

type AppContext struct {
	DbProvider *DbProvider
}

