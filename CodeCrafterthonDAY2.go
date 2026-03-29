package main

import (
	"fmt"
	"strconv"
	"strings"
)

func converthex(a, b string) string {
	var dec string
	if b == "hex" {

		for i := 0; i < len(a); i++ {
			num, err := strconv.ParseInt(a, 16, 64)
			if err != nil {
				return "Invalid hexadecimal"

			} else {

				dec = strconv.FormatInt(num, 10)

			}
		}
	}
	return strings.ToUpper(dec)

}
func convertbin(a, b string) string {
	var dec string
	if b == "bin" {

		for i := 0; i < len(a); i++ {
			num, err := strconv.ParseInt(a, 2, 64)
			if err != nil {
				return "Invalid binary"

			} else {

				dec = strconv.FormatInt(num, 10)

			}
		}
	}
	return dec

}

func convertdec(a, b string) string {
	var dec string
	if b == "dec" {

		for i := 0; i < len(a); i++ {
			num, err := strconv.ParseInt(a, 10, 64)
			if err != nil {
				return "invalid decimal"

			} else {

				dec = strconv.FormatInt(num, 2)

			}
		}
	}
	return dec

}
func convertdec1(a, b string) string {
	var dec string
	if b == "dec" {

		for i := 0; i < len(a); i++ {
			num, err := strconv.ParseInt(a, 10, 64)
			if err != nil {
				return "invalid decimal"

			} else {

				dec = strconv.FormatInt(num, 16)

			}
		}
	}
	return strings.ToUpper(dec)

}

func main() {
	for {
		var userinput string
		fmt.Print("Enter input string: ")
		fmt.Scan(&userinput)

		var userinput1 string
		fmt.Scan(&userinput1)

		var userinput2 string
		fmt.Scan(&userinput2)
		if userinput2 == " " {
			fmt.Println("provide a valid input for conversion")
			return
		}

		var option int
		fmt.Println("1. Hexadecimal to decimal\n 2. Binary to decimal\n 3. Decimal to binary\n 4. Decimal to hexadecimal\n 5. Quit")
		fmt.Scan(&option)
		if option == 5 {
			fmt.Println("Goodbye")
			return
		}

		switch option {
		case 1:
			fmt.Println(converthex(userinput1, userinput2))
		case 2:
			fmt.Println(convertbin(userinput1, userinput2))
		case 3:
			fmt.Println(convertdec(userinput1, userinput2))
		case 4:
			fmt.Println(convertdec1(userinput1, userinput2))
			continue
		default:
			fmt.Println("invalid input")
			break
		}

	}
}
