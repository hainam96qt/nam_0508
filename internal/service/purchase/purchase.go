package purchase

import (
	"database/sql"

	"nam_0508/internal/config"
	db "nam_0508/internal/repo/dbmodel"
)

type (
	Service struct {
		conf *configs.Config

		DatabaseConn *sql.DB
		Query        *db.Queries
	}
)

func NewPurchaseService(conf *configs.Config, DatabaseConn *sql.DB) *Service {
	query := db.New(DatabaseConn)
	return &Service{
		conf:         conf,
		DatabaseConn: DatabaseConn,
		Query:        query,
	}
}
