package xmlbuilder

import (
	"encoding/xml"
	"reflect"
	"unsafe"
)

type Element struct {
	start  xml.StartElement
	tokens []xml.Token
}

func (e Element) MarshalXML(encoder *xml.Encoder, _start xml.StartElement) error {
	return encoder.EncodeElement(e.tokens, e.start)
}

func (e *Element) Append(value Element) *Element {
	if wrapper1, ok := e.tokens[len(e.tokens)-1].(wrapper); ok {
		elements := append(wrapper1.Elements, value)
		e.tokens[len(e.tokens)-1] = wrapper{Elements: elements}
	} else {
		e.tokens = append(e.tokens, wrapper{Elements: []Element{value}})
	}
	return e
}

type wrapper struct {
	Elements []Element `xml:"elements"`
}

type Tag func(tokens ...xml.Token) Element

func makeTag(name xml.Name) Tag {
	return func(args ...xml.Token) Element {
		var elements []Element
		var tokens []xml.Token
		var attrs []xml.Attr
		for _, arg := range args {
			switch v := arg.(type) {
			case Element:
				elements = append(elements, v)
			case Attrs:
				attrs = append(attrs, xmlAttrs(v)...)
			default:
				tokens = append(tokens, arg)
			}
		}
		if len(elements) > 0 {
			tokens = append(tokens, wrapper{Elements: elements})
		}
		return Element{
			start:  xml.StartElement{Name: name, Attr: attrs},
			tokens: tokens,
		}
	}
}

func E(name string, tokens ...xml.Token) Element {
	return makeTag(xml.Name{Local: name})(tokens...)
}

func Tags[T any](tags T, hooks ...func(*xml.Name)) T {
	elem := reflect.ValueOf(&tags).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		if _, ok := field.Interface().(Tag); !ok {
			continue
		}
		fieldType := elem.Type().Field(i)
		space := fieldType.Tag.Get("xmlns")
		local := elem.Type().Field(i).Name
		name := xml.Name{Space: space, Local: local}
		for _, hook := range hooks {
			hook(&name)
		}
		field.Set(reflect.ValueOf(makeTag(name)))
	}
	return tags
}
