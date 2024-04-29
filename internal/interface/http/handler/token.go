package handler

import (
	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/entity"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/domain/service"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/presenter"
	"github.com/gofiber/fiber/v2"
)

type TokenHandler struct {
	TokenService      service.TokenService
	ResponsePresenter *presenter.ResponsePresenter
	Config            *config.Config
}

func NewTokenHandler(
	tokenService service.TokenService,
	responsePresenter *presenter.ResponsePresenter,
	config *config.Config,
) *TokenHandler {
	return &TokenHandler{
		TokenService:      tokenService,
		ResponsePresenter: responsePresenter,
		Config:            config,
	}
}

func (h *TokenHandler) ReadByRefreshToken(c *fiber.Ctx) error {
	failed := "failed to fetch token by refresh token"
	success := "successfully to fetch token by refresh token"

	refreshToken := c.Params("refresh_token")

	token, err := h.TokenService.ReadByRefreshToken(refreshToken)
	if err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, token)
}

func (h *TokenHandler) Create(c *fiber.Ctx) error {
	failed := "failed to create token"
	success := "successfully to create token"

	token := new(entity.Token)

	if err := c.BodyParser(token); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusBadRequest, failed, err)
	}

	if err := h.TokenService.Create(token); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusCreated, success, token)
}

func (h *TokenHandler) DeleteByRefreshToken(c *fiber.Ctx) error {
	failed := "failed to delete token by refresh token"
	success := "successfully to delete token by refresh token"

	refreshToken := c.Params("refresh_token")

	if err := h.TokenService.DeleteByRefreshToken(refreshToken); err != nil {
		return h.ResponsePresenter.SetError(c, fiber.StatusInternalServerError, failed, err)
	}

	return h.ResponsePresenter.SetSuccess(c, fiber.StatusOK, success, nil)
}
