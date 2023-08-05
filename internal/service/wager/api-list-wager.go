package wager

import (
	"context"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
)

func (s *Service) ListWagers(ctx context.Context, req *model.ListWagerRequest) (*model.ListWagerResponse, error) {
	// validate
	wagers, err := s.Query.ListWagers(ctx, db.ListWagersParams{
		Limit:  int32(req.Limit),
		Offset: int32((req.Page - 1) * req.Limit),
	})
	if err != nil {
		return nil, err
	}

	return &model.ListWagerResponse{
		Wagers: convertWagersDBToAPI(wagers),
	}, nil
}

func convertWagersDBToAPI(wagerDBs []db.Wager) []model.Wager {
	var wagerAPIs []model.Wager
	for _, v := range wagerDBs {
		wagerAPIs = append(wagerAPIs, convertWagerDBToAPI(v))
	}
	return wagerAPIs
}
