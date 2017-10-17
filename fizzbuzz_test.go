package fizzbuzz

import (
	"testing"
)

type fizzBuzzTest struct {
	in int
	out string
}

var fizzBuzzTests = []fizzBuzzTest{
	fizzBuzzTest{2,"2"},
	fizzBuzzTest{3,"Fizz"},
	fizzBuzzTest{4,"4"},
	fizzBuzzTest{5,"Buzz"},
	fizzBuzzTest{6,"Fizz"},
	fizzBuzzTest{15,"FizzBuzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, fbt := range fizzBuzzTests {
		res := ToFizzBuzz(fbt.in)
		if res != fbt.out {
			t.Errorf("ToFizzBuzz(%d) = res must be %d",fbt.in, res, fbt.out)
		}
	}
}
