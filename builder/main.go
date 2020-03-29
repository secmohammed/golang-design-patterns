package main

import (
    "fmt"
    "strings"
)

const (
    indentSize = 2
)

type HtmlElement struct {
    name, text string
    elements   []HtmlElement
}

func (e *HtmlElement) String() string {
    return e.string(0)
}
func (e *HtmlElement) string(indent int) string {
    sb := strings.Builder{}
    i := strings.Repeat(" ", indentSize*indent)
    sb.WriteString(
        fmt.Sprintf(
            "%s<%s>\n",
            i,
            e.name,
        ),
    )
    if len(e.text) > 0 {
        sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
        sb.WriteString(e.text)
        sb.WriteString("\n")
    }
    for _, el := range e.elements {
        sb.WriteString(el.string(indent + 1))
    }
    sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
    return sb.String()
}

type HtmlBuilder struct {
    rootName string
    root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
    return &HtmlBuilder{
        rootName,
        HtmlElement{rootName, "", []HtmlElement{}},
    }
}
func (b *HtmlBuilder) String() string {
    return b.root.String()
}
func (b *HtmlBuilder) AddChild(childName, childText string) {
    e := HtmlElement{childName, childText, []HtmlElement{}}
    b.root.elements = append(b.root.elements, e)

}
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
    e := HtmlElement{childName, childText, []HtmlElement{}}
    b.root.elements = append(b.root.elements, e)
    return b
}

func main() {
    text := "hello"
    sb := strings.Builder{}
    sb.WriteString("<p>")
    sb.WriteString(text)
    sb.WriteString("</p>")
    fmt.Println(sb.String())
    words := []string{"Hello", "World"}
    sb.Reset()
    // <ul><li>...</li><li>..</li></ul>
    sb.WriteString("<ul>")
    for _, v := range words {
        sb.WriteString("<li>")
        sb.WriteString(v)
        sb.WriteString("</li>")
    }
    sb.WriteString("</ul>")
    fmt.Println(sb.String())

    // new way using builder.
    b := NewHtmlBuilder("ul")
    for _, v := range words {
        b.AddChild("li", v)
    }
    fmt.Println(b.String())
    // Chaining by returning the reciever.
    b1 := NewHtmlBuilder("ul")
    b1.AddChildFluent("li", "hello").AddChildFluent("li", "world")
    fmt.Println(b1.String())

}
