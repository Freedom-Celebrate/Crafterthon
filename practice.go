package main

import (
	"fmt"
	"strconv"
	"strings"
)

func uppercaselastword(s []string) []string {

	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}

	}
	return s
}

func ConvertNumbers(s []string) []string {
	var k int
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {

			num, _ := strconv.ParseInt(s[i-1], 16, 64)
			s[i] = fmt.Sprint(num)
			k = i - 1
		}
	}

	return append(s[:k], s[k+1:]...)
}

func ConvertNumbers1(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			num, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				// fmt.Println("unsuccessful conversion", err)
				continue
			}
			s[i-1] = strconv.FormatInt(num, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		} else if s[i] == "(bin)" {
			num, _ := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = strconv.FormatInt(num, 10)
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}
func decimaltoHex(s []string) []string {

	var k int

	for i := 0; i < len(s); i++ {
		if s[i] == "(dec)" && i > 0 {
			pword := s[i-1]
			v, _ := strconv.Atoi(pword)
			t, _ := fmt.Printf("%X", v)
			fmt.Println(t)
			u := fmt.Sprint(t)
			s[i] = u
			k = i - 1

		}
	}

	return append(s[:k], s[k+1:]...)
}

func main() {
	// fmt.Println(ConvertNumbers([]string{"first", "1E", "(hex)", "First", "fire"}))
	// //	x := map[string]map[string]string{"user1":{"username":"victor",
	// //
	// // "age":"20",
	// // "gender":"male"}}
	// // y := fmt.Sprintln(x)
	// name := "favour"
	// fmt.Printf("%T", name[1])
	// if name[1] == 'a' {
	// 	fmt.Println(true)

	// }
	// var a string = "sdfajksldfs"
	// b := []byte(a)
	// fmt.Printf("%b", b)
	// fmt.Println()

	// v := 30
	// fmt.Printf("%x", v)
	fmt.Println(decimaltoHex([]string{"a;lokdf", "10", "(dec)", "slkjd", "sdfj"}))
}
