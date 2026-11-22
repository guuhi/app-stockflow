package repo

import "github.com/jmoiron/sqlx"

type ExampleRepository struct {
	db *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) *ExampleRepository {
	return &ExampleRepository{db: db}
}