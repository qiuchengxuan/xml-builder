package xmlbuilder

// Simply a wrapper since Element has no pointer receiver to MarshalXML method
type Document struct {
	Element
}

func Doc(element Element) Document {
	return Document{element}
}
