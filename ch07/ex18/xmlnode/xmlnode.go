// Package xmlnode は、XML のノードのツリーを構築します。
package xmlnode

import (
	"encoding/xml"
	"fmt"
)

// Node は、XML のノードを表します。
type Node interface {
	String() string
}

// CharData は、テキスト文字列を表すノードです。
type CharData string

// Element は、名前付き要素とその属性を表すノードです。
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// Parse は、XML デコーダを用いて、XML のノードツリーを返します。
func Parse(dec *xml.Decoder) (Node, error) {
	var stack []*Element // 走査の途中で、まだ閉じられていないタグ要素です。

	for {
		tok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &element)
			}
			stack = append(stack, &element)
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stack) == 1 {
				// 根の要素の終了タグを発見したので、根の要素を返して終了します。
				return stack[0], nil
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			// 親要素の無いテキスト文字列は無視します。
			if len(stack) > 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, CharData(tok))
			}
		}
	}
}

func (c CharData) String() string {
	return string(c)
}

func (e *Element) String() string {
	var attrs, children string
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>", e.Type.Local, attrs, children, e.Type.Local)
}
