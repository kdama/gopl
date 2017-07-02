package crawl

import (
	"fmt"
	"log"
	"net/url"
)

var tokens = make(chan struct{}, 20)

// Crawl は、対象のページをクロールして、ローカルディスクに保存します。
func Crawl(path, out string) (nextPaths []string) {
	fmt.Println(path)

	tokens <- struct{}{} // acquire a token
	nextPaths, err := crawl(path, out)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	return nextPaths
}

func crawl(path, out string) (nextPaths []string, err error) {
	body, isHTML, err := fetch(path)
	if err != nil {
		return
	}

	// レスポンスが HTML ならば、リンクを抽出し、相対パスへの変換を行います。
	var links []string
	if isHTML {
		links, body, err = extract(path, body)
		if err != nil {
			return
		}
	}

	err = save(path, out, body)
	if err != nil {
		return
	}

	for _, extracted := range links {
		same, err := sameDomain(path, extracted)
		if err != nil {
			log.Print(err)
			continue
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
