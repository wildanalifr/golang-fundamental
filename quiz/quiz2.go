package quiz

import "fmt"

func PrintForLoop(sentences string) {
	for i, sentence := range sentences {
		if i%2 == 0 {
			fmt.Print(string(sentence))
		}
	}
}
