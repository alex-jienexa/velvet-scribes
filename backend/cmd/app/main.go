package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Создаем новое приложение Fiber
	app := fiber.New()

	// Определяем маршрут для GET запроса на корневой путь
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Привет, мир! 👋")
	})

	// Запускаем сервер на порту 3000
	app.Listen(":3000")
}
