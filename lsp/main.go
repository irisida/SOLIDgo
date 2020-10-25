package main

import "fmt"

// Sized interface
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// Rectangle struct
type Rectangle struct {
	width, height int
}

// GetWidth - getter for the width value
func (r *Rectangle) GetWidth() int {
	return r.width
}

// SetWidth sets the width
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

// GetHeight - getter for the height value
func (r *Rectangle) GetHeight() int {
	return r.height
}

// SetHeight sets height value
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

/*
here is where we break the LSP because we're starting to
implement functionality that deviates away from the lower
or base implementation and this causes fragility and some
very unexpected results
*/

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	// the the breaks the principle
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	// the the breaks the principle
	s.width = height
	s.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Printf("expected: %v  actual: %v \n", expectedArea, actualArea)
}

func main() {
	rec := &Rectangle{2, 3}
	UseIt(rec)

	// violates the LSP
	sq := NewSquare(5)
	UseIt(sq)

	// an alternate could be a struct of type square that
	// takes a single parameter, inside that function it
	// generates a new Rectangle and assigns the single
	// param size to both length and width.
}
