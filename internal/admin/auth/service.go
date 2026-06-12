package auth

import (
	error_response "admin-api/pkg/responses"
	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthService interface {
	Login(username string,password string) (*AuthLoginResponse, error)
}

type AuthServiceImpl struct {
	Repo AuthRepo
}

func NewAuthServiceImpl(db *sqlx.DB) *AuthServiceImpl{
	r := NewAuthRepoImpl(db)
	return &AuthServiceImpl{
		Repo: r,
	}
}

func (s *AuthServiceImpl) Login(username string,password string) (*AuthLoginResponse, error) {
	rs, err := s.Repo.Login(username,password)
	if err != nil {
		return nil, &error_response.ErrorResponse{
			MessageID: "database_error",
			Err:       err,
		}
	}

	if !rs {
		return nil, &error_response.ErrorResponse{
			MessageID: "login_failed",
			Err:       errors.New("invalid username or password"),
		}
	}

	var au AuthLoginResponse
	au.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
	au.TokenType = "jwt"

	return &au, nil
   
}