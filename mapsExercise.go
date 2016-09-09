// My answer for the maps exercise at https://tour.golang.org/moretypes/23
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	var m map[string]int
	m = make(map[string]int)
	for i:=0; i < len(words); i++ {
		m[words[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
