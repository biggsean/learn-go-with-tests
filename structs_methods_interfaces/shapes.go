package main

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle representation
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle representation
type Circle struct {
	Radius float64
}

// Area returns area
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Triangle representation
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns area
func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

// Perimeter returns the perimeter
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area returns the area
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
