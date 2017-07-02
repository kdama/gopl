package crawl

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// out は、ページの複製の保存先となるディレクトリです。
const out = "./out"

var tokens = make(chan struct{}, 20)

// Crawl は、対象のページをクロールして、ローカルディスクに保存します。
func Crawl(path string) (nextPaths []string) {
	fmt.Println(path)

	tokens <- struct{}{} // acquire a token
	list, w, err := Extract(path)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	err = save(path, w)

	if err != nil {
		log.Print(err)
	}

	for _, extracted := range list {
		same, err := sameDomain(path, extracted)
		if err != nil {
			log.Fatal(err)
		}
		if same {
			nextPaths = append(nextPaths, extracted)
		}
	}

	return
}

func sameDomain(x, y string) (bool, error) {
	px, err := url.Parse(x)
	if err != nil {
		return false, err
	}
	py, err := url.Parse(y)
	if err != nil {
		return false, err
	}
	return px.Host == py.Host, nil
}

func save(path string, data []byte) error {
	parsed, err := url.Parse(path)
	if err != nil {
		return err
	}

	localPath := out + "/" + parsed.Host + parsed.Path
	fmt.Println(parsed.Path)
	fmt.Println(filepath.Base(parsed.Path))
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
