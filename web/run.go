package main

import "fmt"
import (
	ws "web/first"
)

func main() {
	ws.StartHelloWebServer()
	fmt.Println("starting web server...")
}
