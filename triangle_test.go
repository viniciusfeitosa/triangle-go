package main

import "testing"

func TestIfTriangleIsValidWithNegativeSidesOrZeroValue(t *testing.T) {
	sides := []int64{5, 0, 5}
	triangle := NewTriangle(sides)
	if triangle.GetType() != INVALID {
		t.Error("This test should return an invalide message")
	}

	sides = []int64{5, -1, 5}
	triangle = NewTriangle(sides)
	if triangle.GetType() != INVALID {
		t.Error("This test should return an invalide message")
	}
}

func TestIfTriangleIsOK(t *testing.T) {
	sides := []int64{5, 5, 5}
	triangle := NewTriangle(sides)
	if triangle.GetType() != EQUILATERAL {
		t.Errorf("This test should return EQUILATERAL, but returned %v \n", triangle.GetType())
	}

	sides = []int64{5, 6, 5}
	triangle = NewTriangle(sides)
	if triangle.GetType() != ISOCELES {
		t.Errorf("This test should return ISOCELES, but returned %v \n", triangle.GetType())
	}

	sides = []int64{4, 6, 5}
	triangle = NewTriangle(sides)
	if triangle.GetType() != SCALENE {
		t.Errorf("This test should return SCALENE, but returned %v \n", triangle.GetType())
	}
}
