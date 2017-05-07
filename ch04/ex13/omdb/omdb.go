// Package omdb は、Open Movie Database に対する Go の API を提供します。
package omdb

import "fmt"

import "net/url"
import "strings"

func searchURL(terms []string) string {
	return fmt.Sprintf("http://www.omdbapi.com/?t=%s", url.QueryEscape(strings.Join(terms, " ")))
}

// Movie は、Open Movie Database の映画情報を表します。
type Movie struct {
	Poster   string
	Response string
}
