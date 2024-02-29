package services

import (
	"log/slog"

	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	Mal_id int `json:"mal_id"`
}

func GetFavorites(c echo.Context) error {
	slog.Info("Getting favorites")
	favs, err := selectAllFavorites()
	if err != nil {
		slog.Error("Error getting favorites", "selectAllFavorites", err)
		return c.String(500, "Error getting favorites")
	}
	return c.JSON(200, favs)
}
func AddToFavorites(c echo.Context) error {
	slog.Info("Adding to favorites")
	var favorite Favorite
	if err := c.Bind(&favorite); err != nil {
		slog.Error("Error parsing request body", "AddToFavorites", err)
		return c.String(400, "Error parsing request body")
	}

	err := addToFavorites(favorite.Mal_id, favorite.Title, favorite.Image)
	if err != nil {
		slog.Error("Error adding to favorites", "addToFavorites", err)
		return c.String(500, "Error adding to favorites")
	}

	return c.JSON(200, "Success")
}

func RemoveFromFavorites(c echo.Context) error {
	requestBody := new(RequestBody)
	if err := c.Bind(requestBody); err != nil {
		slog.Error("Error binding request body", "Bind", err)
		return c.String(400, "Error binding request body")
	}
	id := requestBody.Mal_id
	err := removeFromFavorites(id)
	if err != nil {
		slog.Error("Error deleting from favorites", "removeFromFavorites", err)
		return c.String(500, "Error deleting from favorites")
	}
	return c.JSON(200, "Deleted")
}
