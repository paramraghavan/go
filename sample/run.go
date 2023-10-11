package main

import "fmt"

import (
	e "sample/emp"
	js "sample/jsontostruct"
)

func main() {
	var e1 e.Employee
	e1 = e.Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))

	// json str to struct and bacj
	js.Run()
}
