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

	db, err := database.InitDB("../data.db")
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	InitRouter(app, db)

	return &Server{App: app}
}

func (s *Server) Start() error {
	return s.App.Listen(":3000")
}
