package main

import (
	"XMLparser/structureXML"
	"testing"
)

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

// text/template
// https://golang.org/pkg/text/template/#example_Template

// rewrite sql => schema ivi_new
// create schema ivi;
// create table ivi.countries...

// new file with copy from

// begin; rename ivi to ivi_old; rename ivi_new to ivi; drop ivi_old; commit;
