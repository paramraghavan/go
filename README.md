# go
go through the lens of a java/python developer. [Super link for go](https://www.golangprograms.com/)


## Is go compiled or interpreted. Is it VM based
- Go is a compiled language.
- This means we must run our source code files through a compiler, which reads source code and generates a binary, or executable, file that is used to run the program
- It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

## Does Go have a runtime?
*Go does have an extensive library, called the runtime*, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language. Although it is more central to the language, Go's runtime is analogous to libc , the C library.

## Goroutine vs Thread:
|Goroutine|	Thread|
|---------|--------|
|Goroutines are managed by the go runtime.|	Operating system threads are managed by kernel.|
|Goroutine are not hardware dependent.|	Threads are hardware dependent.|
|Goroutines have easy communication medium known as channel.|	Thread does not have easy communication medium.|
|Due to the presence of channel one goroutine can communicate with other goroutine with low latency.|	Due to lack of easy communication medium inter-threads communicate takes place with high latency.|
|Goroutines are cheaper than threads.|The cost of threads are higher than goroutine.|
|faster startup time than threads.| slow startup time than goroutines.|
|growable stack| fixed stack|

## Concurrency and Go
Supports two styles
  - Communicating sequential processes (CSP) *use communication as synchronization primitive*
  - Shared memory multithreading uses locks
  - goroutines communicate via channels
  > https://www.cs.princeton.edu/courses/archive/fall16/cos418/docs/P1-concurrency.pdf
  > https://www.minaandrawos.com/2015/12/06/concurrency-in-golang/

 ### goroutine
 A goroutine is a lightweight thread of execution in the Go programming language. It is similar to a thread in other programming languages, but it is managed by the Go runtime rather than the operating system. Goroutines allow concurrent execution of functions in a program, and they are designed to be efficient and scalable.

In Go, **a program starts with a single goroutine, which executes the main function**. Additional goroutines can be created using the **go keyword followed by a function call**. This starts a new goroutine that runs concurrently with the original goroutine.

Goroutines are very lightweight, and it's possible to create thousands or even millions of them in a single program without significant overhead. This makes it easy to write concurrent programs in Go that take advantage of multiple CPU cores and can perform many tasks simultaneously.

Because goroutines are managed by the Go runtime, they are automatically scheduled and can **communicate with each other using channels**. This makes it easy to write complex concurrent programs without worrying about low-level details such as locking and synchronization.

> https://www.golangprograms.com/goroutines.html

### Channels
In Go, a channel is a built-in data structure that is used for communication and synchronization between goroutines. Channels are a fundamental feature of the language that enable safe and efficient communication and synchronization between goroutines (concurrently executing functions), meaning channels are threadsafe.  Think of channel as a conduit between producer(s) and consumer(s).

A channel is essentially a conduit that allows data to be passed between goroutines. It has a s**pecific type**, which determines the type of data that can be sent through the channel. Channels are created using the **built-in make function** and can be buffered or unbuffered.

Unbuffered channels block the sending goroutine until there is a corresponding receiver ready to receive the value being sent. This means that data is guaranteed to be received in the order it was sent, and that synchronization is built into the channel.

Buffered channels, on the other hand, can hold a limited number of values (determined by the buffer size), and will only block the sending goroutine when the buffer is full. This can allow for some additional concurrency, but requires careful consideration to avoid deadlocks and other synchronization issues.

Channels are often used to coordinate the activities of different goroutines, allowing them to share data and work together without the need for explicit locking or synchronization. 

For example, to create a channel of type int, you can use the following code:
```go
ch := make(chan int)
```
Here's an example of creating a buffered channel of integers with a capacity of 3:
```go
ch := make(chan int, 3)
```

Once a channel is created, you can send values into the channel using the <- operator, and receive values from the channel using the same operator. For example:
```go
ch <- 42 // send 42 into the channel
x := <-ch // receive a value from the channel and assign it to x
```
Channels can also be used to signal between goroutines by sending and receiving values that don't carry any data. For example, a channel can be used to signal the termination of a goroutine:
```go
done := make(chan bool)

// goroutine
go func() {
    // do some work...
    done <- true // signal that the work is done
}()

// wait for the goroutine to finish
<-done
```

> more here: https://www.golangprograms.com/go-language/what-are-channels-in-golang.html


[The Go Programming Language Specification states](https://golang.org/ref/spec#Channel_types):
>“A single channel may be used in send statements, receive operations, and calls to the built-in functions cap and len by any number of goroutines without further synchronization.”
In other words, you can have multiple writers and multiple readers all using a single channel without a mutex or other lock. The channel itself manages the data and ensures the safety of concurrent access.

### Channels thread safe?
- Yes. You should be careful if you use pointer data structure with channels.
-
  





### Waitgroups
WaitGroups are another means of allowing additional threads to complete their process before the main thread runs to completion. They work by blocking the main thread until the goroutines associated with the WaitGroup have completed. To wait for multiple goroutines to finish, we can use a wait group.
Example
```go
// https://gobyexample.com/waitgroups
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup	// wg counter is 0

    for i := 1; i <= 5; i++ {
        wg.Add(1)	// increment WG counter

        i := i

        go func() {
            defer wg.Done() // once this line runs, the wg counter is reduced
            worker(i)
        }()
    }

    wg.Wait()
    // by now the wg counter is 0	
    fmt.Printf("All done! ")

}
```
#### Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/* output
hello	// note first hello
world
*/

package main
import "fmt"

func main() {
	fmt.Println("hello0")
	
	defer fmt.Println("world")

	fmt.Println("hello1")
	fmt.Println("hello2")
	fmt.Println("hello3")
}
/* output
hello0
hello1
hello2
hello3
world
*/

```

### Select statement

A select statement in Go serves a very similar purpose to the switch statement  except that the switch statement works with variables and comparison operations between those variables, select statements are optimized to work with channel operations. 

- **Blocking select**
```go
select {
    case channel operation:
        statements
    case channel operation:
        statements
}
```

We're going to start with the select keyword and then we're going to have a pair of matched curly braces. Notice we don't have any variable that we're going to be testing against in select statements, those are only used with switches. Then we're going to have a series of case statements, and those cases are going to consist of channel operations. So we might try and send a message into a channel or we might try and receive a message from a channel. These are often used where we might have multiple asynchronous processes, multiple Goroutines, that are doing work and the result might come back through several different channels. So for example, you might have one case that's listening for the response to a database query, but you might have another case that's listening for a message from a timer that's indicating if the operation is taking too long. So the first case would be what we would want to succeed, we want our database query to succeed, but the second case would allow us to ensure that the operation doesn't take too long and waste resources. Now in this form where we just have cases, this is what's called a blocking select, which means Go will not proceed, the select statement will stop until one of the cases can be operated on. 

- **a non‑blocking select**
```go
select {
    case channel operation:
        statements
    case channel operation:
        statements
    default:			//optional
        statementsa
}

If you want to create a non‑blocking select, then you can have a default clause in the select statement. When you have a default clause, Go is going to check each case to see if those channel operations are possible to be acted upon, if none of them can be, it's not going to block, it's just going to execute the default clauses statements. Now, the other thing to keep in mind with select statements is this. In a select statement, if more than one case can be acted upon, then one case is chosen randomly. Now this is different than switches. With a switch, the first case that is valid with a switch will be executed, so Go is going to look at the cases from the top down. With a select, if multiple cases can be acted upon, Go is actually going to choose randomly between those. Now that's a very intentional decision that was made by the Go team to ensure that we don't develop a dependency or reliance or an expectation upon how our cases are going to be executed. So, if multiple cases can be executed, one will be chosen at random. 

## Setup
- Install Go toolchain. I has - build, dependencies - like third party libraries, profile code, application tracng/debugging, test, documentation
  - go.dev and download go - go.dev/dl/.
  - Install go
  - go version
  - go
  - 
- Go editor
  - Visual Studio Code + extension + libraries
  - https://code.visualstudio.com/docs/?dv=darwinarm64
  - Install extension Go for Visual Studio code from teh Go team at Google. In extension search for Go
  - Menu - View -> Command Palette -> Go: Install/Update Tools -> selelct all
 
## Module vs Package
- A package is a directory of .go files, and it is the basic building block of a Go program. Packages help to organize code into reusable components. 
- On the other hand, a module is a collection of packages with built-in dependencies and versioning.  A module comes with two additional files **go.mod** and **go.sum**.
  - go.mod gives the name of the module, as well as its dependencies and minimum versions.
  - go.sum is a dependency lock file that is generated automatically.
- go mod options
  - go mod init: creates a whole new module in the current directory.
  - go mod tidy: fixes missing modules and removes others who aren’t in use.
  - go mod download: downloads modules to your device’s cache.
  - and more **go mod help**
  - https://www.workfall.com/learning/blog/how-to-use-go-modules-for-package-management/#:~:text=for%20Go%20beginners.-,Package%20vs%20Module,with%20two%20additional%20files%20go.

## Simple data types
```
Strings, Numbers, Booleans, Errors
```

### Strings
```
"this is a string: \n creates a new line" interpreted string
'this is also a string: \n does not create a new line' raw string
```

### error type
The error type in Go is implemented as the following interface:
```
type error interface {
    Error() string
}
```
So basically, an error is anything that implements the Error() method, which returns an error message as a string. It’s that simple!

#### Constructing Errors
Errors can be constructed on the fly using Go’s built-in errors or fmt packages. For example, the following function uses the errors package to return a new error with a static error message:
```
package main

import "errors"

func DoSomething() error {
    return errors.New("something didn't work")
}
```
- https://earthly.dev/blog/golang-errors/


### more on data type
- https://www.programiz.com/golang/data-types
- https://www.w3schools.com/go/go_data_types.php


## Variable Declaration
```
var myName string            //  declare variable
var myName string = "Mike"   // declare and initialize
var myName = "Mike"         // initialize with inferred type
```
- Following infers the type and creates variable of that type, note **:=", instead of "="
```
myName := "Mike"            // short declaration syntax
```

- Go allows multiple variables to be initialized at once!
```
a, b := 10, 5 // here a and bis inferred as  in datatype, created and initialied with 10 and 5 respectively
```

```
c := a + b  // 15 - addition, here c is inferred as  in datatype, created and initialied with 50
c = a - b   // 5 - subtraction, here we are assiging new values to existing  variable c
c = a * b   // 50 - multiplication
```

## Comparison - value and reference types

### value types
```
a, b : = 10, 5
с := a == b  // false - equality
c = a != b   // true - inequality
```

## Constants 
```
const a = 42  // constant (implicitly typed)
const b string = "Hello" // expliocity typed constant

const c = a  // one constant assigned to another

// group constant
const (
   d = true
   e = 3.14
)
```

## Pointers and Values
Pointers are primarily used to sahre memory.

```
a := 42
b := &a
fmt.print(*b) // prints 42, dereferencing pointer b
fmt.print(b) // prints memeory address of a
*b = 84 // changes the value held in address of a
fmt.print(a) // prints 84

c = new(int) // build-in "new" function creates  pointer to anonymous variable
```



## TYPE CONVERSION
Go doesn't support implicit conversions.
```
var i int = 32
var f float32
f= i  // error! - Go doesn't support implicit conversions
f = float32(i)  //type conversions allow explicit conversion
```

## Arrays
- int
![image](https://github.com/paramraghavan/java-to-go/assets/52529498/90aaee32-f84d-4a0f-9866-47abb2f700d5)

- strings
- 
- 
- compare arrays, **==**
```
arr = [3]string{"foo", "bar", "baz"}

//assigning arrays in go applies copy operation
arr2 := arr						// arrays are copied by value

fmt.PrintIn(arr2)					// /I {"foo" "bar" "baz”}

arr[0] = "quux"					// {"quux" "bar" "baz"}
fmt.Println(arr)					// {“foo' "bar' "baz”}
fmt.Print]n(arr2)

arr == arr2						// false - arrays are comparable
```

## Slice Data type
- Is a Reference data type, similar to pointers
- nil , means not pointing to anything
- ability add data to slices, so we can increase the sixe of the datastructure, unlike the array data type where the size is finite
- ![image](https://github.com/paramraghavan/java-to-go/assets/52529498/810ecee1-faad-4f02-8175-f354215fbf56)
- ```go
   go get golang.org/x/exp/slices // get the new slices library  dependency. exp indicates experimental library
  ```
- go.sum file will show the dependencies for this module

## Map data type
```go
var m map|string]int                    // declare a map
fmt.Print]n(m)                          // mapl] (nil)
m = map[string]int{"foo": 1, "bar": 2} // map literal
fmt.Println(m)                          // map [foo:1 bar :2]

fmt.PrintIn(m[" foo"])                  // lookup value in map
m ["bar" ] = 99                         // update value in map

delete (m, "foo")                       // remove entry from map

m[ "baz" ] = 418                        //add new key/value to map with
fmt.Println(m)                          // map[bar: 99 baz: 418]

fmt.PrintIn(m["foo" ])                  // 0 - foo has been removed, but queries always return results
v, ok := m[ "foo" ]                     // v,ok this syntax verifies presents
fmt.Println(v, ok)                       // 0, false

v, ok := m[ "baz" ] 
fmt.Println(v, ok)                       // 418, true
```

## Structs data type
A struct (short for "structure") is a collection of data fields with declared data types. Golang has the ability to declare and create own data types by combining one or more types, including both built-in and user-defined types. Each data field in a struct is declared with a known type, which could be a built-in type or another user-defined type.

Structs are the only way to create concrete user-defined types in Golang. Struct types are declared by composing a fixed set of unique fields. Structs can improve modularity and allow to create and pass complex data structures around the system. You can also consider Structs as a template for creating a data record, like an employee record or an e-commerce product.

The declaration starts with the keyword type, then a name for the new struct, and finally the keyword struct. Within the curly brackets, a series of data fields are specified with a name and a type.
```
type identifier struct{
  field1 data_type
  field2 data_type
  field3 data_type
}
```
- Creating Instances of Struct Types 
```go
package main
 
import "fmt"
 
type rectangle struct {
	length  int
	breadth int
	color   string
 
	geometry struct {
		area      int
		perimeter int
	}
}
 
func main() {
	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"
 
	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)
 
	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}
```
- Creating a Struct Instance Using a Struct Literal
```go
package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)

	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle{breadth: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)
}
```

- Struct Instantiation Using new Keyword
An instance of a struct can also be created with the new keyword. It is then possible to assign data values to the data fields using dot notation.
```go
package main
 
import "fmt"
 
type rectangle struct {
	length  int
	breadth int
	color   string
}
 
func main() {
	rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)
 
	var rect2 = new(rectangle) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}
```

- Struct Instantiation Using Pointer Address Operator
Creates an instance of rectangle struct by using a pointer address operator is denoted by & symbol.

```go
package main

import "fmt"

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {
	var rect1 = &rectangle{10, 20, "Green"} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}
```

> amd more here : https://www.golangprograms.com/go-language/struct.html


## looping
```
for { ... }					// infinite loop
for condition { ... }				// loop till condition
for initializer; test; post clause { ... }	// counter-based loop
```
> https://www.golangprograms.com/for-range-loops.html

## Documentation on standard libraries
- https://pkg.go.dev/std

## Go for Java developers
- https://github.com/fstab/go-programming-for-java-developers

## Go tutorials on github
  - https://github.com/topics/go-tutorial
  - https://www.workfall.com/learning/blog/how-to-use-go-modules-for-package-management/ **
  - https://www.golangprograms.com/
