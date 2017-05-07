// ch04/ex12 は、xkcd から検索語と一致するコミックの URL と内容を表示します。
package main

import (
	"flag"
	"fmt"

	"encoding/json"
	"os"

	"io/ioutil"

	"github.com/kdama/gopl/ch04/ex12/xkcd"
)

const (
	maxComicID = 1833
)

var fetchFlag = flag.Bool("fetch", false, "fetch all comics")

func main() {
	flag.Parse()

	if *fetchFlag {
		fetch()
	} else {
		if len(flag.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "ch04/ex12: must have at least 1 query")
			os.Exit(1)
		}
		search(flag.Args())
	}
}

func fetch() {
	comicIndex := xkcd.NewComicIndex()
	for comicID := 1; comicID <= maxComicID; comicID++ {
		// for comicID := 1; comicID <= 1833; comicID++ {
		comic, err := xkcd.GetComic(comicID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ch04/ex12: get %d: %v\n", comicID, err)
		} else {
			comicIndex.Comics = append(comicIndex.Comics, comic)
		}
	}
	result, err := json.MarshalIndent(comicIndex, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", result)
	}
}

func search(terms []string) {
	index, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	}

	comics := xkcd.NewComicIndex()
	err = json.Unmarshal(index, &comics)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch04/ex12: %v\n", err)
		os.Exit(1)
	}

	searchResult := xkcd.SearchComic(comics, terms)

	fmt.Printf("%d comics:\n", len(searchResult))
	for _, comic := range searchResult {
		printComic(comic)
	}
}

func printComic(comic *xkcd.Comic) {
	fmt.Printf("\n-- Comic %d --\n", comic.Num)
	fmt.Printf("\nImage URL:\n%s\n", comic.Img)
	fmt.Printf("\nTranscript:\n%s\n", comic.Transcript)
}
