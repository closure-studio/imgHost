package server

import (
	"fmt"
	"log"

	S3 "github.com/closure-studio/objectStorage/server/controller/s3"
	"github.com/closure-studio/objectStorage/server/middleware"
	"github.com/closure-studio/objectStorage/server/utils/resp"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		AppName: "closure-studio img video host",
	})
	app.Use(middleware.JSONResponseMiddleware())

	api := app.Group("/api")
	api.Post("/upload", S3.UploadHandler())

	app.Get("/", func(c fiber.Ctx) error {
		// return an error
		// return fmt.Errorf("Hello, error")
		return resp.Success(c, nil, "Hello, World!")
	})

	return &Server{App: app}
}

func (s *Server) Start() {
	fmt.Println("🚀 Server running on port 8080")

	log.Fatal(s.App.Listen("0.0.0.0:8080"))
}
