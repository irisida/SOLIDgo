package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low level module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high level module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate(lookupName string) {
	for _, p := range r.browser.FindAllChildrenOf(lookupName) {
		fmt.Println(lookupName, "has a child named ", p.name)
	}
}

func main() {
	parent := Person{"Jock"}
	ch1 := Person{"Chris"}
	ch2 := Person{"Lonnie"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &ch1)
	relationships.AddParentAndChild(&parent, &ch2)

	r := Research{&relationships}
	r.Investigate(parent.name)
}
