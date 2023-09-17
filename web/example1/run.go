package main

import "fmt"
import (
	ws "example1/first"
)

func main() {
	ws.StartHelloWebServer()
	fmt.Println("starting web server...")
}
