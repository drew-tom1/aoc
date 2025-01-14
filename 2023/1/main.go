package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	var res int

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {

		var temp string

		left, err1 := leftCalibrate(string(scanner.Text()))
		right, err2 := rightCalibrate(string(scanner.Text()))

		if err1 == nil {
			temp += string(left)
		}
		if err2 == nil {
			temp += string(right)
		}

		intermediate, err := strconv.Atoi(temp)
		if err != nil {
			fmt.Println("Error occurred: ", err)
		}

		res += intermediate
		fmt.Println(res)

		temp = ""

	}

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
