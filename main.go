// package main

// import (
//     "log"
//     "os"
//     "go-fiber/config"
//     "go-fiber/database"
//     "go-fiber/routes"
// )

// func main() {
//     config.LoadEnv()
//     db := database.ConnectDB()
//     defer db.Close()

//     app := config.NewApp(db)

//     // register semua route
//     routes.RegisterRoutes(app, db)

//     port := os.Getenv("APP_PORT")
//     if port == "" {
//         port = "3000"
//     }

//     log.Fatal(app.Listen(":" + port))
// }

package main

import (
    "log"
    "os"
    "go-fiber/config"
    "go-fiber/database"
)

func main() {
    config.LoadEnv()

    // ✅ Panggil ConnectMongo()
    db := database.ConnectMongo()
    if db == nil {
        log.Fatal("❌ Gagal konek ke MongoDB — return nil")
    }

    // ✅ Kirim db ke app
    app := config.NewApp(db)

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}

