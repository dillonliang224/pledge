package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Use(myMiddleware)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userId, _ := ctx.Params().GetUint64("id")
		// ctx.Writef("userId: %d", userId)
		ctx.JSON(iris.Map{"userId": userId})
	})

	app.Listen(":8080")
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
