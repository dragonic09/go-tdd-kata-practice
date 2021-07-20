package comparator

import "testing"

func TestFizzBuzzShouldReturnFizzWhenOnlyDivisibleByThree(t *testing.T) {
	result := FizzBuzz(9)
	if result != "Fizz" {
		t.Errorf("FizzBuzz failed, expect the function to return Fizz when input is divisible only by 3.")
	}
}

func TestFizzBuzzShouldReturnBuzzWhenOnlyDivisibleByFive(t *testing.T) {
	result := FizzBuzz(10)
	if result != "Buzz" {
		t.Errorf("FizzBuzz failed, expect the function to return Fizz when input is divisible only by 5.")
	}
}

func TestFizzBuzzShouldReturnFizzBuzzWhenDivisibleByBothThreeAndFive(t *testing.T) {
	result := FizzBuzz(15)
	if result != "FizzBuzz" {
		t.Errorf("FizzBuzz failed, expect the function to return FizzBuzz when input is divisible by 3 and 5.")
	}
}

func TestFizzBuzzShouldReturnItsInputWhenItsNotDivisibleByThreeAndFive(t *testing.T) {
	result := FizzBuzz(1)
	if result != "1" {
		t.Errorf("FizzBuzz failed, expect the function to return its input when input is not divisible by 3 and 5.")
	}
}

