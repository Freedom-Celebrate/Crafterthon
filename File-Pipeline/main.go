package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CapsToTitle(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == strings.ToUpper(words[i]) {
			words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])

		}

	}
	return strings.Join(words, " ")
}
func LowToUpper(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == strings.ToLower(words[i]) {
			words[i] = strings.ToUpper(string(words[i]))

		}

	}
	return strings.Join(words, " ")
}
func Flag(s string) string {
	if len(s) > 80 {
		return s + " " + "[TRUNCATED]"
	}
	return s
}
func Trim(s string) string {
	p1 := regexp.MustCompile(`^\s+`)
	p2 := regexp.MustCompile(`$\s+`)

	s = p1.ReplaceAllString(s, "")
	s = p2.ReplaceAllString(s, "")
	return s
}
func DashesAndBlanks(s string) string {
	p1 := regexp.MustCompile(`s+`)
	p2 := regexp.MustCompile(`-`)
	p3 := regexp.MustCompile(`_`)

	s = p1.ReplaceAllString(s, "")
	s = p2.ReplaceAllString(s, "")
	s = p3.ReplaceAllString(s, "")
	return s
}
func Todo(s string) string {
	words := strings.Fields(s)
	if words[0] == strings.ToUpper("TODO:") || words[0] == strings.ToLower("todo:") || words[0] == strings.Title("Todo:") {
		words[0] = strings.ToUpper("✦ ACTION:")

	}
	return strings.Join(words, " ")
}

func process(s string) string {
	s = Trim(s)
	s = DashesAndBlanks(s)
	s = CapsToTitle(s)
	s = LowToUpper(s)
	s = Flag(s)
	s = Todo(s)

	return s
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
	}

	outputfile := os.Args[2]

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not open file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanner.Text()
		result := process(scanner.Text())

		err = os.WriteFile(outputfile, []byte(result+"\n"), 0644)
		if err != nil {
			fmt.Println("unsuccessful file writing:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input: ", err)
	}

}
