package authentication

import (
	"database/sql"

	configs "nam_0508/config"
	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
)

type (
	Service struct {
		conf *configs.Config

		DatabaseConn *sql.DB
		Query        *db.Queries

		jwtSvc jwtService
	}

	jwtService interface {
		GenerateTokenPair(userID int, userName string) (*model.TokenPair, error)
	}
)

func NewAuthenticationService(conf *configs.Config, DatabaseConn *sql.DB, jwtSvc jwtService) *Service {
	query := db.New(DatabaseConn)
	return &Service{
		conf:         conf,
		DatabaseConn: DatabaseConn,
		Query:        query,
		jwtSvc:       jwtSvc,
	}
}
