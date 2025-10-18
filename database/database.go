// package database

// import (
//     "database/sql"
//     "log"
//     "os"
//     _ "github.com/lib/pq"
// )

// func ConnectDB() *sql.DB {
//     dsn := os.Getenv("DB_DSN")
//     db, err := sql.Open("postgres", dsn)
//     if err != nil {
//         log.Fatal(err)
//     }

//     if err = db.Ping(); err != nil {
//         log.Fatal("Database tidak connect:", err)
//     }

//     log.Println("DB Connected ✅")
//     return db
// }

package database

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Database {
    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB_NAME")

    if uri == "" {
        uri = "mongodb://localhost:27017"
    }
    if dbName == "" {
        dbName = "alumni_db_tugas4"
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("❌ Gagal konek ke MongoDB:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("❌ MongoDB tidak bisa diakses:", err)
    }

    log.Println("✅ Berhasil konek ke MongoDB:", dbName)
    return client.Database(dbName)
}
