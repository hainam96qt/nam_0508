package wager

import (
	"context"

	db "nam_0508/internal/repo/dbmodel"
)

type (
	Service struct {
		wagerRepo WagerRepository
	}

	WagerRepository interface {
		ListWagers(ctx context.Context, params db.ListWagersParams) ([]db.Wager, error)
		CreateWager(ctx context.Context, wager db.CreateWagerParams) (*db.Wager, error)
	}
)

func NewWagerService(wagerRepo WagerRepository) *Service {
	return &Service{
		wagerRepo: wagerRepo,
	}
}
