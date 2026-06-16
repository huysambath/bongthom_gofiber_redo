package user

import (
	"admin-api/pkg/constants"
	response "admin-api/pkg/http"
	"admin-api/pkg/translate"
	"admin-api/pkg/utls"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) Show(c fiber.Ctx) error {
	rs, e := h.Service.Show(&UserShowRequest{})

	if e != nil {
		msg := translate.Translate(c, e.MessageID)
		return c.Status(fiber.StatusInternalServerError).JSON(
			response.NewResponseError(
				msg,
				constants.Get_user_fail,
				e.Err,
			),
		)
	}

	msg := translate.Translate(c, "get_users_success")
	return c.Status(fiber.StatusOK).JSON(
		response.NewResponse(
			msg,
			constants.Get_user_success,
			rs,
		),
	)
}

func (h *UserHandler) ShowOne(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		msg, e_msg := translate.TranslateWithError(c, "invalid_id")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(msg, constants.User_invaild_id,err),
		)
	}

	rs, e := h.Service.ShowOne(id)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, "get_user_failed")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusNotFound).JSON(
			response.NewResponseError(msg, constants.Get_user_fail, err),
		)
	}

	msg, e_msg := translate.TranslateWithError(c,"get_user_success")
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
			msg,
			constants.Get_user_success,
			rs,
		),
	)
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	req := UserCreateRequest{}
	v := utls.NewValidator()

	if err := req.bind(c,v); err != nil {
		return err
	}

	rs, e := h.Service.Create(&req)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, "create_user_failed")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Create_user_failed,
				e,
			),
		)
	}
	msg, e_msg := translate.TranslateWithError(c, "get_user_success")
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
			msg,
			constants.Creat_user_success,
			rs,
		),
	)
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		msg, e_msg := translate.TranslateWithError(c, "invaild_id")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.User_invaild_id,
				err,
			),
		)
	}

	req := UserUpdateRequest{}
	v := utls.NewValidator()

	if err := req.bind(c,v); err != nil {
		return err
	}

	rs, e := h.Service.Update(id,&req)

	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, "update_user_failed")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Update_user_failed,
				e,
			),
		)
	}

	msg, e_msg := translate.TranslateWithError(c, "update_user_success")
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
		response.NewResponse(msg, constants.Update_user_success, rs),
	)
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		msg, e_msg := translate.TranslateWithError(c, "invaild_id")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.User_invaild_id,
				err,
			),
		)
	}

	e := h.Service.Delete(id)
	if e != nil {
		msg, e_msg := translate.TranslateWithError(c, "delete_user_failed")
		if e_msg != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				response.NewResponseError(
					e_msg.Err.Error(),
					constants.Translate_Failed,
					e_msg.Err,
				),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			response.NewResponseError(
				msg,
				constants.Delete_user_failed,
				e,
			),
		)
	}

	msg, e_msg := translate.TranslateWithError(c ,"deleted_user_success")
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
			msg,
			constants.Translate_succes,
			e,
		),
	)
}