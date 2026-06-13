package share

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt"
)

type UserContent struct {
	UserID float64 `json:"user_id"`
	UserName string `json:"user_name"`
	LoginSession string `json:"login_session"`
	Exp time.Time `json:"exp"`
	UserAgent string `json:"user_agent"`
	Ip string `json:"ip"`
	RoleID int `json:"role_id"`
}

type UserFromCreate struct {
	FirstName string
	LastName string
}

type UserFromUpdate struct {
	ID int8
	FirstName string
	LastName string
	NationalityID int 
	Nationality Nationality
}

type Nationality struct {
	ID int64
	Value string
}

type Paging struct {
	Page int `json:"page" query:"page" validate:"omitepty,min=1"`
	Perpage int `json:"per_page" query:"per_page" validate:"omitempty,min=1"`
}

type Sort struct {
	Property string `json:"property" validate:"requried"`
	Direction string `json:"direction" validate:"required,oneof=asc desc"`
}

type Filter struct {
	Property string `json:"property" validate:"required" query:"preperty"`
	Value interface{} `json:"value" validate:"required" query:"value"`
}

type Status struct {
	ID int 	`json:"id"`
	Name string `json:"name"`
}

var StatusData = []Status{
	{ID: 1, Name: "Active"},
	{ID: 2, Name: "Inactive"},
	{ID: 3, Name: "Suspended"},
	{ID: 4, Name: "Deleted"},
}

func GenerateToken(userID float64, userName string, loginSession string, roleID int) (string, time.Time, error) {
	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Claims is payload(data) stored inside the JWT token
	claims := jwt.MapClaims{
		"uesr_id": userID,
		"user_name": userName,
		"login_session": loginSession,
		"exp": expirationTime.Unix(),
		"role_id": roleID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}