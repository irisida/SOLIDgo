package main

import "fmt"

type Document struct {
	name string
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {}
func (m MultiFunctionPrinter) Fax(d Document)   {}
func (m MultiFunctionPrinter) Scan(d Document)  {}

// ******************
// break the principle
type OldFashionedPrinter struct{}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}
func (o OldFashionedPrinter) Fax(d Document)  { /* wont have capability */ }
func (o OldFashionedPrinter) Scan(d Document) { /* wont have capability */ }

// ***************************************
// proposed solution to work the principle
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}
type Fax interface {
	Fax(d Document)
}

// use the ISP friendly implementation that implements only print interface
type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {
	// implement the interface
}

// implement many with ISP
type MyCopier struct{}

func (m MyCopier) Print(d Document) {
	// implement the interface
	fmt.Println("Document is: ", d.name)
	fmt.Println("I am printing")
}
func (m MyCopier) Scan(d Document) {
	// implement the interface
	fmt.Println("I am scanning")
}

func main() {
	d := Document{name: "ed.txt"}
	m := MyCopier{}
	m.Print(d)
}
