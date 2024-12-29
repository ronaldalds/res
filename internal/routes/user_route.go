package routes

import (
	"github.com/gofiber/fiber/v2"
)

func (r *Router) User(router fiber.Router) {
	router.Post(
		"/",
		r.Middleware.JWTProtected("create_user"),
		r.Controller.CreateUserHandler,
	)
	router.Put(
		"/:id",
		r.Middleware.JWTProtected(),
		r.Controller.UpdateUserHandler,
	)
}

func (r *Router) Role(router fiber.Router) {
	router.Post(
		"/roles",
		r.Middleware.JWTProtected("create_role"),
		r.Controller.RegisterRoleHandler,
	)
}

func (r *Router) Permission(router fiber.Router) {
	router.Post(
		"/permissions",
		r.Middleware.JWTProtected("create_permission"),
		r.Controller.RegisterPermissiontHandler,
	)
}
