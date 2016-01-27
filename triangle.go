package main

// TriangleType is the type name of the a Triangle
type TriangleType string

const (
	// EQUILATERAL - all sides equals
	EQUILATERAL TriangleType = "Equilateral"
	// ISOCELES - Two sides equals
	ISOCELES TriangleType = "Isoceles"
	// SCALENE - No sides equals
	SCALENE TriangleType = "Scalene"
	// INVALID - if triangle is invalid show this message
	INVALID TriangleType = "Your input is invalid to triangle"
)

// Triangle is the struct with information about size sides of the a Triangle
type Triangle struct {
	side1, side2, side3 int64
}

// NewTriangle create an instance of the Triangle
// input side1, side2, side3 int64
func NewTriangle(sides []int64) *Triangle {
	return &Triangle{side1: sides[0], side2: sides[1], side3: sides[2]}
}

// GetType return the name type of the Triangle
func (t *Triangle) GetType() TriangleType {
	if t.side1 <= 0 || t.side2 <= 0 || t.side3 <= 0 || (t.side1 >= t.side2+t.side3 || t.side2 >= t.side1+t.side3 || t.side3 >= t.side1+t.side2) {
		return INVALID
	}

	if t.side1 == t.side2 && t.side2 == t.side3 {
		return EQUILATERAL
	}

	if t.side1 == t.side2 || t.side2 == t.side3 || t.side3 == t.side1 {
		return ISOCELES
	}

	return SCALENE
}
