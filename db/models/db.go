package db

import (
	"context"
	"database/sql"
)

type Queries struct {
	db DB
}

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New creates a new Queries object given a DB
func New(db DB) *Queries {
	return &Queries{db: db}
}
