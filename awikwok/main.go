package main

import (
	"awikwok/routes"

	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"Nama":   "Beni",
		"alamat": "Tangsel",
	})
}
func GetNamaAlamat(c *fiber.Ctx) error {
	nama := c.Params("nama")
	alamat := c.Params("alamat")

	return c.JSON(fiber.Map{
		"nama":   nama,
		"alamat": alamat,
	})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/getdata", GetData)
	/*
		app.Get("/getNama", routes.GetNamaAlamat)*/

	app.Get("/getNama/:nama/:alamat", GetNamaAlamat)
	app.Post("/sendData", routes.Senduser)
	app.Listen(":3000")
}
