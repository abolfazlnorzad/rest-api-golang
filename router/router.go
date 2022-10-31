package router

import (
	"app/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateRouter(db *database.DBHandler) error {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		fmt.Println("done")

		fmt.Println(db.GetAllUsers())
		return ctx.JSON(db.GetAllUsers())
	})

	app.Get("/user/:name", func(ctx *fiber.Ctx) error {
		return ctx.JSON(db.FindUserByName(ctx.Params("name")))
	})

	app.Get("/:id", func(ctx *fiber.Ctx) error {
		fmt.Println("s")
		num, _ := ctx.ParamsInt("id")
		fmt.Println(num)
		return ctx.JSON(db.UpdateUserById(num))
	})

	app.Get("delete/:id", func(ctx *fiber.Ctx) error {
		fmt.Println("s")
		num, _ := ctx.ParamsInt("id")
		fmt.Println(num)
		db.DeleteUserById(num)
		return ctx.SendString("done")
	})
	return app.Listen(":3000")
}
