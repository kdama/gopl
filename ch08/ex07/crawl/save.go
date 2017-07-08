package crawl

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Save は、与えられたバイト配列のデータをローカルディスクに保存します。
func Save(path, out string, data []byte) error {
	parsed, err := url.Parse(path)
	if err != nil {
		return err
	}

	localPath := out + "/" + parsed.Host + parsed.Path
	if parsed.Path == "" || !strings.Contains(filepath.Base(parsed.Path), ".") {
		localPath += "/index.html"
	}
	err = os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(localPath)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return err
}
