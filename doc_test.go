package xmlbuilder

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	e := Tags(struct{ person, gender, firstname, lastname Tag }{})
	doc := Doc(e.person(
		e.gender("female"),
		e.firstname("Anna"),
		e.lastname("Smith"),
	))
	bytes, _ := xml.MarshalIndent(doc, "", "    ")
	expected := `
	<person>
	    <gender>female</gender>
	    <firstname>Anna</firstname>
	    <lastname>Smith</lastname>
	</person>`
	assert.Equal(t, strings.ReplaceAll(strings.TrimSpace(expected), "\t", ""), string(bytes))
	bytes, _ = xml.MarshalIndent(&doc, "", "    ")
	assert.Equal(t, strings.ReplaceAll(strings.TrimSpace(expected), "\t", ""), string(bytes))
}
