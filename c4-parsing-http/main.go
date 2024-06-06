package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"name" query:"name"`
}

func main() {
	r := echo.New()

	// routes
	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})
	// method .Any() menerima segala jenis request dengan method GET, POST, PUT, atau lainnya

	// testing menggunakan Form Data
	// curl -X POST http://localhost:9000/user -d 'name=Fadil' 'email=fadilansyah25.dev@gmail.com'
	// output => {"name":"Fadil", "email":"nope@novalagung.com"}

	// testing menggunakan JSON Payload
	// curl -X POST http://localhost:9000/user -H 'Content-Type: application/json' -d '{"name":"Fadil", "email":"fadilansyah25.dev@gmail.com"}'
	// output => {"name":"Fadil", "email":"fadilansyah25.dev@gmail.com"}

	// testing menggunakan XML Payload
	// curl -X POST http://localhost:9000/user -H 'Content-Type: application/xml' -d '<?xml version="1.0"?><Data><Name>Fadil</Name><Email>fadilansyah25.dev@gmail.com</Email></Data>'
	// output => {"name":"Fadil", "email":"nope@novalagung.com"}

	// testng menggunakan query string
	// curl -X GET http://localhost:9000/user?name=Joe&email=nope@novalagung.com

	fmt.Println("Server started at localhost:9000")
	r.Start(":9000")
}
