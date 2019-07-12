package main

import (
	"XMLparser/db"
	"XMLparser/ivi"
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

func TestCreateTablePersons(t *testing.T) {
	p := ivi.Person{}
	q := db.CreateTable(&p, "persons", "public")
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
	for tag, st := range ivi.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		sql += db.CreateTable(obj, tag, "public") + "\n"
	}
	err := ioutil.WriteFile("testdata/ddl.sql", []byte(sql), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

func TestGenSchema(t *testing.T) {
	oldSchema := "ivi"
	newSchema := "ivi"
	sql := "DROP SCHEMA " + oldSchema + " CASCADE;\nCREATE SCHEMA " + newSchema + ";\n"
	for tag, st := range ivi.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		sql += db.CreateTable(obj, tag, newSchema) + "\n"
	}
	err := ioutil.WriteFile("testdata/create_schema_ivi.sql", []byte(sql), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateSchema(t *testing.T) {
	schema := "ivi"
	newSchema := "new_" + schema
	var tables string
	sql := `CREATE SCHEMA new_@schema;

@tables

BEGIN;
ALTER SCHEMA @schema RENAME TO old_@schema;
ALTER SCHEMA new_@schema RENAME TO @schema;
COMMIT;

DROP SCHEMA old_@schema CASCADE;
`

	// create
	for tag, st := range ivi.Mapping() {
		xType := reflect.TypeOf(st)
		obj := reflect.New(xType).Interface()
		tables += db.CreateTable(obj, tag, newSchema) + "\n"
	}
	// fill
	for tag, _ := range ivi.Mapping() {
		tables += db.CopyTable(tag, newSchema) + "\n"
	}

	sr := strings.NewReplacer("@schema", schema, "@tables", tables)
	sql = sr.Replace(sql)

	err := ioutil.WriteFile("testdata/update_ivi.sql", []byte(sql), os.ModePerm)
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

func TestGenCopyTables(t *testing.T) {
	schema := "ivi"
	sql := ""
	for tag, _ := range ivi.Mapping() {
		sql += db.CopyTable(tag, schema) + "\n"
	}
	err := ioutil.WriteFile("testdata/copy_csv_to_ivi.sql", []byte(sql), os.ModePerm)
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
			{{end}}{{if .IsArray}}{{.Variable}} := "{" + strings.Join(o.{{.Name}},",") + "}"
			{{end}}{{end}}return []string{ {{range .Fields}}{{if .IsSimple}}o.{{.Name}}{{else}}{{.Variable}}{{end}}, {{end}}}
		}
`
	tmpl := template.Must(template.New("").Parse(tpl))

	for tag, st := range ivi.Mapping() {
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
