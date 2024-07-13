package storage

import (
	"database/sql"

	"auth-service/internal/config"

	sq "github.com/Masterminds/squirrel"
)

type (
	AuthSt struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*AuthSt, error) {

	db, err := ConnectDB(*config)
	if err != nil {
		return nil, err
	}

	return &AuthSt{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}
