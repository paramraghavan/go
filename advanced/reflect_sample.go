package main

//TODO work in progress

import (
	"fmt"
	"reflect"
	"strconv"
)

type T struct {
	A int
	B string
	C float64
	D bool
	student Student
}

type Student struct {
	name   string
	branch string
	year   int
}

type ColType reflect.Type

func main() {
	sample()

	var t interface{} = T{}
	mapOfT := mapOfColAndType(t)
	fmt.Println(mapOfT)
}

func sample() {
	t := T{10, "ABCD", 15.20, true, Student{}}
	t.student.year = 2023
	t.student.branch = "BS"
	t.student.name = " joy"
	typeT := reflect.TypeOf(t)

	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}

func parseMapTToStructT( mapData map[string]string, mapOfT map[string]ColType) T {
	var parsedDataStruct = T{}

	for k,v := range mapData {
		strVar := v
		switch mapOfT[k].String() {
			case "string":
			case "int":
				parsedDataStruct.`k` = 	strconv.Atoi(strVar)
			case "float64":
				parsedDataStruct[k] = 	strconv.ParseFloat(strVar, 64)
			case "float32":
				parsedDataStruct[k] = 	strconv.ParseFloat(strVar, 32)
			case "bool":
				parsedDataStruct[k] = 	strconv.ParseBool(strVar)
		}
	}

	return parsedDataStruct

}
func generateMapOfT() map[string]string {
	stuctTMap := map[string]string {}

	stuctTMap["A"] = "10"
	stuctTMap["B"] = "String Value"
	stuctTMap["C"] = "10.00"
	stuctTMap["D"] = "true"

	return stuctTMap
}

func  mapOfColAndType(data interface{} ) map[string]ColType {
	typeT := reflect.TypeOf(data)

	var mapValTyp map[string]ColType = map[string]ColType{}
	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		mapValTyp[field.Name] = field.Type
	}
	return mapValTyp

}
