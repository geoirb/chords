package controllers

import (
	"chords/models"
	"fmt"

	"github.com/labstack/echo"
)

func GetChords(context echo.Context) (err error) {
	author := context.Param("author")
	titel := context.Param("titel")

	fmt.Println(models.GetSong(author, titel))
	return nil
}
