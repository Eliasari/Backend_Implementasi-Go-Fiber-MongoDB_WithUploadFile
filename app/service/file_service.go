package service

import (
	"context"
	"fmt"
	"go-fiber/app/model"
	"go-fiber/app/repository"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadFileService(repo *repository.FileRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(string)

		fileType := c.FormValue("type")
		targetUserID := c.FormValue("user_id")

		// admin bisa upload untuk user lain, user hanya bisa untuk dirinya sendiri
		if role != "admin" {
			targetUserID = userID
		} else if targetUserID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Admin harus menyertakan user_id"})
		}

		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "File tidak ditemukan"})
		}

		// validasi file
		if err := validateFile(fileHeader, fileType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		// simpan ke /uploads/{type}/
		saveDir := fmt.Sprintf("./uploads/%s", fileType)
		os.MkdirAll(saveDir, os.ModePerm)

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		filePath := filepath.Join(saveDir, filename)

		if err := c.SaveFile(fileHeader, filePath); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file"})
		}

		userObjID, _ := primitive.ObjectIDFromHex(targetUserID)
		file := model.File{
			ID:          primitive.NewObjectID(),
			UserID:      userObjID,
			Filename:    fileHeader.Filename,
			Path:        filePath,
			ContentType: fileHeader.Header.Get("Content-Type"),
			Size:        fileHeader.Size,
			Type:        fileType,
			CreatedAt:   time.Now(),
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := repo.CreateFile(ctx, &file); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan ke database"})
		}

		return c.JSON(fiber.Map{
			"message": "File berhasil diupload",
			"data":    file,
		})
	}
}

func GetFilesService(repo *repository.FileRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(string)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		files, err := repo.GetFiles(ctx, userID, role)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil file"})
		}

		return c.JSON(fiber.Map{"data": files})
	}
}

func GetFileByIDService(repo *repository.FileRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		file, err := repo.GetFileByID(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "File tidak ditemukan",
			})
		}

		// Ambil info user dari JWT
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(string)

		// Convert userID JWT ke ObjectID biar bisa dibandingin
		currentUserObjID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID tidak valid",
			})
		}

		// Kalau bukan admin, cek apakah file milik user itu
		if role != "admin" && file.UserID != currentUserObjID {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Kamu tidak punya akses ke file ini",
			})
		}

		return c.JSON(fiber.Map{
			"data": file,
		})
	}
}


func DeleteFileService(repo *repository.FileRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Ambil data file dari DB
		file, err := repo.GetFileByID(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "File tidak ditemukan",
			})
		}

		// Ambil info user dari JWT
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(string)
		currentUserObjID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID tidak valid",
			})
		}

		// Debug log (sementara)
		fmt.Println("ROLE:", role)
		fmt.Println("File.UserID:", file.UserID.Hex())
		fmt.Println("CurrentUser:", currentUserObjID.Hex())

		// Authorization
		if role != "admin" {
			// Kalau user bukan pemilik file → tolak
			if file.UserID != currentUserObjID {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": "Kamu tidak bisa menghapus file milik user lain",
				})
			}
		}

		// Hapus file di sistem (opsional, kalau ada file fisik)
		if err := os.Remove(file.Path); err != nil {
			fmt.Println("Warning: gagal hapus file fisik:", err)
		}

		// Hapus dari database
		if err := repo.DeleteFileByID(ctx, id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menghapus file",
			})
		}

		return c.JSON(fiber.Map{
			"message": "File berhasil dihapus",
		})
	}
}


// --- helper validation
func validateFile(file *multipart.FileHeader, fileType string) error {
	contentType := file.Header.Get("Content-Type")
	size := file.Size

	if fileType == "photo" {
		if !(strings.Contains(contentType, "jpeg") || strings.Contains(contentType, "jpg") || strings.Contains(contentType, "png")) {
			return fmt.Errorf("foto harus berformat jpeg/jpg/png")
		}
		if size > 1*1024*1024 {
			return fmt.Errorf("ukuran foto maksimal 1MB")
		}
	}

	if fileType == "certificate" {
		if !strings.Contains(contentType, "pdf") {
			return fmt.Errorf("sertifikat harus berformat PDF")
		}
		if size > 2*1024*1024 {
			return fmt.Errorf("ukuran sertifikat maksimal 2MB")
		}
	}

	return nil
}
