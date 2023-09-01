package main

import (
	"fmt"
	"strconv"
)

var majorArcana = map[int]string{
	0: "The Fool",
	1: "The Magician",
	2: "The High Priestess",
	3: "The Empress",
	4: "The Emperor",
	5: "The Hierophant",
	6: "The Lovers",
	7: "The Chariot",
	8: "Strength",
	9: "The Hermit",
}

func main() {
	fmt.Print("Please provide a date in the format DD/MM/YYYY: ")
	var date string
	fmt.Scanln(&date)

	sum := calculateDigitSum(date)

	cardNumber := sum % 9
	card := majorArcana[cardNumber]

	fmt.Printf("Your Tarot Major Arcana card for the date %s is: %s\n", date, card)
}

func calculateDigitSum(date string) int {
	digits := make([]int, 0)

	for _, char := range date {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			digits = append(digits, digit)
		}
	}

	for len(digits) > 1 {
		sum := 0
		for _, digit := range digits {
			sum += digit
		}
		digits = digitsFromNumber(sum)
	}

	return digits[0]
}

func digitsFromNumber(num int) []int {
	digits := make([]int, 0)

	for num > 0 {
		digit := num % 10
		digits = append([]int{digit}, digits...)
		num /= 10
	}

	return digits
}
