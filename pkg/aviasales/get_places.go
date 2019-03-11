package aviasales

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

//ResponseAviaSales ...
type ResponseAviaSales struct {
	StateCode       interface{} `json:"state_code"`
	MainAirportName interface{} `json:"main_airport_name,omitempty"`
	Name            string      `json:"name"`
	IndexStrings    []string    `json:"index_strings"`
	Type            string      `json:"type"`
	CountryCases    interface{} `json:"country_cases"`
	Coordinates     Coordinates `json:"coordinates"`
	CountryCode     string      `json:"country_code"`
	Code            string      `json:"code"`
	CountryName     string      `json:"country_name"`
	Weight          int         `json:"weight"`
	Cases           Cases       `json:"cases"`
	CityCode        string      `json:"city_code,omitempty"`
	CityCases       CityCases   `json:"city_cases,omitempty"`
	CityName        string      `json:"city_name,omitempty"`
}

//Cases ...
type Cases struct {
	Pr string `json:"pr"`
	Tv string `json:"tv"`
	Ro string `json:"ro"`
	Da string `json:"da"`
	Vi string `json:"vi"`
}

//CityCases ...
type CityCases struct {
	Pr string `json:"pr"`
	Tv string `json:"tv"`
	Ro string `json:"ro"`
	Da string `json:"da"`
	Vi string `json:"vi"`
}

//Coordinates ...
type Coordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

//todo may be struct?
//GetPlaces with validation
//type GetPlaces struct {
//	Term   string   `validate:"required,min=2"`
//	Locale string   `validate:"required"`
//	Types  []string `validate:"required"`
//}

//GetPlaces work
func (c *Client) GetPlaces(query url.Values) ([]ResponseAviaSales, error) {

	u, err := url.Parse(fmt.Sprintf("%s/places.json", c.URL))
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	var result []ResponseAviaSales
	response, err := c.Client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
