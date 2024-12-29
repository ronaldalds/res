package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ronaldalds/res/internal/database"
)

type Middleware struct {
	App        *fiber.App
	GormStore  *database.GormStore
	RedisStore *database.RedisStore
}

func NewMiddleware(app *fiber.App) *Middleware {
	return &Middleware{
		App:        app,
		GormStore:  database.DB.GormStore,
		RedisStore: database.DB.RedisStore,
	}
}
