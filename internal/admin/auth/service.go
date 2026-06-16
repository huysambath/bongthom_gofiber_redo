package auth

import (
	error_responses "admin-api/pkg/responses"
	"admin-api/pkg/share"
	"fmt"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(username string,password string) (*AuthLoginResponse, *error_responses.ErrorResponse)
	CheckSession(loginSession string, userID float64) (bool, error)
}

type AuthServiceImpl struct {
	Repo AuthRepo
}

func NewAuthServiceImpl(r AuthRepo) *AuthServiceImpl{
	return &AuthServiceImpl{
		Repo: r,
	}
}

func (s *AuthServiceImpl) Login(username string, password string) (*AuthLoginResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	// Validate credentials and fetch user
	user, errResp := s.Repo.Login(username, password)
	if errResp != nil {
		return nil, errResp
	}

	// Generate a new login session
	loginSession := uuid.New().String()

	if err := s.Repo.UpdateLoginSession(user.ID, loginSession); err != nil {
		return nil, msg.NewErrorResponse("session_error", fmt.Errorf("failed to update login session"))
	}

	// Generate JWT token
	tokenString, _, err := share.GenerateToken(float64(user.ID), user.UserName, loginSession, user.RoleId)
	if err != nil {
		return nil, msg.NewErrorResponse("token_error", fmt.Errorf("failed to generate token"))
	}

	// Build and return the response
	var au AuthLoginResponse
	au.Auth.Token = tokenString
	return &au, nil
}

func (s *AuthServiceImpl) CheckSession(loginSession string, userID float64) (bool, error) {
	return s.Repo.CheckSession(loginSession, userID)
}