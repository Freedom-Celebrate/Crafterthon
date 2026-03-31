package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Removetodo(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		if words[i] == "TODO" {
			words[i] = "ACTION"
		}
	}
	return strings.Join(words, " ")
}

func Trim(s string) string {

	p1 := regexp.MustCompile(`^\s+`)
	p2 := regexp.MustCompile(`$\s+`)

	s = p1.ReplaceAllString(s, "")
	s = p2.ReplaceAllString(s, "")
	return s

}
func Title(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		for _, ch := range words {
			if ch >= "A" && ch <= "Z" {
				words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])

			}
		}

	}
	return strings.Join(words, " ")
}
func Flag(s string) string {
	if len(s) > 80 {
		return s + "TRUNCATED"
	}
	return s

}
func DashesBlanks(s string) string {
	p1 := regexp.MustCompile(`\s+`)
	p2 := regexp.MustCompile(`-|_`)

	s = p1.ReplaceAllString(s, "")
	s = p2.ReplaceAllString(s, "")
	return s
}

func main() {

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File not found: input.txt")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(os.Stdout)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		
	}
	if  err := scanner.Err(); err != nil{
		fmt.Println(err)
	}

}
