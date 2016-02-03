package render

import (
	"bytes"
	"reflect"
	"text/template"
)

type Template string

type Node struct {
	Parent   *Node
	Text     []byte
	Children []*Node
}

func RenderStruct(parent interface{}, renders MapRender, transforms MapTransform, rootTransform ParentTransform) []byte {
	structValue := reflect.TypeOf(parent)
	templates := buildTemplates(renders)
	if structValue.Kind() == reflect.Struct {
		parent := new(Node)
		for fieldIndex := 0; fieldIndex < structValue.NumField(); fieldIndex++ {
			fieldValue := structValue.Field(fieldIndex)
			kindField := fieldValue.Type.Kind()
			transform, template := transforms[kindField], templates[kindField]
			text, err := renderField(fieldValue, transform, template)
			if err != nil {
				panic(err)
			}
			childNode := &Node{Parent: parent, Text: text}
			parent.Children = append(parent.Children, childNode)
		}
		result, err := renderParent(parent, rootTransform, templates[reflect.Ptr])
		if err != nil {
			panic(err)
		}
		return result
	}
	return nil
}

func buildTemplates(renders MapRender) map[reflect.Kind]*template.Template {
	mapTemplates := make(map[reflect.Kind]*template.Template)
	for kind, render := range renders {
		tmplKind, err := template.New("").Parse(string(render))
		if err != nil {
			panic(err)
		}
		mapTemplates[kind] = tmplKind
	}
	return mapTemplates
}

func renderField(field reflect.StructField, trans Transform, tmpl *template.Template) ([]byte, error) {
	value := trans(field)
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, value)
	return buffer.Bytes(), err
}

func renderParent(node *Node, transform ParentTransform, tmpl *template.Template) ([]byte, error) {
	interfaceToRender := transform(node)
	buffer := new(bytes.Buffer)
	err := tmpl.Execute(buffer, interfaceToRender)
	return buffer.Bytes(), err
}
