// ch12/ex06 は、Go の値を S 式形式でエンコードします。ただし、フィールドの値がその型のゼロ値である場合には、フィールドをエンコードしません。
package main

import (
	"fmt"
	"log"

	"github.com/kdama/gopl/ch12/ex06/sexpr"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
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
	}

	data, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
