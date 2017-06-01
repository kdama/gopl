// Package sorting は、曲に関する多段ソートを提供します。
package sorting

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"
	"time"
)

// MultiSort は、曲に関する多段ソートを提供します。
func MultiSort(tracks []*Track, columns []string) sort.Interface {
	return customSort{
		tracks,
		func(x, y *Track) bool {
			for i := len(columns) - 1; i >= 0; i-- {
				if columns[i] == "Title" && x.Title != y.Title {
					return x.Title < y.Title
				} else if columns[i] == "Artist" && x.Artist != y.Artist {
					return x.Artist < y.Artist
				} else if columns[i] == "Album" && x.Album != y.Album {
					return x.Album < y.Album
				} else if columns[i] == "Year" && x.Year != y.Year {
					return x.Year < y.Year
				} else if columns[i] == "Length" && x.Length != y.Length {
					return x.Length < y.Length
				}
			}
			return true
		},
	}
}

// Track は、曲を表します。
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// Length は、文字列をパースして、曲の長さとして返します。
func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// FprintTracks は、曲のスライスを文字列による表現を io.Writer に書き込みます。
func FprintTracks(w io.Writer, tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(w, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
