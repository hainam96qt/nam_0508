package purchase

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"nam_0508/internal/model"
	error2 "nam_0508/pkg/error"
	"nam_0508/pkg/util/request"
	"nam_0508/pkg/util/response"
)

type (
	Endpoint struct {
		purchaseSvc PurchaseService
	}

	PurchaseService interface {
		CreatePurchase(ctx context.Context, req *model.CreatePurchaseRequest) (*model.CreatePurchaseResponse, error)
	}
)

func InitPurchaseHandler(r *chi.Mux, purchaseSvc PurchaseService) {
	wagerEndpoint := &Endpoint{
		purchaseSvc: purchaseSvc,
	}
	r.Route("/", func(r chi.Router) {
		r.Post("/buy/{wager_id}", wagerEndpoint.createPurchase)
	})
}

func (e *Endpoint) createPurchase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	wagerID, err := strconv.Atoi(chi.URLParam(r, "wager_id"))
	if err != nil {
		log.Printf("failed to get query 'page' for list wagers: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	var req model.CreatePurchaseRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}
	req.WagerID = int32(wagerID)

	res, err := e.purchaseSvc.CreatePurchase(ctx, &req)
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusCreated, res)
}
