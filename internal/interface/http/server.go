package http

import (
	"fmt"

	"github.com/banggibima/go-fiber-jwt-rbac/config"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/application/service"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/infrastructure/memory"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/infrastructure/repository"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/handler"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/middleware"
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	Config *config.Config
	GORM   *gorm.DB
	Redis  *redis.Client
	Fiber  *fiber.App
}

func NewServer(
	config *config.Config,
	gorm *gorm.DB,
	redis *redis.Client,
	fiber *fiber.App,
) *Server {
	return &Server{
		Config: config,
		GORM:   gorm,
		Redis:  redis,
		Fiber:  fiber,
	}
}

func (s *Server) Router() {
	userRepository := repository.NewUserRepository(s.GORM)
	tokenRepository := memory.NewTokenMemory(s.Redis)

	userService := service.NewUserService(userRepository, s.Config)
	tokenService := service.NewTokenService(tokenRepository, s.Config)

	responsePresenter := presenter.NewResponsePresenter()

	userHandler := handler.NewUserHandler(userService, responsePresenter, s.Config)
	tokenHandler := handler.NewTokenHandler(tokenService, responsePresenter, s.Config)

	authMiddleware := middleware.NewAuthMiddleware(responsePresenter, s.Config)

	router := &Router{
		App:                      s.Fiber,
		UserHandler:              userHandler,
		TokenHandler:             tokenHandler,
		AuthenticationMiddleware: authMiddleware.Authentication,
		AuthorizationMiddleware:  authMiddleware.Authorization,
	}

	router.Public()
	router.Protected()
}

func (s *Server) Start() error {
	s.Router()

	port := s.Config.HTTP.Port

	if err := s.Fiber.Listen(fmt.Sprintf(":%d", port)); err != nil {
		return err
	}

	return nil
}
