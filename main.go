package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	// ErrorInvalidInput is the message when input is wrong
	ErrorInvalidInput = "Invalid input"
)

func main() {
	var input string

	fmt.Println("Please enter the 3 sides values to determinate the triangle type")
	fmt.Print("Split values with commas: ")
	if _, err := fmt.Scanln(&input); err != nil {
		log.Fatalf("Input error: %v", err)
	}

	sides, err := getSidesFromInput(input)
	if err != nil {
		log.Fatalf("get sides error: %v", err)
	}

	fmt.Println(NewTriangle(sides).GetType())
}

func getSidesFromInput(input string) ([]int64, error) {
	sliceSidesString := strings.Split(input, ",")
	if len(sliceSidesString) != 3 {
		return nil, errors.New(ErrorInvalidInput)
	}

	sides := make([]int64, 3)
	for index, value := range sliceSidesString {
		side, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}

		sides[index] = side
	}

	return sides, nil
}
