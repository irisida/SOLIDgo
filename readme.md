# SOLID principles in Go

SOLID principles have a strong root in OOP, of which Go is not a true OOP language so there may be some terms slightly co-opted (highlighted of course) to explain or keep to the essence of the reason behind a given principle.

- `S` - SRP - Single responsibility principle
- `O` - OCP - Open closed principle
- `L` - LSP - Liskov substitution principle
- `I` - ISP - Interface segregation principle
- `D` - DIP - Dependency inversion principle

## SRP - Single responsibility principle

A type should have one responsibility and therefore only one reason to change. In definition terms we typically see words like `class`, `module` and `encapsulate`. Working within the binds of `Go` we should assume that the most direct correlation to class is a struct and therefore:

- A struct should have a single purpose or functionality.
- All functionality of the struct should be strictly, narrowly aligned to that functionality/purpose.
- cross cutting functionality extensions should be avoided or extracted out to another struct. (separation of concerns)
-

## OCP - Open closed principle

## LSP - Liskov substitution principle

## ISP - Interface segregation principle

## DIP - Dependency inversion principle
