package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"myShows/structureXML"
	"os"
)

func processing(categories structureXML.ContentCategories) {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{}
	for _, c := range categories.ContentCategory {
		strWrite = append(strWrite, []string{c.ID, c.Title, c.Description, c.Hru, c.Text})
	}
	err = csvWriter.WriteAll(strWrite)
	if err != nil {
		fmt.Print("Can't write content categories")
		os.Exit(1)
	}
	csvWriter.Flush()
}

func main() {
	xmlFile, _ := os.Open("structureXML/document")

	decoder := xml.NewDecoder(xmlFile)

	counter := 0
	const maxRecordsNum = 5

	var contentCategories, contentCategory2s, contentCategory3s structureXML.ContentCategories
	//var contentGenres structureXML.Genres
	//var persons structureXML.Persons
	//var personsType structureXML.PersonsType
	//var awards structureXML.Awards
	//var compilations structureXML.Compilations
	//var countries structureXML.Countries
	//var instruments structureXML.Instruments
	//var moods structureXML.Moods
	//var rusVersions structureXML.RusVersions
	//var contentFormats structureXML.ContentFormats
	//var langs structureXML.Langs
	//var contentTypes structureXML.ContentTypes
	//var properties structureXML.Properties
	//var propertyValues structureXML.PropertyValues
	//var promo structureXML.Promo
	//var additionalDataType structureXML.AdditionalDataType
	//var contents structureXML.Contents
	//var actionLogSequenceNumber structureXML.Text

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		inElem := true
		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "content_categories":
				for inElem {
					t, _ := decoder.Token()
					if t == nil {
						break
					}
					switch innerElem := t.(type) {
					case xml.StartElement:
						if innerElem.Name.Local == "content_category" {
							var parsed structureXML.ContentCategory
							decoder.DecodeElement(&parsed, &innerElem)
							if parsed.ID != "" || parsed.Hru != "" || parsed.Description != "" || parsed.Title != "" {
								contentCategories.ContentCategory = append(contentCategories.ContentCategory, parsed)
								counter++
								if counter >= maxRecordsNum {
									processing(contentCategories)
								}
							} else {
								inElem = false
							}
						}
					case xml.EndElement:
						// fmt.Print("content_categories: ", innerElem.Name.Local)
						if innerElem.Name.Local == "content_categories" {
							inElem = false
						}
					}
				}
			case "content_category_2s":
				for inElem {
					t, _ := decoder.Token()
					if t == nil {
						break
					}
					switch innerElem := t.(type) {
					case xml.StartElement:
						if innerElem.Name.Local == "content_category" {
							var parsed structureXML.ContentCategory
							decoder.DecodeElement(&parsed, &innerElem)
							if parsed.ID != "" || parsed.Hru != "" || parsed.Description != "" || parsed.Title != "" {
								contentCategory2s.ContentCategory = append(contentCategory2s.ContentCategory, parsed)
								counter++
							} else {
								inElem = false
							}
						}
					case xml.EndElement:
						//fmt.Print("content_category_2s: ", innerElem.Name.Local)
						if innerElem.Name.Local == "content_category_2s" {
							inElem = false
						}
					}
				}
			case "content_category_3s":

				for inElem {
					t, _ := decoder.Token()
					if t == nil {
						break
					}
					switch innerElem := t.(type) {
					case xml.StartElement:
						if innerElem.Name.Local == "content_category" {
							var parsed structureXML.ContentCategory
							decoder.DecodeElement(&parsed, &innerElem)
							if parsed.ID != "" || parsed.Hru != "" || parsed.Description != "" ||
								parsed.Title != "" {

								contentCategory3s.ContentCategory = append(contentCategory3s.ContentCategory, parsed)
								counter++
							} else {
								inElem = false
							}
						}
					case xml.EndElement:
						//fmt.Print("content_category_3s: ", innerElem.Name.Local)
						if innerElem.Name.Local == "content_category_3s" {
							inElem = false
						}
					}
				}
			}
			//case "content_category_2s":
			//	fmt.Print(se.Name.Local)
			//	var parsed structureXML.MyShows
			//	decoder.DecodeElement(&parsed.ContentCategory2s, &se)
			//	for i := 0; i < len(parsed.ContentCategory2s.ContentCategory); i++ {
			//		doc.ContentCategory2s.ContentCategory = append(doc.ContentCategory2s.ContentCategory,
			//			parsed.ContentCategory2s.ContentCategory[i])
			//	}
		default:
		}
	}
	for _, cc := range contentCategories.ContentCategory {
		fmt.Printf("Title: %v; ", cc.Title)
		fmt.Printf("Description: %v; ", cc.Description)
		fmt.Printf("Hru: %v; ", cc.Hru)
		fmt.Printf("ID: %v; ", cc.ID)
		fmt.Printf("Text: %v\n", cc.Text)
	}
	fmt.Print("2s: \n")
	for _, cc := range contentCategory2s.ContentCategory {
		fmt.Printf("Title: %v; ", cc.Title)
		fmt.Printf("Description: %v; ", cc.Description)
		fmt.Printf("Hru: %v; ", cc.Hru)
		fmt.Printf("ID: %v; ", cc.ID)
		fmt.Printf("Text: %v\n", cc.Text)
	}
	fmt.Print("3s: \n")
	for _, cc := range contentCategory2s.ContentCategory {
		fmt.Printf("Title: %v; ", cc.Title)
		fmt.Printf("Description: %v; ", cc.Description)
		fmt.Printf("Hru: %v; ", cc.Hru)
		fmt.Printf("ID: %v; ", cc.ID)
		fmt.Printf("Text: %v\n", cc.Text)
	}

}
