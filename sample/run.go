package main

import "fmt"

import (
	e "sample/emp"
)

func main() {
	var e1 e.Employee
	e1 = e.Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
