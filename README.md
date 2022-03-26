XML builder for Golang
======================

xml-builder is a metaprogramming approach to build XML document like lxml in Python.

**Creating a XML document**

The following example creates an XML document from scratch using the xml-builder package 
and outputs its indented contents to stdout.

```go
e := Tags(struct{ person, gender, firstname, lastname Tag }{})
doc := Doc(e.person(
	e.gender("female"),
	e.firstname("Anna"),
	e.lastname("Smith"),
))
bytes, _ := xml.MarshalIndent(doc, "", "    ")
fmt.Println(string(bytes))
```

Will output:

```
<person>
  <gender>female</gender>
  <firstname>Anna</firstname>
  <lastname>Smith</lastname>
</person>
```

**Attributes**

This example illustrates how to define attributes in element

```go
e := Tags(struct{ person, gender Tag }{})
doc := Doc(e.person(
	e.gender("female", Attrs{"a": "b", "xmlns:c": "xml.builder"}),
))
bytes, _ := xml.MarshalIndent(doc, "", "    ")
fmt.Println(string(bytes))
```

Will output:

```
<person>
  <gender a="b" xmlns:c="xml.builder">female</gender>
</person>
```

**InstantElement**

When unable to define tags, you can create as instant element as following:

```go
person := E("person")
gender := E("gender", "female")
person.Append(gender)
```


**Hook**

When element name is snakecase or kebabcase, assuming you already had a function named snakecase, 
you can add a hook when creating tags as following:

```go
e := Tags(struct{ firstName, lastName Tag }{}, func(name *xml.Name) {
	name.Local = snakecase(name.Local)
})
```
