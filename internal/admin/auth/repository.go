package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	responses "admin-api/pkg/responses"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRepo interface {
	Login(username string, password string) (*Auth, *responses.ErrorResponse)
	UpdateLoginSession(userID int, loginSession string) error
	CheckSession(loginSession string, userID float64) (bool, error)
}

type AuthRepoImpl struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewAuthRepoImpl(db *sqlx.DB, rdb *redis.Client) *AuthRepoImpl {
	return &AuthRepoImpl{
		db:  db,
		rdb: rdb,
	}
}

func (r *AuthRepoImpl) Login(username string, password string) (*Auth, *responses.ErrorResponse) {
	errResp := responses.ErrorResponse{}
	var user Auth
	query := `SELECT id, user_name, password, role_id FROM tbl_users WHERE user_name = $1 AND password = $2 AND deleted_at IS NULL LIMIT 1`

	err := r.db.Get(&user, query, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errResp.NewErrorResponse("invalid_credentials", fmt.Errorf("invalid username or password"))
		}
		return nil, errResp.NewErrorResponse("login_error", err)
	}

	return &user, nil
}

func (r *AuthRepoImpl) UpdateLoginSession(userID int, loginSession string) error {
	// 1. Update database
	query := `UPDATE tbl_users SET login_session = $1 WHERE id = $2`
	_, err := r.db.Exec(query, loginSession, userID)
	if err != nil {
		return err
	}

	// 2. Set in Redis
	ctx := context.Background()
	key := fmt.Sprintf("session:%d", userID)
	err = r.rdb.Set(ctx, key, loginSession, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepoImpl) CheckSession(loginSession string, userID float64) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("session:%d", int(userID))
	val, err := r.rdb.Get(ctx, key).Result()
	if err == nil {
		return val == loginSession, nil
	}
	if !errors.Is(err, redis.Nil) {
		// Log/return redis error or fallback to DB
	}

	// Fallback to database
	var dbSession string
	query := `SELECT login_session FROM tbl_users WHERE id = $1 AND deleted_at IS NULL LIMIT 1`
	dbErr := r.db.Get(&dbSession, query, int(userID))
	if dbErr != nil {
		if errors.Is(dbErr, sql.ErrNoRows) {
			return false, nil
		}
		return false, dbErr
	}
	return dbSession == loginSession, nil
}

