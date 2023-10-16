# OOPs

## Go user defined types
- Examples
```go
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

type Address struct {
      name string 
      street string
      city string
      state string
      Pincode int
}

type Emp int

```

## Add behaviour to types
- adding methods
- implementing interface

### methods
Add method to struct Employee
```go
type Employee struct {
	Name string
	Age  int
    Active bool
}

// this could be pass by value or pointer.
// here it is pass by value
func (e Employee) PrintName() {
	fmt.Println("Employee Age:\t", e.Age)
	fmt.Println("Employee Name:\t", e.Name)
}

// here it is pass by value
func (e Employee) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

// here pass as pointer as i want to update employee status to active
func (e *Employee) setActiveStatus()  {
    *e.Active = true
}

```

### interface
Think about a printer interface that can print Employee as well as Address struct above

```go
import "fmt"

type printLabel interface {
	Print() string
}

type employee struct {
	FirstName, LastName, Email string
	Age                        int
}

type address struct {
	street  string
	city    string
	state   string
	zipcode int
}

func (e employee) Print() string {
	return fmt.Sprintf("%v %v [%v]\n", e.FirstName, e.LastName, e.Email)
}

func (a address) Print() string {
	return fmt.Sprintf("%v %v %v [%d]\n", a.street, a.city, a.state, a.zipcode)
}

func run() {
	var p printLabel
	p = employee{FirstName: "Anita", LastName: "Worcester", Age: 27, Email: "abc@email.com"}
	fmt.Println(p.Print())

	p = address{street: "123 main street", city: "Worcester", state: "MA", zipcode: 1601}
	fmt.Println(p.Print())

	// checking if interface is employee or address
    // type assertions
	e, ok := p.(employee)
	fmt.Println(e, ok) // it's not employee type, ok is false

	a, ok := p.(address)
	fmt.Println(a, ok) // it's address type, pk is true

	switch v := p.(type) {
	case employee:
		fmt.Println("Found employee!", v)
	case address:
		fmt.Println("Found address!", v)
	default:
		fmt.Println("I'm not sure what this is...")
	}
}

/* Output
Anita Worcester [abc@email.com]

123 main street Worcester MA [1601]

{   0} false
{123 main street Worcester MA 1601} true
Found address!
```

### Generics
Interfaces are very powerful constructs in Go. However, they do have a downside, and that is, if we look at this example above, whenever we assigned a concrete type, so for example, employee or address object to printLabel  interface, it loses its identity. We no longer know for sure what concrete type we are working with - employee or address. 

In the example above we got back to the concrete type using type assertions and type switches, but those tend to be fairly heavy options for us to use. We want a family of concrete types to have some polymorphic behavior so they all work in the same way, but then after we're done with that polymorphic behavior, we'd like to get back to that concrete type. Well, interfaces don't offer us a simple way to do that, so instead, we're going to turn to another polymorphic construct that Go contains called generic, or generic programming.
