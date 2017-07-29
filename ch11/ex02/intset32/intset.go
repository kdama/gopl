// Package intset32 は、ビットベクタベースの整数のセットを提供します。
package intset32

import (
	"bytes"
	"fmt"
)

// IntSet は、負ではない小さな整数のセットです。そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint32
}

// NewIntSet は、新しい IntSet を返します。
func NewIntSet() *IntSet {
	return &IntSet{words: []uint32{}}
}

// Has は、負ではない値 x をセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add はセットに負ではない値 x を追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith は、s と t の和集合を s に設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String は、"{1 2 3}" の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len は、要素数を返します。
func (s *IntSet) Len() int {
	count := 0
	for _, w := range s.words {
		for w != 0 {
			w &= w - 1
			count++
		}
	}
	return count
}

// Remove は、セットから x を取り除きます。
func (s *IntSet) Remove(x int) {
	word, bit := x/32, uint(x%32)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// Clear は、セットからすべての要素を取り除きます。
func (s *IntSet) Clear() {
	s.words = []uint32{}
}

// Copy は、セットのコピーを返します。
func (s *IntSet) Copy() *IntSet {
	out := &IntSet{}
	for _, w := range s.words {
		out.words = append(out.words, w)
	}
	return out
}

// AddAll は、セットに負ではない複数の値を追加します。
func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		word, bit := x/32, uint(x%32)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

// IntersectWith は、s と t の積集合を s に設定します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith は、s と t の差集合を s に設定します。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymmetricDifference は、s と t の対称差を s に設定します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

// Elems は、セットの要素を含むスライスを返します。
func (s *IntSet) Elems() []int {
	out := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				out = append(out, 32*i+j)
			}
		}
	}
	return out
}
