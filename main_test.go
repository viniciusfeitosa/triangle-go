package main

import (
	"reflect"
	"testing"
)

func TestGetSidesFromInputWithInvalidValue(t *testing.T) {
	input := "1,1"
	_, err := getSidesFromInput(input)
	if err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}

	input = "a a"
	if _, err := getSidesFromInput(input); err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}

	input = "1 1"
	if _, err := getSidesFromInput(input); err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}

	input = " "
	if _, err := getSidesFromInput(input); err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}

	input = "a"
	if _, err := getSidesFromInput(input); err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}

	input = "1,2,3,4"
	if _, err := getSidesFromInput(input); err.Error() != ErrorInvalidInput {
		t.Error("This test should return an Invalid input message")
	}
}

func TestGetSidesFromInputWithValidValue(t *testing.T) {
	expected := []int64{5, 5, 5}
	input := "5,5,5"
	sides, err := getSidesFromInput(input)
	if !reflect.DeepEqual(expected, sides) && err != nil {
		t.Errorf("This test should return an slide like %v, but returned %v\n", expected, sides)
	}

	expected = []int64{5, 2, 5}
	input = "5,5,5"
	sides, err = getSidesFromInput(input)
	if !reflect.DeepEqual(expected, sides) && err != nil {
		t.Errorf("This test should return an slide like %v, but returned %v\n", expected, sides)
	}
}
