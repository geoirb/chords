package controllers

import (
	"chords/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetChords(context echo.Context) (err error) {
	author := context.Param("author")
	titel := context.Param("titel")

	text, err := models.GetSong(author, titel)

	if err != nil {
		return context.NoContent(http.StatusNotFound)
	}

	return context.JSON(http.StatusOK, text)
}
