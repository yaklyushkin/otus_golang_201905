package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const (
	COMMON = iota
	SLASH
	COUNT
)

func unzip(str string) string {
	if len(str) == 0 {
		return ""
	}
	var runes = []rune(str)
	if unicode.IsDigit(runes[0]) {
		log.Fatalf("Неверная строка: начинается с цифры '%s'", string(runes[0]))
	}
	var result strings.Builder
	var state int = COMMON
	var counts int = 0
	var prevChar rune = 0
	var dummy int
	for _, strChar := range runes {
		switch {
		case strChar == '\\':
			{
				if state == SLASH {
					prevChar = '\\'
					state = COMMON
				} else if state == COMMON {
					if prevChar != 0 {
						result.WriteRune(prevChar)
					}
					prevChar = 0
					state = SLASH
				} else {
					// COUNT
					for i := 0; i < counts; i++ {
						result.WriteRune(prevChar)
					}
					prevChar = 0
					counts = 0
					state = SLASH
				}
			}
		case unicode.IsDigit(strChar):
			{
				if state == SLASH {
					prevChar = strChar
					state = COMMON
				} else if state == COMMON {
					counts, _ = strconv.Atoi(string(strChar))
					state = COUNT
				} else {
					// COUNT
					dummy, _ = strconv.Atoi(string(strChar))
					counts = counts*10 + dummy
				}
			}
		default:
			{
				// COMMON
				if state == SLASH {
					result.WriteRune('\\')
					prevChar = strChar
					state = COMMON
				} else if state == COMMON {
					if prevChar != 0 {
						result.WriteRune(prevChar)
					}
					prevChar = strChar
				} else {
					// COUNT
					for i := 0; i < counts; i++ {
						result.WriteRune(prevChar)
					}
					prevChar = strChar
					counts = 0
					state = COMMON
				}
			}
		}
	}
	if state == SLASH {
		result.WriteRune('\\')
	} else if state == COMMON {
		if prevChar != 0 {
			result.WriteRune(prevChar)
		}
	} else {
		// COUNT
		for i := 0; i < counts; i++ {
			result.WriteRune(prevChar)
		}
	}
	return result.String()
}

func main() {
	/*
	* "a4bc2d5e" => "aaaabccddddde"
	* "abcd" => "abcd"
	* "45" => "" (некорректная строка)
	* `qwe\4\5` => `qwe45` (*)
	* `qwe\45` => `qwe44444` (*)
	* `qwe\\5` => `qwe\\\\\` (*)
	 */
	testData := [][]string{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
		{`\`, `\`},
		{`a\b`, `a\b`},
		{"45", ""},
	}
	for _, data := range testData {
		unzippedStr := unzip(data[0])
		isEqual := (unzippedStr == data[1])
		fmt.Println(
			strconv.FormatBool(isEqual) +
				": '" + data[0] + "' -> '" + data[1] +
				"' unzip = '" + unzip(data[0]) + "'")
	}
}
