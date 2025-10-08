package service

import (
	"database/sql"
	"strconv"
	"math"
	"github.com/gofiber/fiber/v2"
	"go-fiber/app/model"
	"go-fiber/app/repository"


)

// GET ALL PEKERJAAN
// func GetAllPekerjaanService(c *fiber.Ctx, db *sql.DB) error {
// 	pekerjaanList, err := repository.FindAllPekerjaan(db)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal mendapatkan data pekerjaan: " + err.Error(),
// 			"success": false,
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"message": "Berhasil mendapatkan semua data pekerjaan",
// 		"success": true,
// 		"data":    pekerjaanList,
// 	})
// }

// GET PEKERJAAN BY ID
// func GetPekerjaanByIDService(c *fiber.Ctx, db *sql.DB) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "ID tidak valid",
// 			"success": false,
// 		})
// 	}

// 	pekerjaan, err := repository.FindPekerjaanByID(db, id)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return c.Status(404).JSON(fiber.Map{
// 				"message": "Pekerjaan tidak ditemukan",
// 				"success": false,
// 			})
// 		}
// 		return c.Status(500).JSON(fiber.Map{
// 			"message": "Gagal mendapatkan pekerjaan: " + err.Error(),
// 			"success": false,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Berhasil mendapatkan data pekerjaan",
// 		"success": true,
// 		"data":    pekerjaan,
// 	})
// }

// GET PEKERJAAN BY ALUMNI ID
func GetPekerjaanByAlumniIDService(c *fiber.Ctx, db *sql.DB) error {
	alumniID, err := strconv.Atoi(c.Params("alumni_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID alumni tidak valid",
			"success": false,
		})
	}

	pekerjaanList, err := repository.FindPekerjaanByAlumniID(db, alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mendapatkan data pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mendapatkan data pekerjaan untuk alumni",
		"success": true,
		"data":    pekerjaanList,
	})
}

// CREATE PEKERJAAN
func CreatePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	var pekerjaan model.Pekerjaan
	if err := c.BodyParser(&pekerjaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Input tidak valid: " + err.Error(),
			"success": false,
		})
	}

	newPekerjaan, err := repository.CreatePekerjaan(db, pekerjaan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menambahkan pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Pekerjaan berhasil ditambahkan",
		"success": true,
		"data":    newPekerjaan,
	})
}

// UPDATE PEKERJAAN
func UpdatePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
			"success": false,
		})
	}

	var pekerjaan model.Pekerjaan
	if err := c.BodyParser(&pekerjaan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Input tidak valid: " + err.Error(),
			"success": false,
		})
	}

	updatedPekerjaan, err := repository.UpdatePekerjaan(db, id, pekerjaan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal update pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pekerjaan berhasil diupdate",
		"success": true,
		"data":    updatedPekerjaan,
	})
}

// DELETE PEKERJAAN
func DeletePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID tidak valid",
			"success": false,
		})
	}

	if err := repository.DeletePekerjaan(db, id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pekerjaan berhasil dihapus",
		"success": true,
	})
}

func GetAllPekerjaanServiceDatatable(c *fiber.Ctx, db *sql.DB) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	list, err := repository.GetPekerjaanRepo(db, search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mendapatkan data pekerjaan alumni: " + err.Error(),
			"success": false,
		})
	}

	total, err := repository.CountPekerjaanRepo(db, search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghitung total pekerjaan alumni: " + err.Error(),
			"success": false,
		})
	}

	meta := model.MetaInfo{
		Page:   page,
		Limit:  limit,
		Total:  total,
		Pages:  int(math.Ceil(float64(total) / float64(limit))),
		SortBy: sortBy,
		Order:  order,
		Search: search,
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mendapatkan data pekerjaan alumni",
		"success": true,
		"data":    list,
		"meta":    meta,
	})
}

func SoftDeletePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error":   "ID tidak valid",
            "success": false,
        })
    }

    userID := c.Locals("user_id").(int)
    role := c.Locals("role").(string)
    isAdmin := role == "admin"

    repo := repository.NewPekerjaanRepository(db)
    err = repo.SoftDelete(id, userID, isAdmin)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(403).JSON(fiber.Map{
                "error":   "Tidak bisa hapus pekerjaan ini",
                "success": false,
            })
        }
        return c.Status(500).JSON(fiber.Map{
            "error":   "Gagal soft delete pekerjaan: " + err.Error(),
            "success": false,
        })
    }

    return c.JSON(fiber.Map{
        "message": "Pekerjaan berhasil dihapus (soft delete)",
        "success": true,
    })
}

func GetTrashedPekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	userID := c.Locals("user_id").(int)
	role := c.Locals("role").(string)
	isAdmin := role == "admin"

	repo := repository.NewPekerjaanRepository(db)
	data, err := repo.GetTrashed(userID, isAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Gagal mengambil data trash: " + err.Error(),
			"success": false,
		})
	}

	if len(data) == 0 {
		return c.JSON(fiber.Map{
			"message": "Tidak ada data pekerjaan yang dihapus",
			"success": true,
			"data":    []model.Pekerjaan{},
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data trash",
		"success": true,
		"data":    data,
	})
}

func RestorePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "ID tidak valid",
			"success": false,
		})
	}

	userID := c.Locals("user_id").(int)
	role := c.Locals("role").(string)
	isAdmin := role == "admin"

	repo := repository.NewPekerjaanRepository(db)
	if err := repo.Restore(id, userID, isAdmin); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Gagal restore pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pekerjaan berhasil direstore",
		"success": true,
	})
}

func HardDeletePekerjaanService(c *fiber.Ctx, db *sql.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   "ID tidak valid",
			"success": false,
		})
	}

	userID := c.Locals("user_id").(int)
	role := c.Locals("role").(string)
	isAdmin := role == "admin"

	repo := repository.NewPekerjaanRepository(db)
	if err := repo.HardDelete(id, userID, isAdmin); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Gagal hard delete pekerjaan: " + err.Error(),
			"success": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pekerjaan berhasil dihapus permanen (hard delete)",
		"success": true,
	})
}
