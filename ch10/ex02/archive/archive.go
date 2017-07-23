// Package archive は、汎用のアーカイブ読み込み機能を提供します。ZIP と POSIX tar をサポートします。
package archive

import (
	"os"
)

// FileHeader は、アーカイブに含まれるファイルの名前と圧縮前のサイズを表現します。
type FileHeader struct {
	Name string
	Size uint64
}

// List は、アーカイブに含まれるファイルの名前と圧縮前のサイズを報告します。
func List(f *os.File) ([]FileHeader, error) {
	format, err := sniff(f)
	if err != nil {
		return nil, err
	}
	return format.list(f)
}
