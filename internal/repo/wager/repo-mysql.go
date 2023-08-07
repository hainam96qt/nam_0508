package wager

import (
	"context"
	"database/sql"

	db "nam_0508/internal/repo/dbmodel"
)

type Repository struct {
	Query        *db.Queries
	DatabaseConn *sql.DB
}

func NewMysqlRepository(databaseConn *sql.DB) *Repository {
	query := db.New(databaseConn)
	return &Repository{
		Query:        query,
		DatabaseConn: databaseConn,
	}
}

func (r *Repository) ListWagers(ctx context.Context, params db.ListWagersParams) ([]db.Wager, error) {
	// validate
	wagers, err := r.Query.ListWagers(ctx, params)
	if err != nil {
		return nil, err
	}

	return wagers, nil
}

func (r *Repository) CreateWager(ctx context.Context, wager db.CreateWagerParams) (*db.Wager, error) {
	// validate
	err := r.Query.CreateWager(ctx, wager)
	if err != nil {
		return nil, err
	}

	id, err := r.Query.LastInsertID(ctx)
	if err != nil {
		return nil, err
	}

	newWagerDB, err := r.Query.GetWager(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &newWagerDB, nil
}

func (r *Repository) GetWagerForUpdate(ctx context.Context, tx *sql.Tx, wagerID int32) (*db.Wager, error) {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	wagerDB, err := query.GetWagerForUpdate(ctx, wagerID)
	if err != nil {
		return nil, err
	}
	return &wagerDB, nil
}

func (r *Repository) UpdatePurchaseWager(ctx context.Context, tx *sql.Tx, updateWagerParams db.UpdatePurchaseWagerParams) error {
	var query = r.Query
	if tx != nil {
		query = query.WithTx(tx)
	}
	err := query.UpdatePurchaseWager(ctx, updateWagerParams)
	if err != nil {
		return err
	}
	return nil
}
