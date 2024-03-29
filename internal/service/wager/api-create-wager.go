package wager

import (
	"context"
	"net/http"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	error2 "nam_0508/pkg/error"
)

func (s *Service) CreateWager(ctx context.Context, req *model.CreateWagerRequest) (*model.CreateWagerResponse, error) {
	if err := validateWager(req); err != nil {
		return nil, err
	}

	newWager := db.CreateWagerParams{
		TotalWagerValue:     req.TotalWagerValue,
		Odds:                req.Odds,
		SellingPercentage:   req.SellingPercentage,
		SellingPrice:        req.SellingPrice,
		CurrentSellingPrice: req.SellingPrice,
	}
	newWagerDB, err := s.wagerRepo.CreateWager(ctx, newWager)
	if err != nil {
		return nil, err
	}

	return &model.CreateWagerResponse{
		Wager: convertWagerDBToAPI(*newWagerDB),
	}, nil
}

func validateWager(req *model.CreateWagerRequest) error {
	if req.TotalWagerValue < 1 {
		return error2.NewXError("invalid total wager value", http.StatusBadRequest)
	}
	if req.Odds < 1 {
		return error2.NewXError("invalid odds", http.StatusBadRequest)
	}
	if req.SellingPercentage < 1 || req.SellingPercentage > 100 {
		return error2.NewXError("invalid selling percentage", http.StatusBadRequest)
	}
	if req.SellingPrice < float64(req.TotalWagerValue)*(float64(req.SellingPercentage)/100) {
		return error2.NewXError("invalid selling price", http.StatusBadRequest)
	}
	return nil
}

func convertWagerDBToAPI(wager db.Wager) model.Wager {
	return model.Wager{
		ID:                  wager.ID,
		TotalWagerValue:     wager.TotalWagerValue,
		Odds:                wager.Odds,
		SellingPercentage:   wager.SellingPercentage,
		SellingPrice:        wager.SellingPrice,
		CurrentSellingPrice: wager.CurrentSellingPrice,
		PercentageSold:      wager.PercentageSold.Int32,
		AmountSold:          wager.AmountSold.Int32,
		PlacedAt:            wager.CreatedAt.Time,
	}
}
