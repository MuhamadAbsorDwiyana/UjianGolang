package controllers

import (
	db "UjianGolang/Configs"
	models "UjianGolang/Models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListVisitor(c *fiber.Ctx) error {
	var visitors []models.Visitor

	// limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Offset(skip).Find(&visitors).Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "berhasil",
		"data":    visitors,
	})
}

func AddVisitor(c *fiber.Ctx) error {
	var visitor models.Visitor

	name := c.FormValue("name")
	identity := c.FormValue("identity")
	address := c.FormValue("address")
	ageStr := c.FormValue("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "error invalid age value",
		})
	}

	visitor.Name = name
	visitor.Identity = identity
	visitor.Address = address
	visitor.Age = uint8(age)

	if err := db.DB.Create(&visitor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}
