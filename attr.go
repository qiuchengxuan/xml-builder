package xmlbuilder

import (
	"encoding/xml"
)

type Attrs map[string]string

func xmlAttrs(attrs Attrs) []xml.Attr {
	xmlAttrs := make([]xml.Attr, 0, len(attrs))
	for key, value := range attrs {
		attr := xml.Attr{Name: xml.Name{Space: "", Local: key}, Value: value}
		xmlAttrs = append(xmlAttrs, attr)
	}
	return xmlAttrs
}
