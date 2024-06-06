package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]any

func main() {
	r := echo.New()

	// method ctx.String()
	r.GET("/", func(c echo.Context) error {
		data := "Hello from /index"
		return c.String(http.StatusOK, data)
	})
	// test menggunakan: curl -X GET http://localhost:9000/

	// method ctx.HTML()
	r.GET("/html", func(c echo.Context) error {
		data := "<h1>Hello from~ /html</h1>"
		return c.HTML(http.StatusOK, data)
	})
	// test menggunakan: curl -X GET http://localhost:9000/html

	// method ctx.Redirect
	r.GET("/redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	})
	// test menggunakan: curl -X GET http://localhost:9000/redirect

	// method ctx.JSON()
	r.GET("/json", func(c echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return c.JSON(http.StatusOK, data)
	})
	// test menggunakan: curl -X GET http://localhost:9000/json

	// Parsing queries request
	// decode URL queries or path

	// Parsing Query String
	r.GET("/page1", func(c echo.Context) error {
		name := c.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return c.String(http.StatusOK, data)
	})
	// test menggunakan: curl -X GET http://localhost:9000/page1?name=grayson

	// Parsing URL Path param
	r.GET("/page2/:name", func(c echo.Context) error {
		name := c.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return c.String(http.StatusOK, data)
	})
	// test menggunakan: curl -X GET http://localhost:9000/page2/grayson

	// parsing URL Path param dan setelahnya
	r.GET("/page3/:name/*", func(c echo.Context) error {
		name := c.Param("name")
		message := c.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return c.String(http.StatusOK, data)
	})
	// statement ctx.Param("*") mengembalikan semua path sesuai dengan skema url-nya, Misal url adalah
	// /page3/fadil/a/b/c/d/e/f/g/h maka yang dikembalikan adalah: a/b/c/d/e/f/g/h
	// test menggunakan: curl -X GET http://localhost:9000/page3/tim/need/some/sleep

	// parsing Form Data
	r.GET("/page4", func(c echo.Context) error {
		name := c.FormValue("name")
		message := c.FormValue("message")

		data := fmt.Sprintf("Hello %s, I have message for you: %s",
			name, message)

		return c.String(http.StatusOK, data)
	})
	// test menggunakan: curl -X POST -F name=damian -F message=angry http://localhost:9000/page4

	// penggunaan echo.WrapHandler untuk routing handler bertipe func(http.ResponseWriter, *http.Request)
	// atau http.HandlerFunc
	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	// untuk routing handler dengan skema func(http.ResponseWriter, *http.Request), maka harus dibungkus 2 kali
	// pertama menggunakan http.HandlerFunc, lalu dengan echo.WrapHandler
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	// routing static Asset
	r.Static("/static", "assets")
	// testing langsung di browser dengan URL: http://localhost:9000/static/layout.js

	r.Start(":9000")
}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)

var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)
