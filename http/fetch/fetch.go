package fetch

import (
	"encoding/json"
	"net/http"
)

var myClient = &http.Client{}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}