// package routes

// import (
// 	"database/sql"

// 	"github.com/gofiber/fiber/v2"
// )

// // pastikan ini ada di file routes/alumni.go & routes/pekerjaan.go
// func RegisterRoutes(app *fiber.App, db *sql.DB) {
// 	AlumniRoutes(app, db)
// 	PekerjaanRoutes(app, db)
// 	AuthRoutes(app, db) // kasih db juga
// 	UserRoutes(app, db)
// }

package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"
)

// pastikan ini ada di file routes/alumni.go & routes/pekerjaan.go
func RegisterRoutes(app *fiber.App, db *mongo.Database) {
    AuthRoutes(app, db)
    UserRoutes(app, db)
    AlumniRoutes(app, db)
    PekerjaanRoutes(app, db)
}

