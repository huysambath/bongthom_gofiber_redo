package auth

import "github.com/jmoiron/sqlx"

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
   rs, _ := s.Repo.Login(username,password)

   if !rs {
		return nil, nil
   }

   var au AuthLoginResponse
   au.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
   au.TokenType = "jwt"

   return &au, nil
   
}