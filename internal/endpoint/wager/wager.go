package wager

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
		wagerSvc WagerService
	}

	WagerService interface {
		CreateWager(ctx context.Context, req *model.CreateWagerRequest) (*model.CreateWagerResponse, error)
		ListWagers(ctx context.Context, req *model.ListWagerRequest) (*model.ListWagerResponse, error)
	}
)

func InitWagerHandler(r *chi.Mux, wagerSvc WagerService) {
	wagerEndpoint := &Endpoint{
		wagerSvc: wagerSvc,
	}
	r.Route("/api/v1/wagers", func(r chi.Router) {
		r.Post("/", wagerEndpoint.createWager)
		r.Get("/", wagerEndpoint.listWagers)
	})
}

func (e *Endpoint) createWager(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateWagerRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	res, err := e.wagerSvc.CreateWager(ctx, &req)
	if err != nil {
		log.Printf("failed to register new user: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	response.JSON(w, http.StatusCreated, res)
}

func (e *Endpoint) listWagers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		log.Printf("failed to get query 'page' for list wagers: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		log.Printf("failed to get query 'limit' for list wagers: %s \n", err)
		response.Error(w, error2.NewXError(err.Error(), http.StatusBadRequest))
		return
	}

	var req = model.ListWagerRequest{
		Page:  page,
		Limit: limit,
	}

	res, err := e.wagerSvc.ListWagers(ctx, &req)
	if err != nil {
		log.Printf("failed to get list wagers: %s \n", err)
		response.Error(w, err)
		return
	}

	response.JSON(w, http.StatusOK, res)
}
