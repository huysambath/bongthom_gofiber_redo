package user

import (
	error_responses "admin-api/pkg/responses"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Show(userRequest *UserShowRequest) (*UserResponse, *error_responses.ErrorResponse)
	ShowOne(id int) (*UserResponse, *error_responses.ErrorResponse)
	Create(req *UserCreateRequest) (*UserResponse, *error_responses.ErrorResponse)
	Update( id int, req *UserUpdateRequest) (*UserResponse, *error_responses.ErrorResponse)
	Delete(id int) *error_responses.ErrorResponse
}

type UserRepoImpl struct {
	DB *sqlx.DB
}

func NewUserRepoImpl(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{
		DB: db,
	}
}

func (r *UserRepoImpl) Show(userRequest *UserShowRequest) (*UserResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	
	var users []User

	err := r.DB.Select(
		&users,
		`SELECT id, first_name, last_name, user_name, email,
		        role_name, role_id, is_admin, login_session,
		        last_login, currency_id, language_id,
		        status_id, created_at, updated_at
		   FROM tbl_users WHERE deleted_at IS NULL`,
	)

	if err != nil {
		return nil, msg.NewErrorResponse("cannot select user from database", err)
	}
	return &UserResponse{Users: users}, nil
}

func (r *UserRepoImpl) ShowOne(id int) (*UserResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var user User
	err := r.DB.Get(&user,
		`SELECT id, first_name, last_name, user_name, email, role_name, role_id, is_admin, login_session, last_login, currency_id, language_id, status_id, created_at, updated_at FROM tbl_users WHERE id = $1 AND deleted_at IS NULL LIMIT 1`,id,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("user_not_found", err)
	}
	return &UserResponse{Users: []User{user}}, nil
}

func (r *UserRepoImpl) Create(req *UserCreateRequest) (*UserResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}

	var user int
	 err := r.DB.QueryRow(
		`INSERT INTO tbl_users (first_name, last_name, user_name, email, password, role_name, role_id, is_admin, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		req.FirstName, req.LastName, req.UserName, req.Email, req.Password, req.RoleName, req.RoleId, req.IsAdmin, req.CreatedBy,
	).Scan(&user)

	if err != nil {
		return nil, msg.NewErrorResponse("create_user_failed", fmt.Errorf("failed to create user: %w", err))
	}

	return &UserResponse{}, nil
}

func (r *UserRepoImpl) Update( id int, req *UserUpdateRequest) (*UserResponse, *error_responses.ErrorResponse) {
	msg := error_responses.ErrorResponse{}
	_,err := r.DB.Exec(
		`UPDATE tbl_users SET first_name = $1, last_name = $2, user_name = $3, email = $4, role_name = $5, role_id = $6, updated_at = NOW() WHERE id = $7`,
		req.FirstName, req.LastName, req.UserName, req.Email, req.RoleName, req.RoleId, id,
	)
	if err != nil {
		return nil, msg.NewErrorResponse("update_user_failed", fmt.Errorf("failed to update user: %w", err))
	}

	return &UserResponse{},nil
}

func (r *UserRepoImpl) Delete(id int) *error_responses.ErrorResponse {
	msg := error_responses.ErrorResponse{}
	_, err := r.DB.Exec(`DELETE FROM tbl_users WHERE id = $1`, id)
	if err != nil {
		return msg.NewErrorResponse("delete_user_failed", fmt.Errorf("failed to delete user: %w", err))
	}
	return nil
}