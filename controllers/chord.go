package controllers

import (
	"chords/models"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func GetChords(context echo.Context) (err error) {
	author := context.Param("author")
	author = strings.Replace(author, "+", " ", -1)
	titel := context.Param("titel")

	text, err := models.GetSong(author, titel)

	if err != nil {
		return context.NoContent(http.StatusNotFound)
	}

	return context.JSON(http.StatusOK, models.CreateJSON(text))
}
