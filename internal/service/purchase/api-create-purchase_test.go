package purchase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	mock "nam_0508/internal/service/purchase/mock"
	convert_type "nam_0508/pkg/util/convert-type"
)

//go:generate mockgen -source=service.go -destination=mock/purchase_service_mock.go

func TestService_CreatePurchase_HappyCase(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock.NewMockWagerRepository(ctrl)
	purchaseMock := mock.NewMockPurchaseRepository(ctrl)
	dbConn, mockDBConn, err := sqlmock.New()
	assert.Nil(t, err, "Error should be nil")

	testPurchaseService := &Service{
		DatabaseConn: dbConn,
		wagerRepo:    wagerRepoMock,
		purchaseRepo: purchaseMock,
	}

	t.Run("happy case", func(t *testing.T) {
		var wagerTime = time.Now()
		var purchaseTime = wagerTime.Add(100 * time.Second)
		req := &model.CreatePurchaseRequest{
			WagerID:     1111,
			BuyingPrice: 1000,
		}
		dbWagerResult := db.Wager{
			ID:                  1111,
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        10000,
			CurrentSellingPrice: 9999,
			PercentageSold:      convert_type.NewNullInt32(1),
			AmountSold:          convert_type.NewNullInt32(1),
			CreatedAt:           convert_type.NewNullTime(wagerTime),
			UpdatedAt:           convert_type.NewNullTime(wagerTime),
		}
		wagerUpdateParams := db.UpdatePurchaseWagerParams{
			CurrentSellingPrice: req.BuyingPrice,
			PercentageSold:      convert_type.NewNullInt32((dbWagerResult.AmountSold.Int32 + 1) / dbWagerResult.Odds),
			AmountSold:          convert_type.NewNullInt32(dbWagerResult.AmountSold.Int32 + 1),
		}
		newPurchase := db.CreatePurchaseParams{
			WagerID:     req.WagerID,
			BuyingPrice: req.BuyingPrice,
		}
		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		mockDBConn.ExpectClose()
		wagerRepoMock.EXPECT().GetWagerForUpdate(ctx, gomock.Any(), req.WagerID).Return(&dbWagerResult, nil).Times(1)
		wagerRepoMock.EXPECT().UpdatePurchaseWager(ctx, gomock.Any(), wagerUpdateParams).Return(nil).Times(1)
		purchaseMock.EXPECT().CreatePurchase(ctx, gomock.Any(), newPurchase).Return(&db.Purchase{
			ID:          2222,
			WagerID:     1111,
			BuyingPrice: req.BuyingPrice,
			CreatedAt:   convert_type.NewNullTime(purchaseTime),
			UpdatedAt:   convert_type.NewNullTime(purchaseTime),
		}, nil).Times(1)
		result, err := testPurchaseService.CreatePurchase(ctx, req)

		assert.Nil(t, err, "Error should be nil")
		assert.Equal(t, result.ID, int32(2222))
		assert.Equal(t, result.BuyingPrice, req.BuyingPrice)
		assert.Equal(t, result.WagerID, req.WagerID)
		assert.Equal(t, result.BoughtAt, purchaseTime)
	})

}

func TestService_CreatePurchase_InvalidBuyPrice(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock.NewMockWagerRepository(ctrl)
	purchaseMock := mock.NewMockPurchaseRepository(ctrl)
	dbConn, mockDBConn, err := sqlmock.New()
	assert.Nil(t, err, "Error should be nil")

	testPurchaseService := &Service{
		DatabaseConn: dbConn,
		wagerRepo:    wagerRepoMock,
		purchaseRepo: purchaseMock,
	}

	t.Run("failed case - invalid buy price", func(t *testing.T) {
		var wagerTime = time.Now()
		req := &model.CreatePurchaseRequest{
			WagerID:     1111,
			BuyingPrice: 10000, // bigger than currently price
		}
		dbWagerResult := db.Wager{
			ID:                  1111,
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        10000,
			CurrentSellingPrice: 9999,
			PercentageSold:      convert_type.NewNullInt32(1),
			AmountSold:          convert_type.NewNullInt32(1),
			CreatedAt:           convert_type.NewNullTime(wagerTime),
			UpdatedAt:           convert_type.NewNullTime(wagerTime),
		}

		mockDBConn.ExpectBegin()
		mockDBConn.ExpectRollback()
		wagerRepoMock.EXPECT().GetWagerForUpdate(ctx, gomock.Any(), req.WagerID).Return(&dbWagerResult, nil).Times(1)
		_, err := testPurchaseService.CreatePurchase(ctx, req)

		assert.NotNil(t, err, "Error should be not nil")
	})
}

