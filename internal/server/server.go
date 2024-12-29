package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ronaldalds/res/internal/handlers"
)

// New cria uma nova instância do FiberServer, inicializando o Fiber e o banco.
func New() *fiber.App {
	// Cria o servidor Fiber encapsulado na estrutura
	server := fiber.New(fiber.Config{
		ServerHeader: "P.O.N.C.H.E",
		AppName:      "P.O.N.C.H.E",
		ErrorHandler: handlers.ErrorHandler,
	})
	return server
}
