package routes

import (
	"database/sql"
	"go-fiber/app/service"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func PekerjaanRoutes(app *fiber.App, db *sql.DB) {
	pekerjaan := app.Group("/pekerjaan", middleware.AuthRequired())


	// // --- Public / User & Admin ---
	// pekerjaan.Get("/", middleware.AuthRequired(), func(c *fiber.Ctx) error {
	// 	return service.GetAllPekerjaanService(c, db)
	// })

	// pekerjaan.Get("/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
	// 	return service.GetPekerjaanByIDService(c, db)
	// })


    pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
        return service.GetPekerjaanByAlumniIDService(c, db)
    })

	pekerjaan.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.CreatePekerjaanService(c, db)
	})

	pekerjaan.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		return service.UpdatePekerjaanService(c, db)
	})

	pekerjaan.Delete("/:id", func(c *fiber.Ctx) error {
    return service.SoftDeletePekerjaanService(c, db)
	})

	// pekerjaan.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
	// 	return service.DeletePekerjaanService(c, db)
	// })

	pekerjaan.Get("/search", func(c *fiber.Ctx) error {
		return service.GetAllPekerjaanServiceDatatable(c, db)
	})

	pekerjaan.Get("/trash", middleware.AuthRequired(), func(c *fiber.Ctx) error {
		return service.GetTrashedPekerjaanService(c, db)
	})

	pekerjaan.Delete("/hard/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
		return service.HardDeletePekerjaanService(c, db)
	})

	pekerjaan.Put("/restore/:id", middleware.AuthRequired(), func(c *fiber.Ctx) error {
		return service.RestorePekerjaanService(c, db)
	})

}
