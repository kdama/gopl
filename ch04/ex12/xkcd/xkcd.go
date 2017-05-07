// Package xkcd は、xkcd に対する Go の API を提供します。
package xkcd

import "fmt"
import "strconv"

func getComicURL(comicID int) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", strconv.Itoa(comicID))
}

// ComicIndex は、xkcd のコミックのインデックスを表します。
type ComicIndex struct {
	Comics []*Comic
}

// NewComicIndex は、新しい xkcd コミックインデックスを返します。
func NewComicIndex() ComicIndex {
	return ComicIndex{[]*Comic{}}
}

// Comic は、xkcd のコミックを表します。
type Comic struct {
	Alt        string
	Day        string
	Img        string
	Link       string
	Month      string
	News       string
	Num        int
	SafeTitle  string
	Title      string
	Transcript string
	Year       string
}
