package main

import (
	"XMLparser/structureXML"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
	"text/template"
)

func CreateTable(i interface{}, tableName string) string {
	val := reflect.ValueOf(i).Elem()
	var fields []string
	query := "CREATE TABLE " + tableName + "\n("
	for j := 0; j < val.NumField(); j++ {
		typeField := val.Type().Field(j)
		tag := typeField.Tag

		if tag.Get("pg") == "jsonb" {
			fields = append(fields, "\n\t\""+tag.Get("xml")+"\" jsonb")
		} else if typeField.Type.String() == "[]string" {
			fields = append(fields, "\n\t\""+strings.SplitN(tag.Get("xml"), ",", 2)[0]+"\" text[]")
		} else {
			fields = append(fields, "\n\t\""+tag.Get("xml")+"\" text")
		}
	}
	query += strings.Join(fields, ",") + "\n);"
	return query
}

func CopyTable(tag string) string {
	query := "COPY " + tag + "\n" + `FROM 'D:\proger\go\src\XMLparser\tables\` + tag + `.csv' DELIMITER ';' CSV;`
	return query
}

func TestGenCreateTables(t *testing.T) {
	sql := "DROP SCHEMA public CASCADE;\nCREATE SCHEMA public;\n"
	for tag, st := range structureXML.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		sql += CreateTable(obj, tag) + "\n"
	}
	err := ioutil.WriteFile("testdata/ddl.sql", []byte(sql), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

func TestGenCopyTables(t *testing.T) {
	t.Skip()
	sql := ""
	for tag, _ := range structureXML.Mapping() {
		sql += CopyTable(tag) + "\n"
	}
	err := ioutil.WriteFile("testdata/dml.sql", []byte(sql), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

func TestGenLines(t *testing.T) {
	filename := "structureXML/liners.go"
	file, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	code := "package structureXML\nimport (\n\t\"strings\"\n)\n"
	file.Write([]byte(code))

	tpl := "func (o {{.Struct}}) Line() []string {\n" +
		"{{range .Fields}}{{if .IsJSON}} \t{{.Variable}} := getJSON(o.{{.Name}})\n{{end}}" +
		"{{if .IsArray}}\t{{.Variable}} := \"{\" + strings.Join(o.{{.Name}},\",\") + \"}\"\n{{end}}{{end}}" +
		"\treturn []string{ {{range .Fields}}{{if .IsSimple}}o.{{.Name}}{{else}}{{.Variable}}{{end}}, {{end}}}\n}\n"

	tmpl := template.Must(template.New("").Parse(tpl))

	for tag, st := range structureXML.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		val := reflect.ValueOf(obj).Elem()

		if tag == "content_category_2s" || tag == "content_category_3s" {
			continue
		}

		err := fillStruct(val, tmpl, file)
		if err != nil {
			t.Error(err)
		}
	}
}

type FldType struct {
	IsSimple bool
	IsJSON   bool
	IsArray  bool
	Name     string
	Variable string
}

func fillStruct(val reflect.Value, tmpl *template.Template, file *os.File) error {
	data := struct {
		Struct string
		Fields []FldType
	}{}
	data.Struct = val.Type().Name()
	for j := 0; j < val.NumField(); j++ {
		typeField := val.Type().Field(j)
		var fld FldType
		fld.IsArray, fld.IsJSON, fld.IsSimple = false, false, false
		if typeField.Type.String() == "string" {
			fld.IsSimple = true
		} else if typeField.Type.String() == "[]string" {
			fld.IsArray = true
		} else {
			fld.IsJSON = true
		}
		fld.Name = typeField.Name
		fld.Variable = strings.ToLower(typeField.Name)
		data.Fields = append(data.Fields, fld)
	}
	err := tmpl.Execute(file, data)
	return err
}
