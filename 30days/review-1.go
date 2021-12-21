package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Today we will expand our knowledge of strings, combining it with what we have
already learned about loops.

Task: Given a string, S, of length N that is indexed from 0 to N-1, print its
even-indexed and odd-indexed characters as 2 space-separated strings on a single
line.

Example:s = adbecf, prints abc def
*/


func main() {
	var N int
	fmt.Scan(&N)

	words := make([]string, N)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < N; i++ {
		words[i], _ = reader.ReadString('\n')
		words[i] = strings.TrimSuffix(words[i], "\n")
	}

	for j := range words {
		word := words[j]
		for i, char := range word {
			if i % 2 == 0 {
				fmt.Printf("%c", char)
			}
		}
		fmt.Print(" ")
		for i, char := range word {
			if i % 2 != 0 {
				fmt.Printf("%c", char)
			}
		}
		fmt.Printf("\n")
	}
}