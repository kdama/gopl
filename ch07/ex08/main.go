// ch07/ex08 は、曲に対する、状態を持つ多段ソートの実装です。
package main

import (
	"os"
	"sort"

	"github.com/kdama/gopl/ch07/ex08/sorting"
)

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
}

func main() {
	// Year でソートした後に、Title でソートした場合の結果を表示します。
	sort.Sort(sorting.MultiSort(tracks, []string{
		"Year",
		"Title",
	}))
	sorting.FprintTracks(os.Stdout, tracks)
}
