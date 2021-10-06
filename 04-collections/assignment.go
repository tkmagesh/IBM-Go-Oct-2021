package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea temporConsequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"

	//find the size of the word that occurs the most by size and print the word size and the occurrence
	words := strings.Split(str, " ")
	wordsBySizeCount := getWordsBySizeCount(words)
	maxWordSize, maxWordSizeOccurrence := getMaxWordsBySize(wordsBySizeCount)
	fmt.Println("The word size that occurs the most is", maxWordSize, "characters long and occurs", maxWordSizeOccurrence, "times")
}

func getWordsBySizeCount(words []string) map[int]int {
	wordsBySizeCount := map[int]int{}
	for _, word := range words {
		wordsBySizeCount[len(word)]++
	}
	return wordsBySizeCount
}

func getMaxWordsBySize(wordsBySizeCount map[int]int) (int, int) {
	maxWordSize := 0
	maxWordSizeOccurrence := 0
	for wordSize, occurrence := range wordsBySizeCount {
		if occurrence > maxWordSizeOccurrence {
			maxWordSize = wordSize
			maxWordSizeOccurrence = occurrence
		}
	}
	return maxWordSize, maxWordSizeOccurrence
}
