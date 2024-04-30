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

	inputext := string(data)
	lines := strings.Split(inputext, "\n")
	for i, line := range lines {
		line := flags(line)
		line = quotesFixer(line)
		line = vowl(line)
		line = ponctuationFix(line)
		line = standardizeSpaces(line)
		lines[i] = line
	}
	output := strings.Join(lines, "\n")

	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Printf("error")
		return
	}
	file.WriteString(output)
}

func flags(str string) string {
	s := strings.Fields(str)

	for i := 0; i < len(s); i++ {
		if i == 0 {
			match := false
			match1, _ := regexp.MatchString(`^\((hex|bin|cap|up|low)[,\)]$`, s[i])
			if i+1 < len(s) {
				match2, _ := regexp.MatchString(`^\d+\)$`, s[i+1])
				match = match1 && match2
			}
			if match {
				s[i] = ""
				s[i+1] = ""
			} else if (s[i] == "(up)" || s[i] == "(low)" || s[i] == "(cap)") && match1 {
				s[i] = ""
			}

		} else {
			var numberOfConvert int
			switch s[i] {
			case "(hex)":
				s[i] = ""
				s[i-1] = hexConvert(s[i-1])
			case "(bin)":
				s[i] = ""
				s[i-1] = binConvert(s[i-1])
			case "(cap)":
				s[i] = ""
				s[i-1] = Capitalize(s[i-1])
			case "(cap,":
				if i < len(s)-1 && s[i+1][len(s[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(s[i+1])
					if numberOfConvert == -1 {
						continue
					}
					s[i] = ""
					s[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if s[j] == "" {
							continue
						}
						s[j] = Capitalize(s[j])
						numberOfConvert--
					}
				}
			case "(up)":
				s[i] = ""
				s[i-1] = strings.ToUpper(s[i-1])
			case "(up,":
				if i < len(s)-1 && s[i+1][len(s[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(s[i+1])
					if numberOfConvert == -1 {
						continue
					}
					s[i] = ""
					s[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if s[j] == "" {
							continue
						}
						s[j] = strings.ToUpper(s[j])
						numberOfConvert--
					}

				}

			case "(low)":
				s[i] = ""
				s[i-1] = strings.ToLower(s[i-1])
			case "(low,":
				if i < len(s)-1 && s[i+1][len(s[i+1])-1:] == ")" {
					numberOfConvert = TrimAtoi(s[i+1])
					if numberOfConvert == -1 {
						continue
					}
					s[i] = ""
					s[i+1] = ""
					for j := i - 1; numberOfConvert > 0; j-- {
						if j < 0 {
							break
						}
						if s[j] == "" {
							continue
						}
						s[j] = strings.ToLower(s[j])
						numberOfConvert--
					}
				}
			}
		}
	}
	return strings.Join(s, " ")

}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
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

func vowl(str string) string {
	s := strings.Fields(str)

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

	return strings.Join(s, " ")
}

func quotesFixer(text string) string {
	text = " " + text + " "
	text = strings.ReplaceAll(text, "''", " ' ' ")
	text = strings.ReplaceAll(text, "' ", " ' ")
	text = strings.ReplaceAll(text, " '", " ' ")
	text = strings.ReplaceAll(text, "''", " ' ")
	text = strings.ReplaceAll(text, "' ", " ' ")
	text = strings.ReplaceAll(text, " '", " ' ")
	text = strings.TrimSpace(text)

	s := strings.Fields(text)
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
	return strings.Join(s, " ")
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
