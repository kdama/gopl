package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	count := 100000000
	{
		start := time.Now()
		for i := 0; i < count; i++ {
			var s, sep string
			for i := 1; i < len(os.Args); i++ {
				s += sep + os.Args[i]
				sep = " "
			}
			// fmt.Println(s)
		}
		secs := time.Since(start).Seconds()
		fmt.Printf("echo1: %.2f\n", secs)
	}
	{
		start := time.Now()
		for i := 0; i < count; i++ {
			s, sep := "", ""
			for _, arg := range os.Args[1:] {
				s += sep + arg
				sep = " "
			}
			// fmt.Println(s)
		}
		secs := time.Since(start).Seconds()
		fmt.Printf("echo2: %.2f\n", secs)
	}
	{
		start := time.Now()
		for i := 0; i < count; i++ {
			strings.Join(os.Args[1:], " ")
			// fmt.Println(strings.Join(os.Args[1:], " "))
		}
		secs := time.Since(start).Seconds()
		fmt.Printf("echo3: %.2f\n", secs)
	}
}
