package auth

import (

	"github.com/jmoiron/sqlx"
)

type AuthRepo interface {
	Login(username string, password string) (bool, error)
}

type AuthRepoImpl struct {
	db *sqlx.DB
}

func NewAuthRepoImpl(db *sqlx.DB) *AuthRepoImpl{
	return &AuthRepoImpl{
		db: db,
	}
}

func (r *AuthRepoImpl) Login(username string, password string) (bool, error) {
	var user Auth
	query := `SELECT id, user_name, password FROM tbl_users WHERE user_name = $1 AND password = $2 AND deleted_at IS NULL LIMIT 1`

    err := r.db.Get(&user, query, username, password)
    if err != nil {
        return false, nil
    }

	return true, nil
}
