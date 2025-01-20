package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var res int
	var power_sum int
	count := 1

	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error occurred with file")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {

		id, games, sanitize_err := sanitizeData(scanner.Text())

		if sanitize_err != nil {
			fmt.Println(sanitize_err)
			return
		}

		isValid, cube_nums := check_valid(games)

		if isValid {
			res += id
		}
		if cube_nums != nil {
			temp := cube_nums["green"] * cube_nums["blue"] * cube_nums["red"]
			power_sum += temp
		}

	}
	fmt.Println(res)
	fmt.Println(power_sum)
}

func check_valid(line string) (bool, map[string]int) {

	isCorrect := true

	benchmark := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	max_cubes := map[string]int{
		"green": 0,
		"red":   0,
		"blue":  0,
	}

	matches := strings.Split(line, ";") // each match within a game

	for _, match := range matches { // iterating through each match
		fmt.Println(match)

		colorMap := make(map[string]int) // hashmap of the colors and their number

		rolls := strings.Split(match, ",") // the rolls within a particular match

		for _, roll := range rolls {
			parts := strings.Split(roll, " ") // the parts of a roll (number and color)

			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println(err)
				isCorrect = false
			}

			color := parts[2]

			if value > benchmark[color] {
				fmt.Println("Value exceeds benchmark total")
				isCorrect = false
			}
			if colorMap[color]+value > benchmark[color] {
				fmt.Println("Color sum exceeds benchmark with added value")
				isCorrect = false
			}

			if value > max_cubes[color] {
				max_cubes[color] = value
			}

			colorMap[color] += value

		}
	}

	return isCorrect, max_cubes
}

func sanitizeData(line string) (int, string, error) {
	game_id, gameplay, _ := strings.Cut(line, ":")
	game_number, err := strconv.Atoi(game_id[5:])

	if err != nil {
		return -1, "", err
	}

	return game_number, gameplay, nil
}

func find_max_cubes() {

}
