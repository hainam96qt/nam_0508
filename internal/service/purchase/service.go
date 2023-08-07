package purchase

import (
	"context"
	"database/sql"

	db "nam_0508/internal/repo/dbmodel"
)

type (
	Service struct {
		DatabaseConn *sql.DB

		purchaseRepo PurchaseRepository
		wagerRepo    WagerRepository
	}

	PurchaseRepository interface {
		CreatePurchase(ctx context.Context, tx *sql.Tx, purchase db.CreatePurchaseParams) (*db.Purchase, error)
	}

	WagerRepository interface {
		UpdatePurchaseWager(ctx context.Context, tx *sql.Tx, updateWagerParams db.UpdatePurchaseWagerParams) error
		GetWagerForUpdate(ctx context.Context, tx *sql.Tx, wagerID int32) (*db.Wager, error)
	}
)

func NewPurchaseService(DatabaseConn *sql.DB, purchaseRepo PurchaseRepository, wagerRepo WagerRepository) *Service {
	return &Service{
		DatabaseConn: DatabaseConn,
		purchaseRepo: purchaseRepo,
		wagerRepo:    wagerRepo,
	}
}
