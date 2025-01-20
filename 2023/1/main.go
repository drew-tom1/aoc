package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	var res int
	count := 1

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {

		var temp string

		raw_input, input_err := normalizeText(scanner.Text())
		if input_err != nil {
			fmt.Println(input_err)
			return
		}

		left, err1 := leftCalibrate(string(raw_input))
		fmt.Println("LINE NUMBER: ", count)
		fmt.Println("LEFT", string(left))
		right, err2 := rightCalibrate(string(raw_input))
		fmt.Println("RIGHT", string(right))

		if err1 == nil {
			temp += string(left)
		}
		if err2 == nil {
			temp += string(right)
		}

		intermediate, err := strconv.Atoi(temp)
		fmt.Println("INTERMEDIATE VAL", intermediate)
		if err != nil {
			fmt.Println("Error occurred: ", err)
		}

		res += intermediate

		temp = ""
		count++

	}
	fmt.Println(res)

}

func leftCalibrate(line string) (byte, error) {
	for i := 0; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return line[i], nil
		}
	}
	return 0, errors.New("no number detected")
}

func rightCalibrate(line string) (byte, error) {
	for i := len(line) - 1; i >= 0; i-- {
		_, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return line[i], nil
		}
	}
	return 0, errors.New("no number detected")
}

func normalizeText(line string) (string, error) {
	mp := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for word, digit := range mp {
		re := regexp.MustCompile(`(?i)` + word)
		line = re.ReplaceAllString(line, digit)
	}
	return line, nil
}
