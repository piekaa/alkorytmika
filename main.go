package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	s := &Server{}
	// Routes
	e.GET("/Test", s.Test)
	e.Static("/", "static")
	// Start server
	e.Logger.Fatal(e.Start(":8882"))
}

type Server struct {
}


func (s *Server) Test(c echo.Context) error {

	return c.String(http.StatusOK, "Test")
}