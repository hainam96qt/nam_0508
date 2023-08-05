package purchase

import (
	"context"
	"database/sql"
	"net/http"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	error2 "nam_0508/pkg/error"
	convert_type "nam_0508/pkg/util/convert-type"
)

func (s *Service) CreatePurchase(ctx context.Context, req *model.CreatePurchaseRequest) (*model.CreatePurchaseResponse, error) {
	tx, err := s.DatabaseConn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func(tx *sql.Tx) {
		if r := recover(); r != nil {
			err, _ = r.(error)
			tx.Rollback()
			return
		}
		tx.Commit()
	}(tx)

	wagerDB, err := s.Query.WithTx(tx).GetWagerForUpdate(ctx, req.WagerID)
	if err != nil {
		return nil, err
	}

	if req.BuyingPrice > wagerDB.CurrentSellingPrice {
		return nil, error2.NewXError("invalid selling percentage", http.StatusBadRequest)
	}

	err = s.Query.WithTx(tx).UpdatePurchaseWager(ctx, db.UpdatePurchaseWagerParams{
		CurrentSellingPrice: req.BuyingPrice,
		PercentageSold:      convert_type.NewNullInt32((wagerDB.AmountSold.Int32 + 1) / wagerDB.Odds),
		AmountSold:          convert_type.NewNullInt32(wagerDB.AmountSold.Int32 + 1),
	})
	if err != nil {
		return nil, err
	}

	newPurchase := db.CreatePurchaseParams{
		WagerID:     req.WagerID,
		BuyingPrice: req.BuyingPrice,
	}
	err = s.Query.WithTx(tx).CreatePurchase(ctx, newPurchase)
	if err != nil {
		return nil, err
	}

	id, err := s.Query.WithTx(tx).LastInsertID(ctx)
	if err != nil {
		return nil, err
	}

	purchaseDB, err := s.Query.WithTx(tx).GetPurchase(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return &model.CreatePurchaseResponse{
		Purchase: convertPurchaseDBToAPI(purchaseDB),
	}, nil
}

func convertPurchaseDBToAPI(purchase db.Purchase) model.Purchase {
	return model.Purchase{
		ID:          purchase.ID,
		WagerID:     purchase.WagerID,
		BuyingPrice: purchase.BuyingPrice,
		BoughtAt:    purchase.CreatedAt.Time,
	}
}
