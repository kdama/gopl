// ch12/ex05 は、Go の値を JSON 形式でエンコードして、プリティプリントします。
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/kdama/gopl/ch12/ex05/json"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		True            bool
		False           bool
		Float           float64
		Interface       interface{}
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		True:      true,
		False:     false,
		Float:     math.Pi,
		Interface: []int{1, 2, 3},
	}

	data, err := json.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
