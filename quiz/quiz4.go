package quiz

import "errors"

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func Calculate(number1 int, number2 int, operation string) (int, error) {
	var total int
	var errorResult error

	switch operation {
	case "+":
		total = number1 + number2
	case "-":
		total = number1 - number2
	case "*":
		total = number1 * number2
	case "/":
		total = number1 / number2
	default:
		errorResult = errors.New("Unknown operation")
	}

	return total, errorResult
}
