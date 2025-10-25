package routes

import (
	"go-fiber/app/repository"
	"go-fiber/app/service"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func FileRoute(app fiber.Router, db *mongo.Database) {
	fileRepo := repository.NewFileRepository(db)

	route := app.Group("/files")

	route.Post("/upload", middleware.AuthRequired(), service.UploadFileService(fileRepo))
	route.Get("/", middleware.AuthRequired(), service.GetFilesService(fileRepo))
	route.Get("/:id", middleware.AuthRequired(), service.GetFileByIDService(fileRepo))
	route.Delete("/:id", middleware.AuthRequired(), service.DeleteFileService(fileRepo))
}
