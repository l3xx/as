package aviasales

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
func TestClient_GetPlaces(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/places.json?locale=ru&term=%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`[
  {
    "country_cases": null,
    "weight": 1006321,
    "country_code": "RU",
    "main_airport_name": null,
    "cases": {
      "da": "\u041c\u043e\u0441\u043a\u0432\u0435",
      "tv": "\u041c\u043e\u0441\u043a\u0432\u043e\u0439",
      "pr": "\u041c\u043e\u0441\u043a\u0432\u0435",
      "ro": "\u041c\u043e\u0441\u043a\u0432\u044b",
      "vi": "\u0432 \u041c\u043e\u0441\u043a\u0432\u0443"
    },
    "coordinates": {
      "lon": 37.617633,
      "lat": 55.755786
    },
    "name": "\u041c\u043e\u0441\u043a\u0432\u0430",
    "index_strings": [
      "defaultcity",
      "defaultcity",
      "maskava",
      "maskva",
      "mosca",
      "moscou",
      "moscova",
      "moscovo",
      "moscow",
      "mosc\u00fa",
      "moskau",
      "moskou",
      "moskova",
      "moskow",
      "moskva",
      "moskwa",
      "moszkva",
      "\u03bc\u03cc\u03c3\u03c7\u03b1",
      "\u043c\u043e\u0441\u043a\u0432\u0430",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0430\u044f",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0430\u044f",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0441\u043a",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0441\u043a",
      "\u043f\u043e\u043d\u0430\u0435\u0445\u0430\u0432\u0441\u043a",
      "\u043f\u043e\u043d\u0430\u0435\u0445\u0430\u0432\u0441\u043a",
      "\u0574\u0578\u057d\u056f\u057e\u0561",
      "\u05de\u05d5\u05e1\u05e7\u05d1\u05d4",
      "\u0645\u0633\u06a9\u0648",
      "\u0645\u0648\u0633\u0643\u0648",
      "\u092e\u093e\u0938\u094d\u0915\u094b",
      "\u0e21\u0e2d\u0e2a\u0e42\u0e01",
      "\u10db\u10dd\u10e1\u10d9\u10dd\u10d5\u10d8",
      "\u30e2\u30b9\u30af\u30ef",
      "\u83ab\u65af\u79d1",
      "\ubaa8\uc2a4\ud06c\ubc14"
    ],
    "code": "MOW",
    "state_code": null,
    "type": "city",
    "country_name": "\u0420\u043e\u0441\u0441\u0438\u044f"
  },
  {
    "city_cases": {
      "da": "\u041c\u043e\u0441\u043a\u0432\u0435",
      "tv": "\u041c\u043e\u0441\u043a\u0432\u043e\u0439",
      "pr": "\u041c\u043e\u0441\u043a\u0432\u0435",
      "ro": "\u041c\u043e\u0441\u043a\u0432\u044b",
      "vi": "\u0432 \u041c\u043e\u0441\u043a\u0432\u0443"
    },
    "country_cases": null,
    "city_code": "MOW",
    "weight": 613709,
    "cases": null,
    "country_code": "RU",
    "city_name": "\u041c\u043e\u0441\u043a\u0432\u0430",
    "name": "\u0414\u043e\u043c\u043e\u0434\u0435\u0434\u043e\u0432\u043e",
    "coordinates": {
      "lon": 37.899494,
      "lat": 55.414566
    },
    "index_strings": [
      "aeroportul domodedovo moscova",
      "defaultcity",
      "domodedovo",
      "domodedovo internasjonale flyplass",
      "domodedovo luchthaven",
      "domodedowo",
      "domodiedowo",
      "letali\u0161\u010de moskva domodedovo",
      "letisko moskva domodedovo",
      "leti\u0161t\u011b moskva domodedovo",
      "maskava",
      "maskavas domodedovo lidosta",
      "maskva",
      "me\u0111unarodni aerodrom domodedovo",
      "mosca",
      "moscou",
      "moscova",
      "moscovo",
      "moscovo - domodedovo",
      "moscow",
      "moscow domodedovo",
      "moscow domodedovo airport",
      "moscow domodedovos flygplats",
      "mosc\u00fa",
      "moskau",
      "moskou",
      "moskova",
      "moskova domodedovon lentokentt\u00e4",
      "moskow",
      "moskva",
      "moskva-domodedovo internationale lufthavn",
      "moskwa",
      "moszkva",
      "moszkva domodedovo rep\u00fcl\u0151t\u00e9r",
      "zra\u010dna luka moscow domodedovo",
      "\u03bc\u03cc\u03c3\u03c7\u03b1",
      "\u03bc\u03cc\u03c3\u03c7\u03b1 \u03bd\u03c4\u03bf\u03bc\u03bf\u03bd\u03c4\u03ad\u03bd\u03c4\u03bf\u03b2\u03bf \u03b1\u03b5\u03c1\u03bf\u03b4\u03c1\u03cc\u03bc\u03b9\u03bf",
      "\u0434\u043e\u043c\u043e\u0434\u0435\u0434\u043e\u0432\u043e",
      "\u0434\u043e\u043c\u043e\u0434\u0454\u0434\u043e\u0432\u043e",
      "\u043b\u0435\u0442\u0438\u0449\u0435 \u043c\u043e\u0441\u043a\u0432\u0430 \u0434\u043e\u043c\u043e\u0434\u0435\u0434\u043e\u0432\u043e",
      "\u043c\u043e\u0441\u043a\u0432\u0430",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0430\u044f",
      "\u043d\u0435\u0440\u0435\u0437\u0438\u043d\u043e\u0432\u0441\u043a",
      "\u043f\u043e\u043d\u0430\u0435\u0445\u0430\u0432\u0441\u043a",
      "\u0564\u0578\u0574\u0578\u0564\u0565\u0564\u0578\u057e\u0578",
      "\u0574\u0578\u057d\u056f\u057e\u0561",
      "\u05de\u05d5\u05e1\u05e7\u05d1\u05d4",
      "\u05e0\u05de\u05dc \u05d4\u05ea\u05e2\u05d5\u05e4\u05d4 \u05de\u05d5\u05e1\u05e7\u05d1\u05d4 \u05d3\u05d5\u05de\u05d5\u05d3\u05d0\u05d3\u05d5\u05d1\u05d5",
      "\u0641\u0631\u0648\u062f\u06af\u0627\u0647 \u0628\u06cc\u0646\u200c\u0627\u0644\u0645\u0644\u0644\u06cc \u062f\u0648\u0645\u0648\u062f\u0647\u200c\u062f\u0648\u0648",
      "\u0645\u0633\u06a9\u0648",
      "\u0645\u0637\u0627\u0631 \u062f\u0648\u0645\u0648\u062f\u064a\u062f\u0648\u06a4\u0648 \u0645\u0648\u0633\u0643\u0648",
      "\u0645\u0648\u0633\u0643\u0648",
      "\u092e\u093e\u0938\u094d\u0915\u094b",
      "\u092e\u093e\u0938\u094d\u0915\u094b \u0926\u094b\u092e\u094b\u0926\u0947\u0926\u094b\u0935\u094b \u0939\u0935\u093e\u0908 \u0905\u0921\u094d\u0921\u093e",
      "\u0e21\u0e2d\u0e2a\u0e42\u0e01",
      "\u0e2a\u0e19\u0e32\u0e21\u0e1a\u0e34\u0e19\u0e42\u0e14\u0e42\u0e21\u0e40\u0e14\u0e42\u0e14\u0e42\u0e27",
      "\u10d3\u10d0\u10db\u10d0\u10eb\u10d4\u10d3\u10dd\u10d5\u10dd\u10e1 \u10d0\u10d4\u10e0\u10dd\u10de\u10dd\u10e0\u10e2\u10d8",
      "\u10db\u10dd\u10e1\u10d9\u10dd\u10d5\u10d8",
      "\u30e2\u30b9\u30af\u30ef",
      "\u591a\u83ab\u6770\u591a\u6c83\u673a\u573a",
      "\u83ab\u65af\u79d1",
      "\ubaa8\uc2a4\ud06c\ubc14"
    ],
    "code": "DME",
    "state_code": null,
    "type": "airport",
    "country_name": "\u0420\u043e\u0441\u0441\u0438\u044f"
  }
]`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	var cl Client
	val := make(map[string][]string)
	val["locale"] = append(val["locale"], "ru")
	val["term"] = append(val["term"], "Москва")
	cl.Client = client
	_, err := cl.GetPlaces(val)
	assert.Equal(t, nil, err)
}

func TestClient_GetPlacesNotValidJSON(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		assert.Equal(t, "/places.json?locale=ru&term=%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0", req.URL.String())
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`
  {
    "country_name": "\u0420\u043e\u0441\u0441\u0438\u044f"
  }
]`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	var cl Client
	val := make(map[string][]string)
	val["locale"] = append(val["locale"], "ru")
	val["term"] = append(val["term"], "Москва")
	cl.Client = client

	_, err := cl.GetPlaces(val)
	assert.NotNil(t, err)
}