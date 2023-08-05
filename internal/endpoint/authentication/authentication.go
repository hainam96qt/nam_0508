package authentication

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"nam_0508/internal/model"
	error2 "nam_0508/pkg/error"
	"nam_0508/pkg/util/request"
	"nam_0508/pkg/util/response"
)

type (
	Endpoint struct {
		authnSvc authenticationService
	}

	authenticationService interface {
		CreateRegistration(ctx context.Context, req *model.CreateRegistrationRequest) (*model.CreateRegistrationResponse, error)
		Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	}
)

func InitAuthenticationHandler(r *chi.Mux, authnSvc authenticationService) {
	registrationEndpoint := &Endpoint{
		authnSvc: authnSvc,
	}
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/registration", registrationEndpoint.createRegistration)
		r.Post("/login", registrationEndpoint.login)
	})
}

func (e *Endpoint) createRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateRegistrationRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	res, err := e.authnSvc.CreateRegistration(ctx, &req)
	if err != nil {
		log.Printf("failed to register new user: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	response.JSON(w, res)
}

func (e *Endpoint) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.LoginRequest
	if err := request.DecodeJSON(ctx, r.Body, &req); err != nil {
		log.Printf("read request body error: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	res, err := e.authnSvc.Login(ctx, &req)
	if err != nil {
		log.Printf("failed to register new user: %s \n", err)
		response.JSON(w, error2.NewXError(err.Error()))
		return
	}

	response.JSON(w, res)
}
