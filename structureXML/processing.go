package structureXML

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
)

type CSVLiner interface {
	Line() []string
}

func getJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	json2str := string(b)
	return json2str
}

func ProcessingTag(writer *csv.Writer, strWrite []string) error {
	err := writer.Write(strWrite)
	if err != nil {
		return fmt.Errorf("Can't process tag: %v", err)
	}
	return nil
}

func Mapping() map[string]interface{} {
	return map[string]interface{}{
		"content_categories":   ContentCategory{},
		"content_category_2s":  ContentCategory{},
		"content_category_3s":  ContentCategory{},
		"content_genres":       Genre{},
		"persons":              Person{},
		"persons_type":         PersonType{},
		"awards":               Award{},
		"compilations":         Compilation{},
		"countries":            Country{},
		"instruments":          Instrument{},
		"moods":                Mood{},
		"rus_versions":         RusVersion{},
		"content_formats":      ContentFormat{},
		"langs":                Lang{},
		"content_types":        ContentType{},
		"properties":           Property{},
		"property_values":      PropertyValue{},
		"promo":                PromoInfo{},
		"additional_data_type": ContentAdditionalDataType{},
		"contents":             Content{},
	}
}