func TestService_CreatePurchase_FailedUpdate(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock.NewMockWagerRepository(ctrl)
	purchaseMock := mock.NewMockPurchaseRepository(ctrl)
	dbConn, mockDBConn, err := sqlmock.New()
	assert.Nil(t, err, "Error should be nil")

	testPurchaseService := &Service{
		DatabaseConn: dbConn,
		wagerRepo:    wagerRepoMock,
		purchaseRepo: purchaseMock,
	}

	t.Run("failed case - failed to update purchase wager", func(t *testing.T) {
		var wagerTime = time.Now()
		req := &model.CreatePurchaseRequest{
			WagerID:     1111,
			BuyingPrice: 1000,
		}
		dbWagerResult := db.Wager{
			ID:                  1111,
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        10000,
			CurrentSellingPrice: 9999,
			PercentageSold:      convert_type.NewNullInt32(1),
			AmountSold:          convert_type.NewNullInt32(1),
			CreatedAt:           convert_type.NewNullTime(wagerTime),
			UpdatedAt:           convert_type.NewNullTime(wagerTime),
		}
		wagerUpdateParams := db.UpdatePurchaseWagerParams{
			CurrentSellingPrice: req.BuyingPrice,
			PercentageSold:      convert_type.NewNullInt32((dbWagerResult.AmountSold.Int32 + 1) / dbWagerResult.Odds),
			AmountSold:          convert_type.NewNullInt32(dbWagerResult.AmountSold.Int32 + 1),
		}
		mockDBConn.ExpectBegin()
		mockDBConn.ExpectRollback()
		wagerRepoMock.EXPECT().GetWagerForUpdate(ctx, gomock.Any(), req.WagerID).Return(&dbWagerResult, nil).Times(1)
		wagerRepoMock.EXPECT().UpdatePurchaseWager(ctx, gomock.Any(), wagerUpdateParams).Return(errors.New("this is got error")).Times(1)

		_, err := testPurchaseService.CreatePurchase(ctx, req)
		assert.NotNil(t, err, "Error should be not nil")
	})
}

func TestService_CreatePurchase_FailedCreate(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wagerRepoMock := mock.NewMockWagerRepository(ctrl)
	purchaseMock := mock.NewMockPurchaseRepository(ctrl)
	dbConn, mockDBConn, err := sqlmock.New()
	assert.Nil(t, err, "Error should be nil")

	testPurchaseService := &Service{
		DatabaseConn: dbConn,
		wagerRepo:    wagerRepoMock,
		purchaseRepo: purchaseMock,
	}

	t.Run("happy case - failed to create purchase", func(t *testing.T) {
		var wagerTime = time.Now()
		req := &model.CreatePurchaseRequest{
			WagerID:     1111,
			BuyingPrice: 1000,
		}
		dbWagerResult := db.Wager{
			ID:                  1111,
			TotalWagerValue:     1,
			Odds:                1,
			SellingPercentage:   1,
			SellingPrice:        10000,
			CurrentSellingPrice: 9999,
			PercentageSold:      convert_type.NewNullInt32(1),
			AmountSold:          convert_type.NewNullInt32(1),
			CreatedAt:           convert_type.NewNullTime(wagerTime),
			UpdatedAt:           convert_type.NewNullTime(wagerTime),
		}
		wagerUpdateParams := db.UpdatePurchaseWagerParams{
			CurrentSellingPrice: req.BuyingPrice,
			PercentageSold:      convert_type.NewNullInt32((dbWagerResult.AmountSold.Int32 + 1) / dbWagerResult.Odds),
			AmountSold:          convert_type.NewNullInt32(dbWagerResult.AmountSold.Int32 + 1),
		}
		newPurchase := db.CreatePurchaseParams{
			WagerID:     req.WagerID,
			BuyingPrice: req.BuyingPrice,
		}
		mockDBConn.ExpectBegin()
		mockDBConn.ExpectCommit()
		wagerRepoMock.EXPECT().GetWagerForUpdate(ctx, gomock.Any(), req.WagerID).Return(&dbWagerResult, nil).Times(1)
		wagerRepoMock.EXPECT().UpdatePurchaseWager(ctx, gomock.Any(), wagerUpdateParams).Return(nil).Times(1)
		purchaseMock.EXPECT().CreatePurchase(ctx, gomock.Any(), newPurchase).Return(nil, errors.New("this is got error")).Times(1)
		_, err := testPurchaseService.CreatePurchase(ctx, req)
		assert.NotNil(t, err, "Error should be not nil")
	})

}
