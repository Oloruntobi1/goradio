package radio

import (
	"encoding/json"
	"net/http"
)

type API struct {
	RequestResult StationsRecordAPI
	Endpoint      string
}

func RecordAPI() API {
	return API{
		Endpoint: "https://www.radiorecord.ru/api/stations/",
	}
}

func (api *API) GetJSON() error {
	url := "https://www.radiorecord.ru/api/stations/"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var result StationsRecordAPI
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}

	api.RequestResult = result
	return nil
}
