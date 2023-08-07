package purchase

import (
	"context"
	"database/sql"

	db "nam_0508/internal/repo/dbmodel"
)

type Repository struct {
	DatabaseConn *sql.DB
	Query        *db.Queries
}

func NewMysqlRepository(databaseConn *sql.DB) *Repository {
	query := db.New(databaseConn)
	return &Repository{
		Query:        query,
		DatabaseConn: databaseConn,
	}
}

func (r *Repository) CreatePurchase(ctx context.Context, tx *sql.Tx, purchase db.CreatePurchaseParams) (*db.Purchase, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	err := query.CreatePurchase(ctx, purchase)
	if err != nil {
		return nil, err
	}

	id, err := query.LastInsertID(ctx)
	if err != nil {
		return nil, err
	}

	purchaserDB, err := query.GetPurchase(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &purchaserDB, nil
}
