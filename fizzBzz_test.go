package main

import (
	"fmt"
	"testing"
)

// function to generate the expected output
func generateExpectedFizzBuzz() string {

	var result string

	for j := 1; j <= 15; j++ {

		// first condition to check is when the number divided by 3 and 5, calculating the remainder
		if j%3 == 0 && j%5 == 0 {
			result += "FizzBuzz\n" // if the condition is met, append "FizzBuzz" to the count
		} else if j%3 == 0 { // second condition to check is when the number divided by 3, calculating the remainder
			result += "Fizz\n" // if the condition is met, append "Fizz" to the count
		} else if j%5 == 0 { // third condition to check is when the number divided by 5, calculating the remainder
			result += "Buzz\n" // if the condition is met, append "Buzz" to the count
		} else {
			// if the number did not meet any conditions
			result += fmt.Sprintf("%d\n", j) // append the number to the count
		}
	}
	return result //return the generated sequence
}

// function to test fizzBuzz
func TestFizzBuzz(t *testing.T) {
	// t.Run allows you to run subtests within a test function
	t.Run("testing FizzBuzz for 15 iterations", func(t *testing.T) {

		expected := generateExpectedFizzBuzz() // generate the expected fizzBuzz output sequence

		got := fizzBuzz() // call the fizzBuzz function

		if got != expected { // compare the expected and the actual output
			t.Errorf(" expected:\n%s\nbut got:\n%s", expected, got) // if no match report the error
		}
	})
}
