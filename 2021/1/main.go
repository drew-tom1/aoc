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
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	var temp int
	var res int

	for scanner.Scan() {
		line_num, conv_err := strconv.Atoi(scanner.Text())
		if conv_err != nil {
			fmt.Println(conv_err)
			return
		}

		if line_num > temp {
			res++
		}
		temp = line_num
	}
	fmt.Println(res - 1) // the first line was added but it should be registered as N/A and is subtracted
}
