package main

import (
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
}
