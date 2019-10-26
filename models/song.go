package models

import (
	"chords/api"
	"encoding/json"
	"fmt"
	"strings"
)

type Answer struct {
	Songs []Song `json:"objects"`
	Count int    `json:"objects_count"`
}

type Song struct {
	Body    string   `json:"body"`
	Authors []Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
}

func isExistAuthor(authors []Author, searchingAuthor string) bool {
	searchingAuthor = strings.ToUpper(searchingAuthor)

	for _, author := range authors {
		name := strings.ToUpper(author.Name)
		if strings.Index(name, searchingAuthor) != -1 {
			return true
		}
	}

	return false
}

func GetSong(author, titel string) (string, error) {
	found, err := api.SearchingSongs(titel)
	if err != nil {
		return "", err
	}

	var apiResponse Answer
	err = json.Unmarshal(found, &apiResponse)
	if err != nil {
		return "", err
	}

	if apiResponse.Count == 0 {
		return "", fmt.Errorf("Not found song")
	}

	for _, song := range apiResponse.Songs {
		if isExistAuthor(song.Authors, author) {
			return song.Body, nil
		}
	}

	return "", fmt.Errorf("Not found song")
}
