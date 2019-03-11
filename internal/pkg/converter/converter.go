package converter

import (
	"travel/pkg/aviasales"
)

//ResponseConvert ...
type ResponseConvert struct {
	Slug     string `json:"slug"`
	Subtitle string `json:"subtitle"`
	Title    string `json:"title"`
}

//Convert work to convert
func Convert(in []aviasales.ResponseAviaSales) []ResponseConvert {
	var result []ResponseConvert
	for _, v := range in {
		subtitle := v.CityName
		title := v.Name
		if v.Type != "airport" {
			title = v.Name
			subtitle = v.CountryCode
		}
		result = append(result, ResponseConvert{
			Slug:     v.Code,
			Subtitle: subtitle,
			Title:    title,
		})
	}
	return result
}
