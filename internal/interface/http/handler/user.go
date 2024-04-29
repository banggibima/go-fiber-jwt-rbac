package handler

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/service"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService       service.UserService
	ResponsePresenter *presenter.ResponsePresenter
	Config            *config.Config
}

func NewUserHandler(
	userService service.UserService,
	responsePresenter *presenter.ResponsePresenter,
	config *config.Config,
) *UserHandler {
	return &UserHandler{
		UserService:       userService,
		ResponsePresenter: responsePresenter,
		Config:            config,
	}
}

func (h *UserHandler) ReadAll(c *fiber.Ctx) error {
	failed := "failed to fetch users"
	success := "successfully to fetch users"

	users, err := h.UserService.ReadAll()
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, users)
}

func (h *UserHandler) ReadByID(c *fiber.Ctx) error {
	failed := "failed to fetch user by id"
	success := "successfully to fetch user by id"

	id := c.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	user, err := h.UserService.ReadByID(uuid)
	if err != nil {
		if err.Error() == "record not found" {
			return h.ResponsePresenter.SetError(c, fiber.StatusNotFound, failed, err)
		}

		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, user)
}

func (h *UserHandler) ReadByUsername(c *fiber.Ctx) error {
	failed := "failed to fetch user by username"
	success := "successfully to fetch user by username"

	username := c.Params("username")

	user, err := h.UserService.ReadByUsername(username)
	if err != nil {
		if err.Error() == "record not found" {
			return h.ResponsePresenter.SetError(c, fiber.StatusNotFound, failed, err)
		}

		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, user)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	failed := "failed to create user"
	success := "successfully to create user"

	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	if err := h.UserService.Create(user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusCreated, success, user)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	failed := "failed to update user"
	success := "successfully to update user"

	id := c.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	exist, err := h.UserService.ReadByID(uuid)
	if err != nil {
		if err.Error() == "record not found" {
			return h.ResponsePresenter.SetError(c, fiber.StatusNotFound, failed, err)
		}

		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	user := exist

	if err := c.BodyParser(user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	if err := h.UserService.Update(uuid, user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	failed := "failed to delete user"
	success := "successfully to delete user"

	id := c.Params("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	exist, err := h.UserService.ReadByID(uuid)
	if err != nil {
		if err.Error() == "record not found" {
			return h.ResponsePresenter.SetError(c, fiber.StatusNotFound, failed, err)
		}

		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	uuid = exist.ID

	if err := h.UserService.Delete(uuid); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, nil)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	failed := "failed to login"
	success := "successfully to login"

	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	token, err := h.UserService.Login(user.Username, user.Password)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, token)
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	failed := "failed to register"
	success := "successfully to register"

	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	token, err := h.UserService.Register(user)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusCreated, success, token)
}
