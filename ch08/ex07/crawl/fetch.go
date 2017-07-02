package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// fetch は、URL の内容と、HTML かどうか、エラーを返します。
func fetch(path string) (body []byte, isHTML bool, err error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("getting %s: %s", path, resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	isHTML = strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") ||
		strings.HasPrefix(resp.Header.Get("Content-Type"), "text/xhtml")
	return body, isHTML, err
}
