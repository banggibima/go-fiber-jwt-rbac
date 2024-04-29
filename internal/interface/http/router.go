package http

import (
	"github.com/banggibima/go-fiber-jwt-rbac/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App                      *fiber.App
	UserHandler              *handler.UserHandler
	AuthenticationMiddleware fiber.Handler
	AuthorizationMiddleware  func([]string) fiber.Handler
}

func NewRouter(
	app *fiber.App,
) *Router {
	return &Router{
		App: app,
	}
}

func (r *Router) Public() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", r.UserHandler.Login)
	auth.Post("/register", r.UserHandler.Register)
}

func (r *Router) Protected() {
	api := r.App.Group("/api")

	user := api.Group("/users", r.AuthenticationMiddleware)
	user.Get("/", r.AuthorizationMiddleware([]string{"admin", "user"}), r.UserHandler.ReadAll)
	user.Get("/:id", r.AuthorizationMiddleware([]string{"admin", "user"}), r.UserHandler.ReadByID)
	user.Get("/username/:username", r.AuthorizationMiddleware([]string{"admin", "user"}), r.UserHandler.ReadByUsername)
	user.Post("/", r.AuthorizationMiddleware([]string{"admin"}), r.UserHandler.Create)
	user.Put("/:id", r.AuthorizationMiddleware([]string{"admin"}), r.UserHandler.Update)
	user.Delete("/:id", r.AuthorizationMiddleware([]string{"admin"}), r.UserHandler.Delete)
}
