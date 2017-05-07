// Package xkcd は、xkcd に対する Go の API を提供します。
package xkcd

import (
	"strings"
)

// SearchComic は 検索語に関連する xkcd のコミックを返します。
func SearchComic(from ComicIndex, terms []string) []*Comic {
	result := []*Comic{}

	for _, comic := range from.Comics {
		if hit(comic, terms) {
			result = append(result, comic)
		}
	}

	return result
}

func hit(comic *Comic, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(comic.Alt, term) {
			continue
		} else if strings.Contains(comic.Day, term) {
			continue
		} else if strings.Contains(comic.Img, term) {
			continue
		} else if strings.Contains(comic.Link, term) {
			continue
		} else if strings.Contains(comic.Month, term) {
			continue
		} else if strings.Contains(comic.News, term) {
			continue
		} else if strings.Contains(comic.SafeTitle, term) {
			continue
		} else if strings.Contains(comic.Title, term) {
			continue
		} else if strings.Contains(comic.Transcript, term) {
			continue
		} else if strings.Contains(comic.Year, term) {
			continue
		}
		return false
	}
	return true
}
