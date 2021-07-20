package comparator

import "fmt"

func FizzBuzz(number int) string {
	isFizz := isFizz(number)
	isBuzz := isBuzz(number)

	if isFizz && isBuzz {
		return "FizzBuzz"
	}

	if isFizz {
		return "Fizz"
	}

	if isBuzz {
		return "Buzz"
	}

	return fmt.Sprint(number)
}

func isFizz(number int) bool {
	return number % 3 == 0
}

func isBuzz(number int) bool {
	return number % 5 == 0
}