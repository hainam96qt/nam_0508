package wager

import (
	"context"
	"database/sql"

	db "nam_0508/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		wagerRepo WagerRepository
	}

	WagerRepository interface {
		ListWagers(ctx context.Context, params db.ListWagersParams) ([]db.Wager, error)
		CreateWager(ctx context.Context, wager db.CreateWagerParams) (*db.Wager, error)
	}
)

func NewWagerService(DatabaseConn *sql.DB, wagerRepo WagerRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		wagerRepo:    wagerRepo,
	}
}
