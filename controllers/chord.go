package controllers

import (
	"chords/models"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func GetChords(context echo.Context) (err error) {
	author := context.Param("author")
	titel := context.Param("titel")

	text, err := models.GetSong(author, titel)

	models.CreateText(text)

	if err != nil {
		return context.NoContent(http.StatusNotFound)
	}

	return context.JSON(http.StatusOK, strings.TrimSpace(text))
}
