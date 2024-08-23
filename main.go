package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Method: GET
	// Resource http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<html><body style='background-color:white;'><h1 style='color:black;'>Welcome to Sample Go Web App</h1></body></html>")
	})

	// same as app.Handle("GET", "/ping", ..)
	// Method: GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method: GET
	// Resource: http://localhost:8080/mypath
	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from %s", ctx.Path())
	})

	// Method: GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello World!"})
	})

	// Any custom fields here. Handler and ErrorLog are set to the server automatically
	srv := &http.Server{Addr: ":8080"}

	// http://localhost:8080/
	// http://localhost:8080/mypath
	app.Run(iris.Server(srv)) // same as app.Listen(":8080")

}
