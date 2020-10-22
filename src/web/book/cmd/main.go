package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Use(middleware)

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"ok":      true,
			"message": "pong",
		})
	})

	app.Listen(":8080")
}

func middleware(ctx iris.Context) {
	ctx.Application().Logger().Info("runs before %s", ctx.Path())
	ctx.Next()
}
