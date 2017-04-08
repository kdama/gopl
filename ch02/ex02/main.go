package main

import (
	"bufio"
	"fmt"
	"os"

	"./conv"
)

func main() {
	if len(os.Args) > 1 {
		fromArgs()
	} else {
		fromStdin()
	}
}

func fromArgs() {
	for _, arg := range os.Args[1:] {
		run(arg)
	}
}

func fromStdin() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		run(s.Text())
	}
	if s.Err() != nil {
		fmt.Fprintf(os.Stderr, "conv: %v\n", s.Err())
		os.Exit(1)
	}
}

func run(str string) {
	value, unit, err := conv.Parse(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conv: %s\n", err)
		os.Exit(1)
	}
	from, to, err := conv.Convert(value, unit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conv: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s = %s\n", from, to)
}
