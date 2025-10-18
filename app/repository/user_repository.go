// versi sebelum menggunakan mongoDB

// package repository

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"go-fiber/app/model"
// )

// func FindUserByUsernameOrEmail(db *sql.DB, identifier string) (*model.User, string, error) {
// 	var user model.User
// 	var passwordHash string

// 	err := db.QueryRow(`
// 		SELECT id, username, email, password_hash, role, created_at
// 		FROM users WHERE username = $1 OR email = $1
// 	`, identifier).Scan(
// 		&user.ID, &user.Username, &user.Email, &passwordHash, &user.Role, &user.CreatedAt,
// 	)

// 	if err != nil {
// 		return nil, "", err
// 	}
// 	return &user, passwordHash, nil
// }

// func GetUsersRepo(db *sql.DB, search, sortBy, order string, limit, offset int) ([]model.User, error) {
// 	// Pastikan sortBy & order udah divalidasi di service

// 	query := fmt.Sprintf(`
// 		SELECT id, username, email, password_hash, role, created_at
// 		FROM users
// 		WHERE username ILIKE $1 OR email ILIKE $1
// 		ORDER BY %s %s
// 		LIMIT $2 OFFSET $3
// 	`, sortBy, order)

// 	log.Println("SQL:", query)
// 	log.Println("Params:", "%"+search+"%", limit, offset)

// 	rows, err := db.Query(query, "%"+search+"%", limit, offset)
// 	if err != nil {
// 		log.Println("Query error:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []model.User

// 	for rows.Next() {
//     var u model.User
//     var passwordHash string

//     if err := rows.Scan(&u.ID, &u.Username, &u.Email, &passwordHash, &u.Role, &u.CreatedAt); err != nil {
//         return nil, err
//     }
//     users = append(users, u)
// }

// return users, nil
// }

// func CountUsersRepo(db *sql.DB, search string) (int, error) {
// 	var total int
// 	query := `SELECT COUNT(*) FROM users WHERE username ILIKE $1 OR email ILIKE $1`
// 	err := db.QueryRow(query, "%"+search+"%").Scan(&total)
// 	if err != nil && err != sql.ErrNoRows {
// 		return 0, err
// 	}
// 	return total, nil
// }

// versi setelah menggunakan mongoDB

package repository

import (
    "context"
    "go-fiber/app/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
    collection *mongo.Collection
}


func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        collection: db.Collection("users"),
    }
}

func (r *UserRepository) FindUserByUsernameOrEmail(ctx context.Context, identifier string) (*model.User, error) {
    var user model.User
    filter := bson.M{"$or": []bson.M{{"username": identifier}, {"email": identifier}}}
    err := r.collection.FindOne(ctx, filter).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context, search string, limit, offset int64) ([]model.User, error) {
    filter := bson.M{
        "$or": []bson.M{
            {"username": bson.M{"$regex": search, "$options": "i"}},
            {"email": bson.M{"$regex": search, "$options": "i"}},
        },
    }

    opts := &options.FindOptions{
        Skip:  &offset,
        Limit: &limit,
    }

    cursor, err := r.collection.Find(ctx, filter, opts)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []model.User
    if err := cursor.All(ctx, &users); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *UserRepository) CountUsers(ctx context.Context, search string) (int64, error) {
    filter := bson.M{
        "$or": []bson.M{
            {"username": bson.M{"$regex": search, "$options": "i"}},
            {"email": bson.M{"$regex": search, "$options": "i"}},
        },
    }

    count, err := r.collection.CountDocuments(ctx, filter)
    if err != nil {
        return 0, err
    }

    return count, nil
}