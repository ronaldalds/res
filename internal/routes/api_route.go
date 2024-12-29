package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ronaldalds/res/internal/controllers"
	"github.com/ronaldalds/res/internal/middlewares"
)

type Router struct {
	App        *fiber.App
	Controller *controllers.Controller
	Middleware *middlewares.Middleware
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		App:        app,
		Controller: controllers.NewController(),
		Middleware: middlewares.NewMiddleware(app),
	}
}

func (r *Router) RegisterFiberRoutes() {
	r.Middleware.CorsMiddleware()
	r.Middleware.SecurityMiddleware()
	apiV2 := r.App.Group("/api/v2")
	apiV2.Get("/health", r.Controller.HealthHandler)
	// Grupo de autenticação
	authGroup := apiV2.Group("/auth", r.Middleware.Limited(10))
	r.Auth(authGroup)

	// Grupo de Users
	usersGroup := apiV2.Group("/users")
	r.User(usersGroup)
	r.Role(usersGroup)
	r.Permission(usersGroup)
}
