package xmlnode

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		xml  string
		want string
	}{
		{
			"<xml></xml>",
			"<xml></xml>",
		},
		{
			"<html><body></body></html>",
			"<html><body></body></html>",
		},
		{
			"<html><body><div></div></body></html>",
			"<html><body><div></div></body></html>",
		},
		{
			`<html><body id="body"></body></html>`,
			`<html><body id="body"></body></html>`,
		},
		{
			`<html><body id="body"><div id="i" class="c"></div></body></html>`,
			`<html><body id="body"><div id="i" class="c"></div></body></html>`,
		},
	}
	for _, test := range tests {
		dec := xml.NewDecoder(strings.NewReader(test.xml))
		got, err := Parse(dec)
		if err != nil {
			t.Errorf("Parse(%q): %v", test.xml, err)
		} else if got.String() != test.want {
			t.Errorf("Parse(%q) = %q, want %q", test.xml, got.String(), test.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		xml  Node
		want string
	}{
		{
			&Element{
				Type: xml.Name{
					Local: "html",
				},
				Children: []Node{
					&Element{
						Type: xml.Name{
							Local: "body",
						},
						Children: []Node{
							&Element{
								Type: xml.Name{
									Local: "div",
								},
								Attr: []xml.Attr{
									xml.Attr{
										Name: xml.Name{
											Local: "id",
										},
										Value: "foo",
									},
									xml.Attr{
										Name: xml.Name{
											Local: "class",
										},
										Value: "bar",
									},
								},
							},
						},
					},
				},
			},
			`<html><body><div id="foo" class="bar"></div></body></html>`,
		},
	}
	for _, test := range tests {
		if got := test.xml.String(); got != test.want {
			t.Errorf("(%v).String() = %q, want %q", test.xml, got, test.want)
		}
	}
}
