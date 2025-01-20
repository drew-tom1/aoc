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

	var res int
	var depth int
	var pos int

	for scanner.Scan() {
		output := strings.Split(scanner.Text(), " ")
		direction := output[0]
		value, conv_err := strconv.Atoi(output[1])
		if conv_err != nil {
			fmt.Println(conv_err)
			return
		}

		switch direction {
		case "forward":
			pos += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	res = depth * pos
	fmt.Println(res)
}
