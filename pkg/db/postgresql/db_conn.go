package postgresql

import (
	"database/sql"
	"fmt"

	"avito-banners/config"

	_ "github.com/lib/pq"
)

func NewPsqlDB(cfg *config.Config) (*sql.DB, error) {
	// connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Pass,
		cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return nil, nil
}
