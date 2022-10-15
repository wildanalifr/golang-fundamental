package quiz

import "fmt"

func Average(numbers []int) float32 {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return float32(total) / float32(len(numbers))
}

func GetBiggerNumber(numbers []int) string {
	for _, number := range numbers {
		if number >= 90 {
			fmt.Println(number)
		}
	}
	return ""
}
