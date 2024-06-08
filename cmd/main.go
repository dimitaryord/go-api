package main

import (
	"errors"
	"log"
	"github.com/dimitaryord/go-api/pkg/db"
	"github.com/dimitaryord/go-api/pkg/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)


func main() {
	database := db.Init()
	app := fiber.New()

	app.Post("/person/create", func(c fiber.Ctx) error {
		payload := models.Person{}

		if err := c.Bind().Body(&payload); err != nil {
			log.Fatalln(err)
		}

		log.Println(payload.Name)
		log.Println(payload.Age)

		database.Create(&payload)
		return c.SendString("User created successfully!")	
	})


	app.Get("/person/find", func(c fiber.Ctx) error {
		name := c.Query("name")
		log.Println(name)
		foundPerson := models.Person{}

		database.Where("name = ?", name).First(&foundPerson)
		return c.JSON(foundPerson)
	})
	
	app.Delete("/person/delete", func(c fiber.Ctx) error {
		name := c.Query("name")
		
		result := database.Where("name = ?", name).First(&models.Person{})
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(400).SendString("No user found")
		}

		database.Where("name = ?", name).Delete(&models.Person{})
		return c.SendString("User deleted successfully")
	})

	app.Put("/person/update", func(c fiber.Ctx) error {
		name := c.Query("name")
		result := database.Where("name = ?", name).First(&models.Person{})

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(400).SendString("No user found")
		}

		bodyPerson := models.Person{}

		if err := c.Bind().Body(&bodyPerson); err != nil {
			log.Fatalln(err)	
		}

		database.Model(models.Person{}).Where("name = ?", name).Updates(bodyPerson)
		return c.SendString("User updated successfully")
	})

	app.Listen(":3000")
}
