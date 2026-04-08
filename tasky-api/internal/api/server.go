package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewServer() *Server {
	app := fiber.New()
	db, err := gorm.Open(sqlite.Open("../data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	InitRouter(app, db)

	return &Server{App: app, DB: db}
}

func (s *Server) Start() error {
	return s.App.Listen(":3000")
}
