package db

import "github.com/jmoiron/sqlx"

type Storage struct{
	db *sqlx.DB
	logger Logger
}
type Logger interface {
}

func NewStorage ()