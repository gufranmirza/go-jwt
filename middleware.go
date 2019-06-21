package main

import (
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo"
    "github.com/dgrijalva/jwt-go"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
    SigningKey: []byte("secret"),
})

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        isAdmin := claims["admin"].(bool)
        if isAdmin == false {
            return echo.ErrUnauthorized
        }
        return next(c)
    }
}
