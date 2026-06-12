package auth

import (
	// Community packages
	"github.com/gofiber/fiber/v3"

	// Internal packages
	constants "admin-api/pkg/constants"
	response "admin-api/pkg/http"
	"admin-api/pkg/translate"
	"admin-api/pkg/utls"
)

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(s AuthService) *AuthHandler{
	return &AuthHandler{
		Service: s,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req AuthRequest
	v := utls.NewValidator()

	if err := req.bind(c,v); err != nil {
		return err
	}

	rs, err := h.Service.Login(req.Username,req.Password)

	if err != nil {
		msg, e_msg := translate.TranslateWithError(c, "login_invalid")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Login_failed,
				err,
			),
		)
	} else {
		msg, e_msg := translate.TranslateWithError(c, "login_success")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			response.NewResponse(
				msg, constants.Login_success, rs,
			),
		)
	}
} 