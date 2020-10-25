package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	demo the SRP (single responsibility principle) where we adhere
	to the mantra of one reason to exist and one reason to change
	this avoiding breaking the principle and creating god objects
	that have multiple reasons (functionalities) and therefore
	many reasons that could trigger change.

	We have a package dealing with the management of journal entries
	if we were to add how persistence worked with a journal to this
	package we would now break the SRP as we have two main reasons
	to exits and multiple reasons for change.
*/

var entryCount = 0

// Journal struct is responsible for managing
// the entries of a journal
type Journal struct {
	entries []string
}

// Persistence is responsible for reading from and
// writing to files, a crosscurting responsibility
// and therefore separated out
type Persistence struct {
	LineSeparator string
}

// AddEntry - adds to journal
func (j *Journal) AddEntry(txt string) int {
	entryCount++
	newEntry := fmt.Sprintf("%d, %s", entryCount, txt)
	j.entries = append(j.entries, newEntry)
	return entryCount
}

// DeleteEntry has not been implemented
func (j *Journal) DeleteEntry(index int) {
	// remove entry
}

// String prints entries separated by a newline
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// SaveToFile Separation of concerns
func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

// func (j *Journal) Load (filename string) {
// 	 ...
// }

// func (j *Journal) LoadFromWeb (filename string) {
// 	...
// }

func main() {
	j := Journal{}
	j.AddEntry("I created a journal")
	j.AddEntry("I added some stuff to it")
	fmt.Println(j.String())

	p := Persistence{"\n"}
	p.SaveToFile(&j, "./testoutput.txt")
}
