package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func help() {
	fmt.Println()
	fmt.Println("Instruction:\n 1. input two integers seperated by a single space\n 2. Then chose an operation to use")
	fmt.Println()

}

func main() {

	var userinput, userinput1 string
	for {
		fmt.Println("- Enter two input seperated by newlines: ")

		fmt.Scan(&userinput)

		fmt.Scan(&userinput1)

		num, err := strconv.Atoi(userinput)
		if err != nil {
			fmt.Println("invalid input\n", err)
			continue
		}
		num1, err := strconv.Atoi(userinput1)
		if err != nil {
			fmt.Println("invalid input\n", err)
			continue
		}

		var operation int
		fmt.Println("OPERATIONS:\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Help \n 6. Exit")
		fmt.Scanln(&operation)

		if operation == 6 {
			fmt.Println("GoodBye!!")
			return
		}

		switch operation {
		case 1:
			fmt.Println(add(num, num1))
		case 2:
			fmt.Println(sub(num, num1))
		case 3:
			fmt.Println(mul(num, num1))
		case 4:
			if num1 == 0 || num == 0 {
				fmt.Println("invalid input: number cannot be divided by zero")
				continue
			} else {
				fmt.Println(div(num, num1))
			}
		case 5:
			help()
		default:
			fmt.Println("Invalid Input")
			continue
		}
	}
}
