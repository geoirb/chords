package controllers

import (
	"fmt"

	"github.com/labstack/echo"
)

func GetChords(context echo.Context) (err error) {
	author := context.Param("author")
	titel := context.Param("titel")

	fmt.Println(author, titel)

	return nil
}
