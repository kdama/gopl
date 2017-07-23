package archive

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

// ErrFormat は、未知のフォーマットを報告します。
var ErrFormat = errors.New("archive: unknown format")

type magic struct {
	str    string
	offset int
}

type format struct {
	name  string
	magic magic
	list  func(*os.File) ([]FileHeader, error)
}

var formats []format

// RegisterFormat は、フォーマットを登録します。
func RegisterFormat(name string, magicStr string, magicOffset int, list func(*os.File) ([]FileHeader, error)) {
	formats = append(formats, format{name, magic{magicStr, magicOffset}, list})
}

type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

func sniff(file *os.File) (format, error) {
	for _, f := range formats {
		file.Seek(0, io.SeekStart)
		r := asReader(file)
		b, err := r.Peek(f.magic.offset + len(f.magic.str))
		if err == nil && bytes.Equal([]byte(f.magic.str), b[f.magic.offset:]) {
			file.Seek(0, io.SeekStart)
			return f, nil
		}
	}
	return format{}, ErrFormat
}
