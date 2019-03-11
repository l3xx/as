package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"travel/pkg/aviasales"

	"github.com/magiconair/properties/assert"
)

func TestConvertValid(t *testing.T) {
	jsonFile, err := os.Open("./../../../fixture/response.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var responseAviaSales []aviasales.ResponseAviaSales
	err = json.Unmarshal(byteValue, &responseAviaSales)
	if err != nil {
		fmt.Println(err)
	}
	res := Convert(responseAviaSales)
	assert.Equal(t, len(res), 5)

}

func TestConvertNotValid(t *testing.T) {
	jsonFile, err := os.Open("./../../../fixture/responseWithoutData.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var responseAviaSales []aviasales.ResponseAviaSales
	err = json.Unmarshal(byteValue, &responseAviaSales)
	if err != nil {
		fmt.Println(err)
	}
	res := Convert(responseAviaSales)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].Slug, "MOW")
	assert.Equal(t, res[1].Slug, "DME")

	assert.Equal(t, res[0].Title, "Москва")
	assert.Equal(t, res[1].Title, "Домодедово")

}
