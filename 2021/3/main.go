package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	numMap := make(map[byte]int)

	for scanner.Scan() {
		for index, digit := range scanner.Text() {
			conv_digit, conv_err := strconv.Atoi(string(digit))
			if conv_err != nil {
				fmt.Println(conv_err)
				return
			}
			numMap[byte(index)] += conv_digit
		}
	}
	fmt.Print(numMap)
}

func calculate_rates(mp map[byte]int) (int, int, error) {
	var epsilon int
	var gamma int

}
