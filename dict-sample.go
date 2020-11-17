
package main

import (
	"fmt"
	"strings"
)

func AddWord(word string, wc map[string]int) map[string]int {
	if _, ok := wc[word]; ok {
		wc[word]++
	} else {
		wc[word] = 1
	}
	return wc;
}

func main() {

	const str string = "uhhhhhhhhhhhhhh life is better with clean hands hands"
	words := strings.Fields(str)

	wordcounts := make( map[string]int )

	for i := 0; i < len(words); i++ {

		wordcounts = AddWord(words[i],wordcounts)

	}

	for word, count := range wordcounts {

		fmt.Printf("word: %s\tcount: %d\n", word, count)

	}

}
