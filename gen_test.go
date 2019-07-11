package main

import (
	"XMLparser/structureXML"
	"bytes"
	"go/format"
	"io"
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

		var sqlType string
		if tag.Get("pg") == "jsonb" {
			sqlType = "jsonb"
		} else if typeField.Type.String() == "[]string" {
			sqlType = "text[]"
		} else {
			sqlType = "text"
		}
		fields = append(fields, "\n\t\""+strings.SplitN(tag.Get("xml"), ">", 2)[0]+"\" "+sqlType)
	}
	query += strings.Join(fields, ",") + "\n);"
	return query
}

func TestCreateTablePersons(t *testing.T) {
	p := structureXML.Person{}
	q := CreateTable(&p, "persons")
	expected := `CREATE TABLE persons
(
	id text,
	posters jsonb,
	name text,
	is_star text,
	eng_title text,
	gen_title text,
	hru text,
	fake_hru text,
	blog_post text,
	blog_link text,
	blog_added text,
	bio text,
	facts text,
	rss text,
	kinopoisk_id text
);`
	if q != expected {
		t.Errorf("Table 'persons' was created incorrectly, got: \n%s \nwant: \n%s ", q, expected)
	}
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

// text/template
// https://golang.org/pkg/text/template/#example_Template

// rewrite sql => schema ivi_new
// create schema ivi;
// create table ivi.countries...

// new file with copy from

// begin; rename ivi to ivi_old; rename ivi_new to ivi; drop ivi_old; commit;

func CopyTable(tag string) string {
	query := "COPY " + tag + "\n" + `FROM 'D:\proger\go\src\XMLparser\tables\` + tag + `.csv' DELIMITER ';' CSV;`
	return query
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
	var buf bytes.Buffer

	code := "package structureXML\nimport (\n\t\"strings\"\n)\n"
	buf.Write([]byte(code))

	tpl := `
		func (o {{.Struct}}) Line() []string { 
			{{range .Fields}}{{if .IsJSON}} {{.Variable}} := getJSON(o.{{.Name}})
			{{end}}{{if .IsArray}}{{.Variable}} := strings.Join(o.{{.Name}},",")
			{{end}}{{end}}return []string{ {{range .Fields}}{{if .IsSimple}}o.{{.Name}}{{else}}{{.Variable}}{{end}}, {{end}}}
		}
`
	tmpl := template.Must(template.New("").Parse(tpl))

	for tag, st := range structureXML.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		val := reflect.ValueOf(obj).Elem()

		if tag == "content_category_2s" || tag == "content_category_3s" {
			continue
		}

		err := writeTemplate(val, tmpl, &buf)
		if err != nil {
			t.Error(err)
		}
	}

	fmtCode, err := format.Source(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	filename := "structureXML/liners.go"
	err = ioutil.WriteFile(filename, fmtCode, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

type FldType struct {
	IsSimple bool
	IsJSON   bool
	IsArray  bool
	Name     string
	Variable string
}

func writeTemplate(val reflect.Value, tmpl *template.Template, file io.Writer) error {
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
