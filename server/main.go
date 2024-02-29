package main

import (
	"log/slog"
	"net/http"
	"os"
	"task/services"
	"task/system"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lmittmann/tint"
)

func main() {
	// Set up the logger
	w := os.Stderr
	var log slog.Level
	if system.LOG_LEVEL == "info" {
		log = slog.LevelInfo
	} else {
		log = slog.LevelDebug
	}
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			AddSource:  true,
			Level:      log,
			TimeFormat: time.Kitchen,
		}),
	))

	// Connect to the database
	err := system.Connect()
	if err != nil {
		slog.Error("Error opening database", "db.Connect", err)
		panic(err)
	}
	slog.Info("Database connected")

	// Run migrations
	err = system.Migrations()
	if err != nil {
		slog.Error("Error running migrations", "db.Migrations", err)
		panic(err)
	}
	slog.Info("Migrations completed")

	// Run the HTTP server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		id := 0
		err := system.Db.QueryRow("SELECT 1").Scan(&id)
		if err != nil {
			slog.Error("Error pinging database", "Db.QueryRow", err)
			return c.String(http.StatusInternalServerError, "Error pinging database")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/favorites", func(c echo.Context) error {
		return services.GetFavorites(c)
	})
	e.POST("/favorites", func(c echo.Context) error {
		return services.AddToFavorites(c)
	})
	e.DELETE("/favorites/:id", func(c echo.Context) error {
		return services.RemoveFromFavorites(c)
	})

	slog.Info("HTTP server listening on", "port", system.HTTP_PORT)
	err = e.Start(":" + system.HTTP_PORT)
	if err != nil {
		slog.Error("Error serving HTTP", "e.Start", err)
		panic(err)
	}
}
