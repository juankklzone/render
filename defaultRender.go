package render

import (
	"bytes"
	"reflect"
)

var defaultMapRender map[reflect.Kind]Template = map[reflect.Kind]Template{reflect.Int: templateInts, reflect.String: templateString, reflect.Struct: templateTime, reflect.Ptr: templateParent}

type value struct {
	Name string
}

func defaultTransform(field reflect.StructField) interface{} {
	return value{Name: field.Name}
}

func defaultParentTransform(n *Node) interface{} {
	buffer := new(bytes.Buffer)
	for _, child := range n.Children {
		buffer.Write(child.Text)
	}
	return string(buffer.Bytes())
}

func DefaultRender(parent interface{}) []byte {
	defaultMapTransf := make(map[reflect.Kind]Transform)
	for kind := range defaultMapRender {
		defaultMapTransf[kind] = defaultTransform
	}
	return RenderStruct(parent, defaultMapRender, defaultMapTransf, defaultParentTransform)
}
