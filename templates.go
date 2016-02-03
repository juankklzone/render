package render

import (
	"reflect"
	"text/template"
)

type MapRender map[reflect.Kind]Template

type Transform func(reflect.StructField) interface{}

type ParentTransform func(*Node) interface{}

type MapTransform map[reflect.Kind]Transform

type mapTemplate map[reflect.Kind]*template.Template

var (
	templateInts   Template = `<label for="{{.Name}}">{{.Name}}</label><br><input type="number" name="{{.Name}}"></input><br>`
	templateString Template = `<label for="{{.Name}}">{{.Name}}</label><br><input type="text" name="{{.Name}}"></input><br>`
	templateTime   Template = `<label for="{{.Name}}">{{.Name}}</label><br><input type="date" name="{{.Name}}"></input><br>`
	templateParent Template = `<form>{{.}}</form>`
)
