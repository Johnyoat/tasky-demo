package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	database "github.com/johnyoat/tasky-demo/tasky-api/internal/db"
)

// Server defines the HTTP server structure using the Fiber framework
type Server struct {
	App *fiber.App
}

// NewServer creates and configures a new Server instance with CORS and database connection
func NewServer() *Server {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin, Content-Type, Accept"},
		AllowMethods: []string{"GET, POST, PUT, DELETE"},
	}))

	db, err := database.InitDB("../data.db")
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	InitRouter(app, db)

	return &Server{App: app}
}

// Start begins the HTTP server and listens for incoming requests on port 3000
func (s *Server) Start() error {
	return s.App.Listen(":3000")
}
