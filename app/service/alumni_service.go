// package service

// import (
// 	"database/sql"
// 	"strconv"
// 	"math"

// 	"github.com/gofiber/fiber/v2"
// 	"go-fiber/app/model"
// 	"go-fiber/app/repository"
// )

// // GET ALL ALUMNI
// // func GetAllAlumniService(c *fiber.Ctx, db *sql.DB) error {
// // 	alumniList, err := repository.FindAllAlumni(db)
// // 	if err != nil {
// // 		return c.Status(500).JSON(fiber.Map{
// // 			"message": "Gagal mendapatkan data alumni: " + err.Error(),
// // 			"success": false,
// // 		})
// // 	}

// // 	return c.JSON(fiber.Map{
// // 		"message": "Berhasil mendapatkan semua data alumni",
// // 		"success": true,
// // 		"data":    alumniList,
// // 	})
// // }

// // GET ALUMNI BY ID
// // func GetAlumniByIDService(c *fiber.Ctx, db *sql.DB) error {
// // 	id, err := strconv.Atoi(c.Params("id"))
// // 	if err != nil {
// // 		return c.Status(400).JSON(fiber.Map{
// // 			"message": "ID tidak valid",
// // 			"success": false,
// // 		})
// // 	}

// // 	alumni, err := repository.FindAlumniByID(db, id)
// // 	if err != nil {
// // 		if err == sql.ErrNoRows {
// // 			return c.Status(404).JSON(fiber.Map{
// // 				"message": "Alumni tidak ditemukan",
// // 				"success": false,
// // 			})
// // 		}
// // 		return c.Status(500).JSON(fiber.Map{
// // 			"message": "Gagal mendapatkan alumni: " + err.Error(),
// // 			"success": false,
// // 		})
// // 	}

// // 	return c.JSON(fiber.Map{
// // 		"message": "Berhasil mendapatkan data alumni",
// // 		"success": true,
// // 		"data":    alumni,
// // 	})
// // }

// // CREATE ALUMNI
// func CreateAlumniService(c *fiber.Ctx, db *sql.DB) error {
// 	var alumni model.Alumni
// 	if err := c.BodyParser(&alumni); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "Input tidak valid: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	newAlumni, err := repository.CreateAlumni(db, alumni)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal menambahkan alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	return c.Status(201).JSON(fiber.Map{
// 		"message": "Alumni berhasil ditambahkan",
// 		"success": true,
// 		"data":    newAlumni,
// 	})
// }

// // UPDATE ALUMNI
// func UpdateAlumniService(c *fiber.Ctx, db *sql.DB) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "ID tidak valid",
// 			"success": false,
// 		})
// 	}

// 	var alumni model.Alumni
// 	if err := c.BodyParser(&alumni); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "Input tidak valid: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	updatedAlumni, err := repository.UpdateAlumni(db, id, alumni)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal update alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Alumni berhasil diupdate",
// 		"success": true,
// 		"data":    updatedAlumni,
// 	})
// }

// // DELETE ALUMNI
// func DeleteAlumniService(c *fiber.Ctx, db *sql.DB) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "ID tidak valid",
// 			"success": false,
// 		})
// 	}

// 	if err := repository.DeleteAlumni(db, id); err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal menghapus alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Alumni berhasil dihapus",
// 		"success": true,
// 	})
// }

// func GetAllAlumniServiceDatatable(c *fiber.Ctx, db *sql.DB) error {
// 	// Ambil query params
// 	page, _ := strconv.Atoi(c.Query("page", "1"))
// 	limit, _ := strconv.Atoi(c.Query("limit", "10"))
// 	sortBy := c.Query("sortBy", "id")
// 	order := c.Query("order", "asc")
// 	search := c.Query("search", "")

// 	if page < 1 {
// 		page = 1
// 	}
// 	offset := (page - 1) * limit

// 	// Ambil data alumni
// 	alumniList, err := repository.GetAlumniRepo(db, search, sortBy, order, limit, offset)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal mendapatkan data alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	// Ambil total data
// 	total, err := repository.CountAlumniRepo(db, search)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal menghitung total alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	meta := model.MetaInfo{
// 		Page:   page,
// 		Limit:  limit,
// 		Total:  total,
// 		Pages:  int(math.Ceil(float64(total) / float64(limit))),
// 		SortBy: sortBy,
// 		Order:  order,
// 		Search: search,
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Berhasil mendapatkan data alumni",
// 		"success": true,
// 		"data":    alumniList,
// 		"meta":    meta,
// 	})
// }

// func CountAlumniByStatusService(c *fiber.Ctx, db *sql.DB) error {
// 	result, err := repository.CountAlumniByStatusRepo(db)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal menghitung status alumni: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Berhasil menghitung status alumni",
// 		"success": true,
// 		"data":    result,
// 	})
// }

package service

import (
	"context"
	"math"
	"strconv"
	"time"

	"go-fiber/app/model"
	"go-fiber/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive" 

	"github.com/gofiber/fiber/v2"
)

func CreateAlumniService(c *fiber.Ctx, repo *repository.AlumniRepository) error {
	var alumni model.Alumni
	if err := c.BodyParser(&alumni); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Input tidak valid: " + err.Error(),
		})
	}

	// Ambil user_id dari JWT (disimpan di c.Locals)
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok || userIDStr == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "User ID tidak valid di token",
		})
	}

	// Convert string ke ObjectID
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User ID bukan ObjectID valid",
		})
	}

	// Set waktu dan user_id
	alumni.UserID = userID
	alumni.CreatedAt = time.Now()
	alumni.UpdatedAt = time.Now()

	// Simpan ke Mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newAlumni, err := repo.CreateAlumni(ctx, &alumni)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menambahkan alumni: " + err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil ditambahkan",
		"data":    newAlumni,
	})
}

func UpdateAlumniService(c *fiber.Ctx, repo *repository.AlumniRepository) error {
	id := c.Params("id")

	var update model.UpdateAlumni
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Input tidak valid: " + err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updated, err := repo.UpdateAlumni(ctx, id, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal update alumni: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil diupdate",
		"data":    updated,
	})
}


func DeleteAlumniService(c *fiber.Ctx, repo *repository.AlumniRepository) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.DeleteAlumni(ctx, id); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal hapus alumni: " + err.Error(), "success": false})
	}

	return c.JSON(fiber.Map{"message": "Alumni berhasil dihapus", "success": true})
}

func GetAllAlumniService(c *fiber.Ctx, repo *repository.AlumniRepository) error {
	page, _ := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
	search := c.Query("search", "")
	offset := (page - 1) * limit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	alumni, err := repo.GetAlumni(ctx, search, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal fetch data: " + err.Error(), "success": false})
	}

	total, _ := repo.CountAlumni(ctx, search)

	meta := model.MetaInfo{
		Page:   int(page),
		Limit:  int(limit),
		Total:  int(total),
		Pages:  int(math.Ceil(float64(total) / float64(limit))),
		SortBy: "nama",
		Order:  "asc",
		Search: search,
	}

	return c.JSON(model.AlumniResponse{
		Data: alumni,
		Meta: meta,
	})
}
