package middleware

import (
	"errors"
	"strings"

	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/presenter"
	pkgjwt "github.com/banggibima/go-fiber-jwt-rbac/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	ResponsePresenter *presenter.ResponsePresenter
	Config            *config.Config
}

func NewAuthMiddleware(
	responsePresenter *presenter.ResponsePresenter,
	config *config.Config,
) *AuthMiddleware {
	return &AuthMiddleware{
		ResponsePresenter: responsePresenter,
		Config:            config,
	}
}

func (m *AuthMiddleware) Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", errors.New("missing token"))
	}

	if !strings.HasPrefix(token, "Bearer ") {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", errors.New("invalid token"))
	}

	token = strings.TrimPrefix(token, "Bearer ")

	claims, err := pkgjwt.ValidateToken(m.Config, token)
	if err != nil {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", err)
	}

	c.Locals("claims", claims)

	return c.Next()
}

func (m *AuthMiddleware) Authorization(roles []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*jwt.Token)
		if !ok {
			return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", errors.New("missing claims"))
		}

		role := claims.Claims.(jwt.MapClaims)["sub"].(map[string]interface{})["role"].(string)

		match := false
		for _, r := range roles {
			if r == role {
				match = true
				break
			}
		}

		if !match {
			return m.ResponsePresenter.SetError(c, fiber.StatusForbidden, "forbidden", errors.New("role is not allowed"))
		}

		return c.Next()
	}
}
