package postgresql

import (
	"database/sql"
	"fmt"

	"avito-banners/config"
)

func NewPsqlDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Name,
		cfg.Postgres.User, cfg.Postgres.Pass,
	)

	db, err := sql.Open("postgre", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return nil, nil
}
