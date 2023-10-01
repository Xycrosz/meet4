package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetNamaAlamat(c *fiber.Ctx) error {
	nama := c.Params("nama")
	alamat := c.Params("alamat")

	return c.JSON(fiber.Map{
		"nama":   nama,
		"alamat": alamat,
	})
}

func Senduser(c *fiber.Ctx) error {

	var data map[string]int

	π := 3.14

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	jaril := data["jari-jari lingkaran"]

	sisip := data["sisi persegi"]
	alastig := data["alas segitiga"]
	tinggitig := data["tinggi segitiga"]
	sisia := data["sisi a"]
	sisib := data["sisi b"]
	sisic := data["sisi c"]
	i := float64(jaril)
	return c.JSON(fiber.Map{
		"luas lingkaran":    i * i * π,
		"luas persegi":      sisip * sisip,
		"luas segitigat":    alastig * tinggitig / 2,
		"keliling lingkata": 2*i*π,
		"keliling persegi":  sisip * 4,
		"keliling segitiga": sisia + sisib + sisic,
		"pesan":             "Sukses untuk lofin lurr",
		"status":            fiber.StatusOK,
	})
}

/*var data map[string]string

if err := c.BodyParser(&data); err != nil {
	return err
}
	ebid := data["username"]
	password := data["password"]

	return c.JSON(fiber.Map{
		"username":     ebid,
		"password": password,
		"pesan": "Sukses untuk lofin lurr",
		"status": fiber.StatusOK,
	})
}*/
