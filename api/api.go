package api

import (
	"io/ioutil"
	"net/http"
)

const key string = "22e80fd1664c355311fb67cb9de79b87459b2f0b"
const query string = "http://api.guitarparty.com/v2/songs/?query=%s"

func SearchingSongs(titel string) ([]byte, error) {
	netClient := &http.Client{}

	req, _ := http.NewRequest("GET", query, nil)
	req.Header.Set("Guitarparty-Api-Key", key)

	response, err := netClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
