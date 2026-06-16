package user

import(
	error_responses "admin-api/pkg/responses"
)

type UserService interface {
	Show(userRequest *UserShowRequest) (*UserResponse, *error_responses.ErrorResponse)
	ShowOne(id int) (*UserResponse, *error_responses.ErrorResponse)
	Create(req *UserCreateRequest) (*UserResponse, *error_responses.ErrorResponse)
	Update( id int, req *UserUpdateRequest) (*UserResponse, *error_responses.ErrorResponse)
	Delete(id int) *error_responses.ErrorResponse
}

type UserServiceImpl struct {
	Repo UserRepo
}

func NewUserServiceImpl(r UserRepo) *UserServiceImpl {
	return &UserServiceImpl{
		Repo: r,
	}
}

func (r *UserServiceImpl) Show(userRequest *UserShowRequest) (*UserResponse, *error_responses.ErrorResponse){
	return r.Repo.Show(userRequest)
}

func (r *UserServiceImpl) ShowOne(id int) (*UserResponse, *error_responses.ErrorResponse) {
	return r.Repo.ShowOne(id)
}

func (r *UserServiceImpl) Create(req *UserCreateRequest) (*UserResponse, *error_responses.ErrorResponse) {
	return r.Repo.Create(req)
}

func (r *UserServiceImpl) Update(id int, req *UserUpdateRequest) (*UserResponse, *error_responses.ErrorResponse) {
	return r.Repo.Update(id, req)
}

func (r *UserServiceImpl) Delete(id int) *error_responses.ErrorResponse {
	return r.Repo.Delete(id)
}