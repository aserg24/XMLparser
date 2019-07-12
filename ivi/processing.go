package ivi

import "encoding/json"

type CSVLiner interface {
	Line() []string
}

func getJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	json2str := string(b)
	return json2str
}
