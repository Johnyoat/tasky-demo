package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	database "github.com/johnyoat/tasky-demo/tasky-api/internal/db"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	app := fiber.New()

	db := database.InitDB("../data.db")
	if db == nil {
		log.Fatal("Failed to initialize database")
	}

	InitRouter(app)

	return &Server{App: app}
}

func (s *Server) Start() error {
	return s.App.Listen(":3000")
}
