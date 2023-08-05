package model

import (
	"time"
)

type Purchase struct {
	ID          int32     `json:"id"`
	WagerID     int32     `json:"wager_id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}

type Wager struct {
	ID                  int32     `json:"id"`
	TotalWagerValue     int32     `json:"total_wager_value"`
	Odds                int32     `json:"odds"`
	SellingPercentage   int32     `json:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price"`
	PercentageSold      int32     `json:"percentage_sold"`
	AmountSold          int32     `json:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at"`
}

type CreateWagerRequest struct {
	TotalWagerValue     int32   `json:"total_wager_value"`
	Odds                int32   `json:"odds"`
	SellingPercentage   int32   `json:"selling_percentage"`
	SellingPrice        float64 `json:"selling_price"`
	CurrentSellingPrice float64 `json:"current_selling_price"`
	PercentageSold      int32   `json:"percentage_sold"`
	AmountSold          int32   `json:"amount_sold"`
}

type CreateWagerResponse struct {
	Wager
}

type ListWagerRequest struct {
	Page  int // from header
	Limit int // from header
}

type ListWagerResponse struct {
	Wagers []Wager
}

type CreatePurchaseRequest struct {
	WagerID     int32   // from header
	BuyingPrice float64 `json:"buying_price"`
}

type CreatePurchaseResponse struct {
	Purchase
}
