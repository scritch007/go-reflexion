package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

type Bar struct {
	NomFamille string
	Prenom     string
	Age        int
}

type SubType struct {
	FirstName string
	Age       int
}

func update(f interface{}, b interface{}, mapping map[string]string) {
	val := reflect.ValueOf(f).Elem()
	val2 := reflect.ValueOf(b).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
		m, found := mapping[typeField.Name]
		if found {
			r := val2.FieldByName(m)
			t, found2 := val2.Type().FieldByName(m)
			if found2 {
				fmt.Printf("%s %v\n", t.Name, r.Interface())
				r.Set(valueField)
			}
		}

	}
}

func subType(a interface{}, b interface{}) {
	toUpdate := reflect.ValueOf(a).Elem()
	from := reflect.ValueOf(b).Elem()
	for i := 0; i < toUpdate.NumField(); i++ {
		valueField := toUpdate.Field(i)
		typeField := toUpdate.Type().Field(i)
		v := from.FieldByName(typeField.Name)
		valueField.Set(v)
	}
}

func main() {
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}
	b := &Bar{}

	mapping := make(map[string]string)
	mapping["FirstName"] = "NomFamille"
	update(f, b, mapping)
	fmt.Printf("Nouveau Nom de Famille %s\n", b.NomFamille)
	s := &SubType{}
	subType(s, f)
	fmt.Printf("%s %d\n", s.FirstName, s.Age)
}
