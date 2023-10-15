package methods

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

	// checking if inerface is employee or address
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
Found address! {123 main street Worcester MA 1601}
*/
