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
}

type ColType reflect.Type

func main() {
	var t interface{} = T{}
	mapOfT := mapOfColAndType(t)
	fmt.Println(mapOfT)


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
