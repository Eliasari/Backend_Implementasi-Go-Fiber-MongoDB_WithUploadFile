package config

import (
    "github.com/gofiber/fiber/v2"
    "go-fiber/middleware"
    "go-fiber/routes"
    "go.mongodb.org/mongo-driver/mongo"
)

func NewApp(db *mongo.Database) *fiber.App {
    app := fiber.New()
    app.Use(middleware.LoggerMiddleware)

    if db == nil {
        panic("❌ Database tidak boleh nil di NewApp()")
    }

    routes.RegisterRoutes(app, db)

    return app
}
