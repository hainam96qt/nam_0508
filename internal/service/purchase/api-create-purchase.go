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

	wagerDB, err := s.wagerRepo.GetWagerForUpdate(ctx, tx, req.WagerID)
	if err != nil {
		return nil, err
	}

	if req.BuyingPrice > wagerDB.CurrentSellingPrice {
		return nil, error2.NewXError("invalid buy price", http.StatusBadRequest)
	}

	err = s.wagerRepo.UpdatePurchaseWager(ctx, tx, db.UpdatePurchaseWagerParams{
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
	purchaseDB, err := s.purchaseRepo.CreatePurchase(ctx, tx, newPurchase)
	if err != nil {
		return nil, err
	}

	return &model.CreatePurchaseResponse{
		Purchase: convertPurchaseDBToAPI(*purchaseDB),
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
