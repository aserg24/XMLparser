package main

import (
	"XMLparser/structureXML"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	xmlFile, err := os.Open("structureXML/document")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	outputDir := "tables"
	err = processFile(xmlFile, outputDir)
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(xmlFile io.Reader, dir string) error {
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

			err := parseXML(decoder, se.Name.Local, dir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseXML(decoder *xml.Decoder, tag, dir string) error {
	log.Println(tag)
	tags := structureXML.Mapping()
	st, ok := tags[tag]
	if !ok {
		return fmt.Errorf("could not find tag=%v", tag)
	}

	filename := dir + "/" + tag + ".csv"
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

			if w, ok := obj.(structureXML.CSVLiner); !ok {
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
