package main

import "fmt"

//function fizzBuzz generates the iteration sequence for numbers and return string when conditions is met
func fizzBuzz() string {

	//defining an counter string to store the iteration sequence
	var count string

	//for loop to iterate between 1 and 15
	for i := 1; i <= 15; i++ {

		// first condition to check is when the number divided by 3 and 5, calculating the remainder
		if i%3 == 0 && i%5 == 0 {
			count += "FizzBuzz\n" // if the condition is met, append "FizzBuzz" to the count
		} else if i%3 == 0 { // second condition to check is when the number divided by 3, calculating the remainder
			count += "Fizz\n" // if the condition is met, append "Fizz" to the count
		} else if i%5 == 0 { // third condition to check is when the number divided by 5, calculating the remainder
			count += "Buzz\n" // if the condition is met, append "Buzz" to the count
		} else {
			// if the number did not meet any conditions
			count += fmt.Sprintf("%d\n", i) // append the number to the count
		}
	}
	return count //return the generated sequence
}

func main() {
	//call the fizzBuzz function and print the result
	fmt.Println(fizzBuzz())
}
