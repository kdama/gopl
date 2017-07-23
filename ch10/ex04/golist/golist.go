// Package golist は、go list ツールに対する API を提供します。
package golist

// Package は、go list ツールにより出力されるパッケージのメタデータを表します。
type Package struct {
	ImportPath string
	Deps       []string
}
