package controllers

import (
	"os"
	"strconv"
	"strings"

	db "github.com/MuhamadAbsorDwiyana/UjianGolang/Configs"
	models "github.com/MuhamadAbsorDwiyana/UjianGolang/Models"

	"github.com/gofiber/fiber/v2"
)

func ListVisitor(c *fiber.Ctx) error {
	var visitors []models.Visitor

	// limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Offset(skip).Find(&visitors).Count(&count)
	return c.Status(200).Render("dashboard/read", fiber.Map{
		"Title":    "Read Database",
		"AppName":  os.Getenv("APP_NAME"),
		"Visitors": visitors,
	}, "layouts/app")
}

func GetVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor

	if err := db.DB.Where("id = ?", id).First(&visitor).Error; err != nil {
		return c.Status(404).Render("404", fiber.Map{
			"Message": err,
		}, "layouts/app")
	}
	return c.Status(fiber.StatusAccepted).Render("dashboard/get", fiber.Map{
		"Visitor": visitor,
	}, "layouts/app")
}

func AddVisitor(c *fiber.Ctx) error {
	var visitor models.Visitor

	name := c.FormValue("name")
	identity := c.FormValue("identity")
	address := c.FormValue("address")
	ageStr := c.FormValue("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("500", nil, "layouts/app")
	}

	visitor.Name = name
	visitor.Identity = identity
	visitor.Address = address
	visitor.Age = uint8(age)

	// Ambil file gambar profil jika ada
	// file, err := c.FormFile("avatar")
	// if err != nil {
	// 	if err.Error() != "multipart: FormFile: no file" {
	// 		return c.Status(fiber.StatusBadRequest).Render("500", nil, "layouts/app")
	// 	}
	// 	// Jika tidak ada file, jangan set field Avatar
	// } else {
	// 	// Simpan file gambar
	// 	fileName := file.Filename
	// 	savePath := "/Storage/" + fileName

	// 	// Simpan file ke disk
	// 	if err := c.SaveFile(file, savePath); err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).Render("500", nil, "layouts/app")
	// 	}

	// 	// Simpan nama file ke field Avatar
	// 	visitor.Avatar = fileName
	// }

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to get file: " + err.Error())
	}

	// Save the file
	savePath := "./Public/app/storage/" + strings.ReplaceAll(file.Filename, " ", "_")
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save file: " + err.Error())
	}

	visitor.Avatar = strings.ReplaceAll(file.Filename, " ", "_")

	if err := db.DB.Create(&visitor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("500", nil, "layouts/app")
	}

	return c.Redirect("/visitor")
}

func EditVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor

	// Mengambil data visitor berdasarkan ID
	if err := db.DB.First(&visitor, "id = ?", id).Error; err != nil {
		return c.Status(404).Render("404", fiber.Map{
			"Message": "Visitor not found",
		}, "layouts/app")
	}

	// Memperbarui data visitor
	visitor.Name = c.FormValue("name")
	visitor.Identity = c.FormValue("identity")
	visitor.Address = c.FormValue("address")
	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.Status(400).Render("400", fiber.Map{
			"Message": "Invalid age format",
		}, "layouts/app")
	}
	visitor.Age = uint8(age)

	// Menyimpan perubahan data
	if err := db.DB.Save(&visitor).Error; err != nil {
		return c.Status(500).Render("500", fiber.Map{
			"Message": "Failed to update visitor",
		}, "layouts/app")
	}

	return c.Status(200).Redirect("/visitor")
}

func EditVisitorForm(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor

	// Mengambil data visitor berdasarkan ID
	if err := db.DB.First(&visitor, "id = ?", id).Error; err != nil {
		return c.Status(404).Render("404", fiber.Map{
			"Message": "Visitor not found",
		}, "layouts/app")
	}

	// Menampilkan form edit dengan data visitor
	return c.Render("dashboard/update", fiber.Map{
		"Visitor": visitor,
	}, "layouts/app")
}

func DeleteVisitor(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor

	// Mengambil data visitor berdasarkan ID
	if err := db.DB.First(&visitor, "id = ?", id).Error; err != nil {
		return c.Status(404).Render("404", fiber.Map{
			"Message": "Visitor not found",
		}, "layouts/app")
	}

	// Menghapus data visitor
	if err := db.DB.Delete(&visitor).Error; err != nil {
		return c.Status(500).Render("500", fiber.Map{
			"Message": "Failed to delete visitor",
		}, "layouts/app")
	}

	// Redirect ke halaman visitor setelah penghapusan
	return c.Status(204).Redirect("/visitor")
}

func DeleteVisitorConfirmation(c *fiber.Ctx) error {
	id := c.Params("id")
	var visitor models.Visitor

	// Mengambil data visitor berdasarkan ID
	if err := db.DB.First(&visitor, "id = ?", id).Error; err != nil {
		return c.Status(404).Render("404", fiber.Map{
			"Message": "Visitor not found",
		}, "layouts/app")
	}

	// Menampilkan halaman konfirmasi penghapusan
	return c.Render("dashboard/delete", fiber.Map{
		"Visitor": visitor,
	}, "layouts/app")
}
