package xmlbuilder

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func marshal(element Element) string {
	bytes, _ := xml.Marshal(element)
	return string(bytes)
}

func TestSingle(t *testing.T) {
	e := Tags(struct{ a Tag }{})
	assert.Equal(t, `<a>a</a>`, marshal(e.a("a")))
}

func TestNested(t *testing.T) {
	e := Tags(struct{ a, b Tag }{})
	assert.Equal(t, `<a><b>b</b></a>`, marshal(e.a(e.b("b"))))
}

func TestMultipleElements(t *testing.T) {
	e := Tags(struct{ a, b, c Tag }{})
	assert.Equal(t, `<a><b>b</b><c>c</c></a>`, marshal(e.a(e.b("b"), e.c("c"))))
}

func TestAppendElement(t *testing.T) {
	e := Tags(struct{ a, b, c Tag }{})
	doc := e.a(e.b("b"))
	doc.Append(e.c("c"))
	assert.Equal(t, `<a><b>b</b><c>c</c></a>`, marshal(doc))
}

func TestElementWithAttributes(t *testing.T) {
	e := Tags(struct{ a, b Tag }{})
	root := e.a(e.b("b", Attrs{"c": "d", "e:f": "g"}), Attrs{"xmlns:e": "x.y.z"})
	assert.Equal(t, `<a xmlns:e="x.y.z"><b c="d" e:f="g">b</b></a>`, marshal(root))
}

func TestInstantElement(t *testing.T) {
	root := E("a", "a")
	assert.Equal(t, `<a>a</a>`, marshal(root))
}
