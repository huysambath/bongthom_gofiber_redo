package user

import (
	"admin-api/pkg/utls"
	"time"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	Id         int        `json:"id" db:"id"`
	FirstName  string     `json:"first_name" db:"first_name"`
	LastName   string     `json:"last_name" db:"last_name"`
	UserName   string     `json:"user_name" db:"user_name"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"password" db:"password"`
	RoleName   string     `json:"role_name" db:"role_name"`
	RoleId     int        `json:"role_id" db:"role_id"`
	IsAdmin    bool       `json:"is_admin" db:"is_admin"`
	LoginSession *string  `json:"login_session" db:"login_session"`
	LastLogin    *time.Time `json:"last_login" db:"last_login"`
	CurrencyId  *int     `json:"currency_id" db:"currency_id"`
	LanguageId  *int     `json:"language_id" db:"language_id"`
	StatusId    *int     `json:"status_id" db:"status_id"`
	Order       *int     `json:"order" db:"order"`
	CreatedBy   *int     `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedBy   *int     `json:"updated_by" db:"updated_by"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
	DeletedBy   *int     `json:"deleted_by" db:"deleted_by"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
	RoleId int `json:"role_id"`
	IsAdmin bool `json:"is_admin"`
	CreatedBy int `json:"created_by"`
}

func (r *UserCreateRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
	RoleId int `json:"role_id"`
	IsAdmin bool `json:"is_admin"`
	CreatedBy int `json:"created_by"`
}

func (r *UserUpdateRequest) bind(c fiber.Ctx, v *utls.Validator) error {
	if err := c.Bind().Body(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	return nil
}


type UserResponse struct {
	Users []User `json:"users"`
}

// type UserShowRequest struct {
// 	PageOption share.Paging   `json:"paging_options" query:"paging_options"`
// 	Sorts      []share.Sort   `json:"sorts,omitempty" query:"sorts"`
// 	Filters    []share.Filter `json:"filters.omitempty" query:"filters"`
// 	Search     string         `json:"q,omitempty" query:"q"`
// 	CurrencyID int            `json:"currency_id,omitempty" query:"currency_id"`
// }

// func (u *UserShowRequest) bind(c fiber.Ctx, v *utls.Validator) error {
// 	if err := c.Bind().Query(u); err != nil {
// 		return err
// 	}

// 	for i := range u.Filters {
// 		// request takes &filters[index][value]= int,
// 		value := c.Query(fmt.Sprintf("filters[%d][value]", i))
// 		if intValue, err := strconv.Atoi(value); err == nil {
// 			u.Filters[i].Value = intValue
// 		} else if boolValue, err := strconv.ParseBool(value); err == nil {
// 			u.Filters[i].Value = boolValue
// 		} else {
// 			u.Filters[i].Value = value
// 		}
// 	}
// 	if u.Search == "" {
// 		u.Search = c.Query("q")
// 	}
// 	if u.CurrencyID == 0 {
// 		if v := strings.TrimSpace(c.Query("currency_id")); v != "" {
// 			if n, err := strconv.Atoi(v); err == nil {
// 				u.CurrencyID = n
// 			}
// 		}
// 	}

// 	if err := v.Validate(u); err != nil {
// 		return err
// 	}
// 	return nil
// }

type UserShowRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
	RoleId int `json:"role_id"`
	IsAdmin bool `json:"is_admin"`
	CreatedBy int `json:"created_by"`
}