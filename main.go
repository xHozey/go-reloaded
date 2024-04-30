package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sample := os.Args[1]
	data, err := os.ReadFile(sample)
	if err != nil {
		fmt.Printf("Error")
		return
	}

	text := string(data)
	fmt.Print(text)
	text = " " + text + " "
	text = strings.ReplaceAll(text, "''", " ' ' ")
	text = strings.ReplaceAll(text, "' ", " ' ")
	text = strings.ReplaceAll(text, " '", " ' ")
	text = strings.ReplaceAll(text, "'\n", " '  \n ")
	text = strings.ReplaceAll(text, "\n'", "   '   \n ")
	text = strings.ReplaceAll(text, "''", " ' ")
	text = strings.ReplaceAll(text, "' ", " ' ")
	text = strings.ReplaceAll(text, " '", " ' ")
	text = strings.ReplaceAll(text, "'\n", " ' \n ")
	text = strings.ReplaceAll(text, "\n'", " ' \n ")
	text = strings.ReplaceAll(text, "\n", "\n ")
	text = strings.TrimSpace(text)
	text = ponctuationFix(text)
	input := strings.Split(text, " ")

	for i := 0; i < len(input); i++ {
		if i == 0 {
			match := false
			match1, _ := regexp.MatchString(`^\((hex|bin|cap|up|low)[,\)]$`, input[i])
			if i+1 < len(input) {
				match2, _ := regexp.MatchString(`^\d+\)$`, input[i+1])
				match = match1 && match2
			}
			if match {
				input[i] = ""
				input[i+1] = ""
			} else if (input[i] == "(up)" || input[i] == "(low)" || input[i] == "(cap)") && match1 {
				input[i] = ""
			}

		} else {
			var numberOfConvert int
			switch input[i] {
			case "(hex)":
				input[i] = ""
				input[i-1] = hexConvert(input[i-1])
			case "(bin)":
				input[i] = ""
				input[i-1] = binConvert(input[i-1])
			case "(cap)":
				input[i] = ""
				input[i-1] = Capitalize(input[i-1])
			case "(cap,":
				if i < len(input)-1 && input[i+1][len(input[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(input[i+1])
					if numberOfConvert == -1 {
						continue
					}
					input[i] = ""
					input[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if input[j] == "" {
							continue
						}
						input[j] = Capitalize(input[j])
						numberOfConvert--
					}
				}
			case "(up)":
				input[i] = ""
				input[i-1] = strings.ToUpper(input[i-1])
			case "(up,":
				if i < len(input)-1 && input[i+1][len(input[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(input[i+1])
					if numberOfConvert == -1 {
						continue
					}
					input[i] = ""
					input[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if input[j] == "" {
							continue
						}
						input[j] = strings.ToUpper(input[j])
						numberOfConvert--
					}

				}

			case "(low)":
				input[i] = ""
				input[i-1] = strings.ToLower(input[i-1])
			case "(low,":
				if i < len(input)-1 && input[i+1][len(input[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(input[i+1])
					if numberOfConvert == -1 {
						continue
					}
					input[i] = ""
					input[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if input[j] == "" {
							continue
						}
						input[j] = strings.ToLower(input[j])
						numberOfConvert--
					}
				}
			}
		}
	}

	input = quotesFixer(input)
	input = vowl(input)
	myString := Join(input, " ")
	myString = ponctuationFix(myString)
	myString = standardizeSpaces(myString)

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Printf("error")
		return
	}
	file.WriteString(myString)
}

func Join(strs []string, sep string) string {
	var h string
	for _, arg := range strs {
		if h == "" {
			h = arg
		} else if h != "\n" {
			h = h + sep + arg
		} else if h == "\n" {
			h = h + arg
		}
	}
	return h
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Split(s, " "), " ")
}

func CheckFirstCharIsvowel(str string) bool {
	vowls := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i := 0; i < len(vowls); i++ {
		if len(str) > 0 {
			if string(str[0]) == vowls[i] {
				return true
			}
		}
	}
	return false
}

func vowl(s []string) []string {
	for i := range s {
		if s[i] == "a" || s[i] == "A" {
			if i+1 < len(s) {
				if CheckFirstCharIsvowel(s[i+1]) {
					s[i] += "n"
				}
			}
		}
		if s[i] == "'a" || s[i] == "'A" {
			if i+1 < len(s) {
				if CheckFirstCharIsvowel(s[i+1]) {
					s[i] += "n"
				}
			}
		}
		if s[i] == "an" || s[i] == "An" {
			if i+1 < len(s) {
				if !CheckFirstCharIsvowel(s[i+1]) {
					s[i] = "a"
				}
			}
		}
		if s[i] == "'an" || s[i] == "'An" {
			if i+1 < len(s) {
				if !CheckFirstCharIsvowel(s[i+1]) {
					s[i] = "'a"
				}
			}
		}
	}

	return s
}

func quotesFixer(s []string) []string {
	offOn := false

	for i := 0; i < len(s); i++ {

		if s[i] == "'" && !offOn {
			offOn = true
			s[i+1] = "'" + s[i+1]
			s[i] = ""

		} else if s[i] == "'" && offOn {
			offOn = false
			s[i-1] = s[i-1] + "'"
			s[i] = ""
		}
	}
	return s
}

func ponctuationFix(s string) string {
	search := regexp.MustCompile(` *([\.,!\?\:;]+) *(\w)`)
	s = search.ReplaceAllString(s, "$1 $2")

	search2 := regexp.MustCompile(` *([\.,!\?\:;]+)`)
	s = search2.ReplaceAllString(s, "$1")
	return s
}

func TrimAtoi(s string) int {
	r := 0
	N := false
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			r = r*10 + int(s[i]-'0')
		} else if s[i] == '-' && r == 0 {
			N = true
		} else if s[i] < '0' || s[i] > '9' {
			r = -1
		}
	}
	if N {
		r *= -1
	}
	return r
}

func hexConvert(s string) string {
	hex, _ := strconv.ParseInt(s, 16, 64)
	conv := strconv.Itoa(int(hex))
	return conv
}

func binConvert(s string) string {
	hex, _ := strconv.ParseInt(s, 2, 64)
	conv := strconv.Itoa(int(hex))
	return conv
}

func Capitalize(s string) string {
	s = strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	return s
}
