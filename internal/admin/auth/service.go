package auth

import (
	error_responses "admin-api/pkg/responses"
	"admin-api/pkg/share"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Login(username string,password string) (*AuthLoginResponse, *error_responses.ErrorResponse)
	CheckSession(loginSession string, userID float64) (bool, error)
}

type AuthServiceImpl struct {
	Repo AuthRepo
}

func NewAuthServiceImpl(db *sqlx.DB, rdb *redis.Client) *AuthServiceImpl{
	r := NewAuthRepoImpl(db, rdb)
	return &AuthServiceImpl{
		Repo: r,
	}
}

func (s *AuthServiceImpl) Login(username string, password string) (*AuthLoginResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	// If it wrong username and password it return error 
	user, errRes := s.Repo.Login(username, password)
	if errRes != nil {
		return nil, errRes
	}

	loginSession := uuid.New().String()

	// Call UpdateLoginSession
	if err := s.Repo.UpdateLoginSession(user.ID, loginSession); err != nil {
		return nil, msg.NewErrorResponse("session_error", fmt.Errorf("failed to update login session"))
	}

	// Generate JWT token
	tokenString, _, err := share.GenerateToken(float64(user.ID), user.UserName, loginSession, user.RoleId)
	if err != nil {
		return nil, msg.NewErrorResponse("token_error", fmt.Errorf("failed to gernerate token"))
	}

	var au AuthLoginResponse
	au.Auth.Token = tokenString
	au.Auth.TokenType = "jwt"
	return &au, nil
}

func (s *AuthServiceImpl) CheckSession(loginSession string, userID float64) (bool, error) {
	return s.Repo.CheckSession(loginSession, userID)
}