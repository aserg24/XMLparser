package db

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/go-pg/pg"
)

func CreateTable(i interface{}, tableName string, schema string) string {
	val := reflect.ValueOf(i).Elem()
	var fields []string
	query := "CREATE TABLE " + schema + "." + tableName + "\n("
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

func CopyTable(table string, schema string) string {
	query := "COPY " + schema + "." + table + "\n" + `FROM 'D:\proger\go\src\XMLparser\tables\` + table + `.csv' DELIMITER ';' CSV;`
	return query
}

type DB struct {
	*pg.DB
}

// New is a function that returns DB as wrapper on PostgreSQL connection.
func New(db *pg.DB) DB {
	return DB{db}
}

func CSVreader(table, dir string) (io.Reader, error) {
	r, err := os.Open("ivi/" + dir + "/" + table + ".csv")
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return r, nil
}

// CreateNewSchema just creates a new schema
func (db DB) CreateNewSchema(schema string) error {
	sql := fmt.Sprintf("CREATE SCHEMA %s;", schema)
	_, err := db.Exec(sql)
	return err
}

// CreateTable creates schema.table with fields like obj's ones
func (db DB) CreateTable(obj interface{}, table, schema string) error {
	_, err := db.Exec(CreateTable(obj, table, schema))
	if err != nil {
		return err
	}
	return nil
}

// FillTable fills schema.table with data from CSV using COPY FROM STDIN
func (db DB) FillTable(schema, table, dir string) error {
	r, err := os.Open("ivi/" + dir + "/" + table + ".csv")
	if err != nil {
		return err
	}
	defer r.Close()
	sql := fmt.Sprintf("COPY %s.%s FROM STDIN DELIMITER ';' CSV;", schema, table)
	_, err = db.CopyFrom(r, sql)
	if err != nil {
		return err
	}
	return nil
}

// AlterSchemasNames changes current schema name to a temporary one for its future removal
// and changes new schema name to previous schema name in one transaction
func (db DB) AlterSchemasNames(schema string) error {
	sql := `BEGIN;
ALTER SCHEMA %s RENAME TO old_%s;
ALTER SCHEMA new_%s RENAME TO %s;
COMMIT;`
	sql = fmt.Sprintf(sql, schema, schema, schema, schema)
	_, err := db.Exec(sql)
	return err
}

// Drop OldSchema drops the old schema which is no longer relevant
func (db DB) DropOldSchema(schema string) error {
	sql := fmt.Sprintf("DROP SCHEMA old_%s CASCADE;", schema)
	_, err := db.Exec(sql)
	return err
}
