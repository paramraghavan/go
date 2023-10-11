package jsontostruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func Run() {
	e := getEmployee()
	empJsonStr, err := structToJson(e)

	if err != nil {
		fmt.Println("Error struct to json string conversion")
	} else {
		fmt.Println(string(empJsonStr))
	}

	e1, err := jsonStrToStruct(empJsonStr)

	fmt.Println(e1)

}

func jsonStrToStruct(jsonStr string) (Employee, error) {
	data := Employee{}
	err := json.Unmarshal([]byte(jsonStr), &data)

	if err != nil {
		log.Println(err)
		msg := fmt.Sprintf("Error parsing json string to struct, %s", err)
		return data, errors.New(msg)
	}
	return data, nil
}

func structToJson(e Employee) (string, error) {
	b, err := json.Marshal(e)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}
	return string(b), nil
}

func getEmployee() Employee {
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	return e
}
