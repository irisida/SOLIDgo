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

The open closed principles dictates that software entities (structs in the realms of Go) should be open to extension but closed to modification. This aligns well with the `specification pattern`.

## LSP - Liskov substitution principle

Named after Barbara Liskov, the Liskov substitution principle deals primarily with inheritance and is therefore not directly applicable to Go. The principle dictates that if you have an object of a derived class it should also work correctly with the base class, or that a base class should not have implemented logic that is negated in a class derived from it. They should be interchangeable in the functionality that is covered in the base class.

## ISP - Interface segregation principle

The simplest of principles in the SOLID design principles. The idea is that an interface should not contain more than it needs to. It makes more sense to break up a large interface to smaller and more reusable interfaces instead of a large interface that would force objects implementing it to needlessly implement unused methods.

## DIP - Dependency inversion principle
