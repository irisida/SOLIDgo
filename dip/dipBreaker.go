package main

// import "fmt"

// type Relationship int

// const (
// 	Parent Relationship = iota
// 	Child
// 	Sibling
// )

// type Person struct {
// 	name string
// }

// type Info struct {
// 	from         *Person
// 	relationship Relationship
// 	to           *Person
// }

// // low level module
// type Relationships struct {
// 	relations []Info
// }

// func (r *Relationships) AddParentAndChild(parent, child *Person) {
// 	r.relations = append(r.relations, Info{parent, Parent, child})
// 	r.relations = append(r.relations, Info{child, Child, parent})
// }

// // high level module
// type Research struct {
// 	// break the DIP
// 	relationships Relationships
// }

// func (r *Research) Investigate(lookupName string) {
// 	relations := r.relationships.relations
// 	for _, rel := range relations {
// 		if rel.from.name == lookupName && rel.relationship == Parent {
// 			fmt.Println(rel.from.name, " has a child named ", rel.to.name)
// 		}
// 	}
// }

// func main() {
// 	parent := Person{"Jock"}
// 	ch1 := Person{"Chris"}
// 	ch2 := Person{"Lonnie"}

// 	relationships := Relationships{}
// 	relationships.AddParentAndChild(&parent, &ch1)
// 	relationships.AddParentAndChild(&parent, &ch2)

// 	r := Research{relationships}
// 	r.Investigate(parent.name)
// }
