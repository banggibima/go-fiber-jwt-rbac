package presenter

import (
	"github.com/gofiber/fiber/v2"
)

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ResponsePresenter struct {
	Success *Success
	Error   *Error
}

func NewResponsePresenter() *ResponsePresenter {
	return &ResponsePresenter{
		Success: &Success{},
		Error:   &Error{},
	}
}

func (r *ResponsePresenter) SetSuccess(c *fiber.Ctx, code int, message string, data interface{}) error {
	r.Success.Message = message
	r.Success.Data = data
	return c.Status(code).JSON(r.Success)
}

func (r *ResponsePresenter) SetError(c *fiber.Ctx, code int, message string, err error) error {
	r.Error.Message = message
	r.Error.Error = err.Error()
	return c.Status(code).JSON(r.Error)
}
