package wager

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	mock_wager "nam_0508/internal/service/wager/mock"
)

//go:generate mockgen -source=service.go -destination=mock/wager_service_mock.go

func TestService_CreateWager(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock_wager.NewMockWagerRepository(ctrl)

	var dbConn = sql.DB{}
	testWagerService := &Service{
		DatabaseConn: &dbConn,
		wagerRepo:    wagerRepoMock,
	}

	var req = &model.CreateWagerRequest{
		TotalWagerValue:   1,
		Odds:              1,
		SellingPercentage: 1,
		SellingPrice:      100.20,
	}
	newWager := db.CreateWagerParams{
		TotalWagerValue:   req.TotalWagerValue,
		Odds:              req.Odds,
		SellingPercentage: req.SellingPercentage,
		SellingPrice:      req.SellingPrice,
	}
	var now = time.Now()
	dbResult := db.Wager{
		ID:                  1111,
		TotalWagerValue:     req.TotalWagerValue,
		Odds:                req.Odds,
		SellingPercentage:   req.SellingPercentage,
		SellingPrice:        req.SellingPrice,
		CurrentSellingPrice: 0,
		PercentageSold:      sql.NullInt32{},
		AmountSold:          sql.NullInt32{},
		CreatedAt:           sql.NullTime{Time: now, Valid: true},
		UpdatedAt:           sql.NullTime{Time: now, Valid: true},
	}
	t.Run("happy case", func(t *testing.T) {
		wagerRepoMock.EXPECT().CreateWager(ctx, newWager).Return(&dbResult, nil).Times(1)
		result, err := testWagerService.CreateWager(ctx, req)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, result.ID, 1111)
		assert.Equal(t, result.TotalWagerValue, req.TotalWagerValue)
		assert.Equal(t, result.Odds, req.Odds)
		assert.Equal(t, result.SellingPercentage, req.SellingPercentage)
		assert.Equal(t, result.SellingPrice, req.SellingPrice)
		assert.Equal(t, result.CurrentSellingPrice, 0.0)
		assert.Equal(t, result.PercentageSold, int32(0))
		assert.Equal(t, result.AmountSold, int32(0))
		assert.Equal(t, result.PlacedAt, now)
	})

	t.Run("failed case - db error", func(t *testing.T) {
		wagerRepoMock.EXPECT().CreateWager(ctx, newWager).Return(nil, errors.New("this is db error")).Times(1)
		_, err := testWagerService.CreateWager(ctx, req)
		assert.NotNil(t, err, "Error should not be nil")
	})

	t.Run("failed case - invalid total wager value", func(t *testing.T) {
		var req = &model.CreateWagerRequest{
			TotalWagerValue:   0, // fail
			Odds:              0,
			SellingPercentage: 0,
			SellingPrice:      0,
		}
		_, err := testWagerService.CreateWager(ctx, req)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, err.Error(), "invalid total wager value")
	})

	t.Run("failed case - invalid odds", func(t *testing.T) {
		var req = &model.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              0, // fail
			SellingPercentage: 0,
			SellingPrice:      0,
		}
		_, err := testWagerService.CreateWager(ctx, req)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, err.Error(), "invalid odds")
	})

	t.Run("failed case - invalid selling percentage", func(t *testing.T) {
		var req = &model.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 0, // fail
			SellingPrice:      0,
		}
		_, err := testWagerService.CreateWager(ctx, req)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, err.Error(), "invalid selling percentage")
	})

	t.Run("failed case - invalid selling price", func(t *testing.T) {
		var req = &model.CreateWagerRequest{
			TotalWagerValue:   1,
			Odds:              1,
			SellingPercentage: 1,
			SellingPrice:      0, // fail
		}
		_, err := testWagerService.CreateWager(ctx, req)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, err.Error(), "invalid selling price")
	})
}
