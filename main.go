package main

import (
    "net/http"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    h := &handler{}
    e.POST("/login", h.login)
    e.GET("/is-loggedin", h.private, IsLoggedIn)
    e.GET("/is-admin", h.private, IsLoggedIn, isAdmin)
    e.POST("/refresh", h.token)

    e.Logger.Fatal(e.Start(":1323"))
}
