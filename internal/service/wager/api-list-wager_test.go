package wager

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	mock_wager "nam_0508/internal/service/wager/mock"
	convert_type "nam_0508/pkg/util/convert-type"
)

func TestService_ListWagers(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock_wager.NewMockWagerRepository(ctrl)

	testWagerService := &Service{
		wagerRepo: wagerRepoMock,
	}

	t.Run("happy case", func(t *testing.T) {
		var req = &model.ListWagerRequest{
			Page:  2,
			Limit: 10,
		}
		var now = time.Now()
		var dbResult = []db.Wager{
			{
				ID:                  1,
				TotalWagerValue:     1,
				Odds:                1,
				SellingPercentage:   1,
				SellingPrice:        100.20,
				CurrentSellingPrice: 10000,
				PercentageSold:      convert_type.NewNullInt32(1),
				AmountSold:          convert_type.NewNullInt32(1),
				CreatedAt:           convert_type.NewNullTime(now),
				UpdatedAt:           convert_type.NewNullTime(now),
			},
			{
				ID: 2,
			},
		}
		wagerRepoMock.EXPECT().ListWagers(ctx, db.ListWagersParams{
			Limit:  int32(10),
			Offset: int32(10),
		}).Return(dbResult, nil).Times(1)
		result, err := testWagerService.ListWagers(ctx, req)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, len(result.Wagers), 2)
		assert.Equal(t, result.Wagers[0].TotalWagerValue, dbResult[0].TotalWagerValue)
		assert.Equal(t, result.Wagers[0].Odds, dbResult[0].Odds)
		assert.Equal(t, result.Wagers[0].SellingPercentage, dbResult[0].SellingPercentage)
		assert.Equal(t, result.Wagers[0].SellingPrice, dbResult[0].SellingPrice)
		assert.Equal(t, result.Wagers[0].CurrentSellingPrice, dbResult[0].CurrentSellingPrice)
		assert.Equal(t, result.Wagers[0].PercentageSold, dbResult[0].PercentageSold.Int32)
		assert.Equal(t, result.Wagers[0].AmountSold, dbResult[0].AmountSold.Int32)
		assert.Equal(t, result.Wagers[0].PlacedAt, dbResult[0].CreatedAt.Time)
	})

	t.Run("failed case", func(t *testing.T) {
		var req = &model.ListWagerRequest{
			Page:  2,
			Limit: 10,
		}

		wagerRepoMock.EXPECT().ListWagers(ctx, db.ListWagersParams{
			Limit:  int32(10),
			Offset: int32(10),
		}).Return(nil, errors.New("this is db error")).Times(1)
		_, err := testWagerService.ListWagers(ctx, req)

		assert.NotNil(t, err, "Error should not be nil")
	})
}
