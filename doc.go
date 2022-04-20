package xmlbuilder

type element = Element

// Simply a wrapper since Element has no pointer receiver to MarshalXML method
type Document struct {
	element
}

func Doc(element Element) Document {
	return Document{element}
}
