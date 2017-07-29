// Package mapset は、組み込みのマップに基づく整数のセットを提供します。
package mapset

import (
	"bytes"
	"fmt"
	"sort"
)

// IntSet は、整数のセットです。そのゼロ値は空セットを表しています。
type IntSet struct {
	intmap map[int]struct{}
}

// NewIntSet は、新しい IntSet を返します。
func NewIntSet() *IntSet {
	return &IntSet{intmap: make(map[int]struct{})}
}

// Has は、負ではない値 x をセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	_, ok := s.intmap[x]
	return ok
}

// Add はセットに負ではない値 x を追加します。
func (s *IntSet) Add(x int) {
	s.intmap[x] = struct{}{}
}

// UnionWith は、s と t の和集合を s に設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for k := range t.intmap {
		s.intmap[k] = struct{}{}
	}
}

// String は、"{1 2 3}" の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, e := range s.Elems() {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", e)
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len は、要素数を返します。
func (s *IntSet) Len() int {
	return len(s.intmap)
}

// Remove は、セットから x を取り除きます。
func (s *IntSet) Remove(x int) {
	delete(s.intmap, x)
}

// Clear は、セットからすべての要素を取り除きます。
func (s *IntSet) Clear() {
	s.intmap = make(map[int]struct{})
}

// Copy は、セットのコピーを返します。
func (s *IntSet) Copy() *IntSet {
	out := NewIntSet()
	for k := range s.intmap {
		out.intmap[k] = struct{}{}
	}
	return out
}

// AddAll は、セットに負ではない複数の値を追加します。
func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		s.intmap[x] = struct{}{}
	}
}

// IntersectWith は、s と t の積集合を s に設定します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for k := range s.intmap {
		_, ok := t.intmap[k]
		if !ok {
			delete(s.intmap, k)
		}
	}
}

// DifferenceWith は、s と t の差集合を s に設定します。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for k := range s.intmap {
		_, ok := t.intmap[k]
		if ok {
			delete(s.intmap, k)
		}
	}
}

// SymmetricDifference は、s と t の対称差を s に設定します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for k := range t.intmap {
		_, ok := s.intmap[k]
		if ok {
			delete(s.intmap, k)
		} else {
			s.intmap[k] = struct{}{}
		}
	}
}

// Elems は、セットの要素を含むスライスを返します。
func (s *IntSet) Elems() []int {
	out := []int{}
	for k := range s.intmap {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}
