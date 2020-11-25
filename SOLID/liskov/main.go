package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func testIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	// rectangle
	expectedResult := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedResult, ", but got ", actualArea, "\n")
}

func main() {
	rectangle := &Rectangle{2,3}
	testIt(rectangle)
}

