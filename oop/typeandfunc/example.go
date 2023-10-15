package typeandfunc

import "fmt"

// Emp user-defined type
type EmpSample int

/*
* Here we cam pass EmpSample by value or pointer.
 1. If the EmpSample is mutable, then you can poss it as pointer
 2. If the type object is very big then also you can pass by pointer
*/
func (e *EmpSample) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

func (e *EmpSample) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}
