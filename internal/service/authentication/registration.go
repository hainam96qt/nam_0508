package authentication

import (
	"context"
	"database/sql"

	"nam_0508/internal/model"
	db "nam_0508/internal/repo/dbmodel"
	convert_type "nam_0508/pkg/util/convert-type"
	"nam_0508/pkg/util/password"
)

func (s *Service) CreateRegistration(ctx context.Context, req *model.CreateRegistrationRequest) (*model.CreateRegistrationResponse, error) {
	tx, err := s.DatabaseConn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	passwordHashed, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := db.CreateUserParams{
		UserName: convert_type.NewNullString(req.UserName),
		Password: convert_type.NewNullString(passwordHashed),
	}
	err = s.Query.WithTx(tx).CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user, err := s.Query.WithTx(tx).GetUser(ctx, convert_type.NewNullString(req.UserName))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tokenPair, err := s.jwtSvc.GenerateTokenPair(int(user.ID), user.UserName.String)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &model.CreateRegistrationResponse{
		TokenPair: *tokenPair,
	}, nil
}
