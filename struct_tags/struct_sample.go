package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation,omitempty"`
}

func (p User) String() string {

	return fmt.Sprintf("User id=%v, name=%v, occupation=%v",
		p.Id, p.Name, p.Occupation)
}

type User1 struct {
	FirstName string `json:"first_name"`
	BirthYear int    `json:"birth_year"`
	Email     string
	Tax       float64 `json:"tax,omitempty"`
}

func marshallEx() {
	dat, _ := json.Marshal(`User1{
    FirstName: "Lane",
    BirthYear: 1990,
    Email:     "example@gmail.com",
}`)
	fmt.Println(string(dat))
}

func unMarshallEx() {
	dat := []byte(`{
    "first_name":"Lane",
    "birth_year":1990,
    "Email":"example@gmail.com"
}`)
	user := User1{}
	err := json.Unmarshal(dat, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

type Test struct {
	A string
	B string
	C string
}

func unMarshallEx1() {
	var example []byte = []byte(`{"A": "1", "C": "3"}`)

	out := Test{
		A: "default a",
		B: "default b",
		// default for C will be "", the empty value for a string
	}
	err := json.Unmarshal(example, &out) // <--
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", out)
}

func main() {

	user := User{Id: 1, Name: "John Doe", Occupation: "gardener"}
	res, _ := json.MarshalIndent(user, " ", "  ")

	fmt.Println(string(res))

	user2 := User{Id: 1, Name: "John Doe"}
	res2, _ := json.MarshalIndent(user2, " ", "  ")

	fmt.Println(string(res2))

	//
	marshallEx()
	//
	unMarshallEx()
	//
	unMarshallEx1()

}
