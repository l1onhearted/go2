package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (self *Rectangle) Area() float32 {
	return self.length * self.width
}
func main() {
	sq1 := new(Square)
	sq1.side = 5
	r1 := new(Rectangle)
	r1.length = 3
	r1.width = 7
	var areaIntf interface{}
	areaIntf = *sq1
	areaRe := r1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Println(areaIntf)
	// fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Println("rectangle area:", areaRe.Area())
}
