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
	if amount%15 == 0 {
		return "FizzBuzz"
	} else if amount%3 == 0 {
		return "Fizz"
	} else if amount%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(amount)
	}
}
