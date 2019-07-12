package ivi

import (
	"XMLparser/db"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

const schema = "ivi"

type IviLoader struct {
	XMLfile   string
	outputDir string
	dbc       *db.DB
}

func NewIviLoader(XMLFile, outputDir string, dbc *db.DB) *IviLoader {
	return &IviLoader{XMLfile: XMLFile, outputDir: outputDir, dbc: dbc}
}

// Parse() unmarshals XML file to Go structures
func (loader IviLoader) Parse() error {
	xmlFile, err := os.Open("ivi/" + loader.XMLfile)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "xml" || se.Name.Local == "action_log_sequence_number" {
				continue
			}

			err := loader.parseElement(decoder, se.Name.Local)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (loader IviLoader) parseElement(decoder *xml.Decoder, tag string) error {
	log.Println(tag)
	tags := loader.Mapping()
	st, ok := tags[tag]
	if !ok {
		return fmt.Errorf("could not find tag = %v", tag)
	}

	filename := "ivi/" + loader.outputDir + "/" + tag + ".csv"
	log.Println(filename)
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	csvWriter.Comma = ';'

	xType := reflect.TypeOf(st)

	for {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			obj := reflect.New(xType).Interface()

			if w, ok := obj.(CSVLiner); !ok {
				return err
			} else if err := decoder.DecodeElement(&obj, &se); err != nil {
				log.Print(err)
				return err
			} else if err := csvWriter.Write(w.Line()); err != nil {
				return err
			}

		case xml.EndElement:
			if se.Name.Local == tag {
				return nil
			}
		}
	}
	return nil
}

// Save() atomically updates an old scheme to the new one with the updated data
func (loader IviLoader) Save() error {
	newSchema := "new_" + schema

	err := loader.dbc.CreateNewSchema(newSchema)
	if err != nil {
		return err
	}
	for tag, st := range loader.Mapping() {

		// create table
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		err = loader.dbc.CreateTable(obj, tag, newSchema)
		if err != nil {
			return err
		}

		// fill table
		err = loader.dbc.FillTable(newSchema, tag, loader.outputDir)
		if err != nil {
			return err
		}
	}

	err = loader.dbc.AlterSchemasNames(schema)
	if err != nil {
		return err
	}

	err = loader.dbc.DropOldSchema(schema)
	if err != nil {
		return err
	}

	return nil
}

// Mapping() matches tables names with structures
func (loader IviLoader) Mapping() map[string]interface{} {
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
