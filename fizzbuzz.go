package fizzbuzz

import (
	"fmt"
	"strconv"
)

func SayFizzBuzz() {
	for i:= 1; i <= 100; i++ {
		fmt.Println(ToFizzBuzz(i))
	}
}

func ToFizzBuzz(amount int) string {
	switch {
	case amount%15 == 0:
		return "FizzBuzz"
	case amount%3 == 0:
		return "Fizz"
	case amount%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(amount)
	}
}
