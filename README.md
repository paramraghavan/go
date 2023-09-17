# go
A quick **go** tour.

- [Super link for go](https://www.golangprograms.com/).
- [Best Golang Packages](https://www.golangprograms.com/go-programming-language-packages.html)
- [Go Sandbox](https://go.dev/play/)
>work in progress

## Web and go
**echo** is the golang equivalent for flask. [Example here](./web) 
> https://github.com/labstack/echo

## Is go compiled or interpreted. Is it VM based
- Go is a compiled language.
- This means we must run our source code files through a compiler, which reads source code and generates a binary, or executable file that is used to run the program
- It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.
- Go does not provide any VM, such as Java JVM. This language only compiles to binary like c++/c.
- Go has a Garbage Collector like java and no free/delete functions/operators like in C/C++.

## Does Go have a runtime?
*Go does have an extensive library, called the runtime*, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language. Although it is more central to the language, Go's runtime is analogous to libc , the C library.

## Concurrency and Go
Supports two styles
  - Communicating sequential processes (CSP) *use communication as synchronization primitive*
  - Shared memory multithreading using locks ???
#### CSP - Communicating Sequential Processes
Concept: CSP (Communicating Sequential Processes)
Imagine, if you will, that we take our program and we break it down into discrete execution units. So when viewing our program this way, we're going to have several workers or several tasks that our program is going to be able to execute. With Go, what we do is we envision each one of these workers as separate execution units that communicate with each other through something called a channel. So the first worker is going to do some work, and then it's going to pass the result of its work into a channel. The second worker will receive that work, and then it's going to do its own work upon that and send its result out to a channel as well. And then we might have another worker, and that might be the end of the execution chain. That worker takes the work from the second worker, and it finishes the job.

The idea of breaking a program up into a dicrete unit of work, each unit being a worker/goroutine is called Communicating Sequential Processes, or CSP. The idea here is that our processes, workers, in this case, are communicating via these channels sequentially. Each one acts independently, but they communicate with each other with the results of their work, and then they potentially take inputs in from other workers as well. 

The beautiful thing about this model is the workers can work independently, which means we can actually have each one of these workers act concurrently, as long as we have some sort of a synchronization mechanism to make sure that they are able to communicate when they need to.

We could have three workers that are generating input values, and then we can have a single worker that's taking in the results of the work of those workers. As long as we have a single channel, it doesn't matter where that message comes from that's going into that channel, it's going to be passed on to that worker that's taking in the results, and it's able to work on that. This is what's called a fan‑in pattern, where we've got multiple input sources that are generating results.

Since all of these can work concurrently, this allows us to have many input workers, and as a result, we can generate input values much faster. Similarly, we could flip the model around. We could have a single worker that's generating input values, and we could have multiple workers that are receiving that. Once again, as long as we have a single channel that's communicating those messages, the worker on the producer side can generate its message and pass that into the channel, and then it doesn't matter, or it doesn't care which one of the workers on the consumer side gets that message. This is ideal if our messages are slow to process, so maybe our producer can generate messages much, much faster than the workers can consume. Since all of these are working in their own concurrent tasks, then we can balance that load out. This is called a fan‑out pattern, where you've got a few input sources and many output sources.
And the whole idea here is to balance the level of concurrency in our program, so our programs are running efficiently as possible. As long as we're using CSP, Communicating Sequential Processes, these concurrency patterns become very, very easy to implement. 

##### Share Memory By Communicating
Traditional threading models (commonly used when writing Java, C++, and Python programs, for example) require the programmer to communicate between threads using shared memory. Typically, shared data structures are protected by locks, and threads will contend over those locks to access the data. In some cases, this is made easier by the use of thread-safe data structures such as Python’s Queue.

Go’s concurrency primitives - goroutines and channels - provide an elegant and distinct means of structuring concurrent software. Instead of explicitly using locks to mediate access to shared data, Go encourages the use of channels to pass references to data between goroutines. This approach ensures that only one goroutine has access to the data at a given time. 

*Do not communicate by sharing memory; instead, share memory by communicating.*
>> https://go.dev/blog/codelab-share


#### Shared memory multithreading ??????
Although we can do everything with CSP, sometimes less convenient than shared memory

  > https://www.cs.princeton.edu/courses/archive/fall16/cos418/docs/P1-concurrency.pdf
  > https://www.minaandrawos.com/2015/12/06/concurrency-in-golang/

 ### goroutine
 A goroutine is a lightweight thread of execution in the Go programming language. It is similar to a thread in other programming languages, but it is managed by the Go runtime rather than the operating system. Goroutines allow concurrent execution of functions in a program, and they are designed to be efficient and scalable.

In Go, **a program starts with a single goroutine, which executes the main function**. Additional goroutines can be created using the **go keyword followed by a function call**. This starts a new goroutine that runs concurrently with the original goroutine.

Goroutines are very lightweight, and it's possible to create thousands or even millions of them in a single program without significant overhead. This makes it easy to write concurrent programs in Go that take advantage of multiple CPU cores and can perform many tasks simultaneously.

Because goroutines are managed by the Go runtime, they are automatically scheduled and can **communicate with each other using channels**. This makes it easy to write complex concurrent programs without worrying about low-level details such as locking and synchronization.

### Goroutine vs Thread:
|Goroutine|	Thread|
|---------|--------|
|Goroutines are managed by the go runtime.|	Operating system threads are managed by kernel.|
|Goroutine are not hardware dependent.|	Threads are hardware dependent.|
|Goroutines have easy communication medium known as channel.|	Thread does not have easy communication medium.|
|Due to the presence of channel one goroutine can communicate with other goroutine with low latency.|	Due to lack of easy communication medium inter-threads communicate takes place with high latency.|
|Goroutines are cheaper than threads.|The cost of threads are higher than goroutine.|
|faster startup time than threads.| slow startup time than goroutines.|
|growable stack| fixed stack|

> https://www.golangprograms.com/goroutines.html

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


## Package Main - how does it work
When you build reusable pieces of code, you will develop a package as a shared library. But when you develop executable programs, you will use the package “main” for making the package as an executable program. _The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. The main function in the package “main” will be the entry point of our executable program_. When you build shared libraries, you will not have any main package and main function in the package.
Example
```go
package main
 
import (
"fmt"
)
func main(){
 fmt.Println("Hello, Gopher!")
}
```

### Import Packages
The keyword “import” is used for importing a package into other packages. In the Code Listing -1, we have imported the package “fmt” into the sample program for using the function Println. The package “fmt” comes from the Go standard library. When you import packages, the Go compiler will look on the locations specified by the environment variable GOROOT and GOPATH. Packages from the standard library are available in the GOROOT location. The packages that are created by yourself, and third-party packages which you have imported, are available in the GOPATH location.

### Install Third-Party Packages
We can download and install third-party Go packages by using “Go get” command. The Go get command will fetch the packages from the source repository and put the packages on the GOPATH location.

The following command in the terminal will install “mgo”,  a third-party Go driver package for MongoDB, into your GOPATH, which can be used across the projects put on the GOPATH directory:
```go
go get gopkg.in/mgo.v2
```

After installing the mgo, put the import statement in your programs for reusing the code, as shown below:
```go
import (        
        "gopkg.in/mgo.v2" 
        "gopkg.in/mgo.v2/bson"       
)
```
The MongoDB driver, mgo,  provides two packages that we have imported in the above import statement.

> more here... https://thenewstack.io/understanding-golang-packages/



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
- Yes. 

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
            defer wg.Done() // once this line runs, the wg counter is decremented by 1
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
```
If you want to create a non‑blocking select,_ then you can have a default clause_ in the select statement. When you have a default clause, Go is going to check each case to see if those channel operations are possible to be acted upon, if none of them can be, it's not going to block, it's just going to execute the default clauses statements. Now, the other thing to keep in mind with select statements is this. In a select statement, if more than one case can be acted upon, then one case is chosen randomly. Now this is different than switches. With a switch, the first case that is valid with a switch will be executed, so Go is going to look at the cases from the top down. With a select, if multiple cases can be acted upon, Go is actually going to choose randomly between those. Now that's a very intentional decision that was made by the Go team to ensure that we don't develop a dependency or reliance or an expectation upon how our cases are going to be executed. So, if multiple cases can be executed, one will be chosen at random.


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
| more in details... https://earthly.dev/blog/golang-errors/


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
a, b := 10, 5 // here a and b is inferred as  in datatype, created and initialied with 10 and 5 respectively
```

```
c := a + b  // 15 - addition, here c is inferred as  in datatype, created and initialied with 50
c = a - b   // 5 - subtraction, here we are assiging new values to existing  variable c
c = a * b   // 50 - multiplication
```

## sprintf
```go
import "fmt"

dbName := "master"
tableName := "employee"
str := fmt.Sprintf("%s.%s", dbName, fromTable),

fmt.Println(str)

/*output
master.employee
*/
 
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
const b string = "Hello" // explicity typed constant

const c = a  // one constant assigned to another

// group constant
const (
   d = true
   e = 3.14
)
```

## Pointers and Values
Pointers are primarily used to share memory.

```
a := 42
b := &a
fmt.print(*b) // prints 42, dereferencing pointer b
fmt.print(b) // prints memeory address of a
*b = 84 // changes the value held in address of a
fmt.print(a) // prints 84

c := new(int) // build-in "new" function creates  pointer to anonymous variable
var d = new(int) // build-in "new" function creates  pointer to anonymous variable
*c = 10
fmt.Println(c)	// output 10
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
```go
var arr [3]int		// array of 3 ints
fmt.Println(arr)	// [0 0 0]
arr = [3]int{1, 2, 3}	// array literal

fmt.Println(arr[1])	// 2
arr[1] = 99		// update value
fmt.Println(arr)	// [1 99 3]

fmt.Println(len(arr))	// 3

// for loop and range over array	
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
	sum += num
}
fmt.Println("sum:", sum)	// sum: 9

for i, num := range nums {
   if num == 3 {
     fmt.Println("index:", i)	// index: 1
   }
}



```  

- strings

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
- ability add data to slices, so we can increase the size of the datastructure, unlike the array data type where the size is finite
-
```go
var s []int				// slices of ints
fmt.Println(s)				// (nil) 
s = []int{1, 2, 3)			// slice litera 

fmt.Println(s[1])			// 2 
s[1] = 99				// update value 
fmt.Println(s)				// [1 99 3]
	
s = append(s, 5, 10, 15)		// add elements to the slice
fmt.Println(s)				// [1 99 3 5 10 15]

s = slices.Delete(s, 1, 3) 		// remove indices 1, 2 from slice (golang.org/x/exp/slices)
fmt.Println s)				// [1 5 10 15]
```

- **make** to allocate memeory, slice and append to slice.
- If there's sufficient capacity in the underlying slice, the element is placed after the last element and the length get incremented. However, if there is not sufficient capacity, a new slice is created, all of the existing elements are copied over, the new element is added onto the end, and the new slice is returned. Notcie after appending the - _Length is 9 Capacity is 12_
```go
a := make([]int, 2, 5)
a[0] = 10
a[1] = 20
fmt.Println("Slice A:", a)
fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

a = append(a, 30, 40, 50, 60, 70, 80, 90)
fmt.Println("Slice A after appending data:", a)
fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

/* output
Slice A: [10 20]
Length is 2 Capacity is 5
Slice A after appending data: [10 20 30 40 50 60 70 80 90]
Length is 9 Capacity is 12
*/

```  

-
```go
   go get golang.org/x/exp/slices // get the new slices library  dependency. exp indicates experimental library
```
- go.sum file will show the dependencies for this module

## Map data type
```go
var m map[string]int                    // declare a map
fmt.Println(m)                          // map (nil)
m = map[string]int{"foo": 1, "bar": 2} // map literal
fmt.Println(m)                          // map [foo:1 bar :2]

fmt.Println(m[" foo"])                  // lookup value in map
m ["bar" ] = 99                         // update value in map

delete (m, "foo")                       // remove entry from map

m[ "baz" ] = 418                        //add new key/value to map with
fmt.Println(m)                          // map[bar: 99 baz: 418]

fmt.Println(m["foo" ])                  // 0 - foo has been removed, but queries always return results
v, ok := m[ "foo" ]                     // v,ok this syntax verifies presents
fmt.Println(v, ok)                       // 0, false

v, ok := m[ "baz" ] 
fmt.Println(v, ok)                       // 418, true


kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
   fmt.Printf("%s -> %s\n", k, v)
}
/* output
a -> apple
b -> banana
*/
	
for k := range kvs {
  fmt.Println("key:", k)
}
/* output
key: a
key: b
*/

for i, c := range "go" { // range on strings iterates over Unicode code points.
  fmt.Println(i, c)
}
/* output 
0 103
1 111
*/
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


## Interface
An Interface is an abstract type. Interface describes all the methods and provides the signatures for each method. An interfaces act as a blueprint for method sets, they must be implemented before being used. 
_Type that satisfies an interface is said to implement it._

### Define Type that Satisfies an Interface
- Defines an interface type named Employee with two methods.
- Then it defines a type named Emp that satisfies Employee.
- _We define all the methods on Emp that it needs to satisfy Employee_
```go
package main

import "fmt"

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

//
// Emp type defines all the methods that it needs to satisfy Employee
// see below
//
// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
```

### Define Type that Satisfies Multiple Interfaces
Interfaces allows any user-defined type to satisfy multiple interface types at once.
```go
package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
```
> and more ..  https://www.golangprograms.com/go-language/interface.html

### Empty Interface Type
The type interface{} is known as the empty interface, and it is used to accept values of any type. The empty interface doesn't have any methods that are required to satisfy it, and so every type satisfies it.

In the following example, the manyType variable is declared to be of the type interface{} and it is able to be assigned values of different types. The printType() function takes a parameter of the type interface{},
hence this function can take the values of any valid type.

```go
package main

import "fmt"

func printType(i interface{}) {
	fmt.Println(i)
}

func main() {
	var manyType interface{}
	manyType = 100
	fmt.Println(manyType)

	manyType = 200.50
	fmt.Println(manyType)

	manyType = "Germany"
	fmt.Println(manyType)

	printType("Go programming language")
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	printType(countries)

	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	printType(employee)

	country := [3]string{"Japan", "Australia", "Germany"}
	printType(country)
}
```


## looping
```
for { ... }					// infinite loop
for condition { ... }				// loop till condition
for initializer; test; post clause { ... }	// counter-based loop
```
> https://www.golangprograms.com/for-range-loops.html

## Functions
### Naming Conventions for Golang Functions
* A name must begin with a letter, and can have any number of additional letters and numbers.
* A function name cannot start with a number.
* A function name cannot contain spaces.
* __If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.__
* If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
* function names are case-sensitive (car, Car and CAR are three different variables).

### Golang Passing Address to a Function
Passing the address of variable to the function and the value of a variables modified using dereferencing inside body of function.  
```go
package main

import "fmt"

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)
}
/* Output
Before: John 20
After : John Doe 25
*/
```
### Anonymous Functions in Golang
An anonymous function is a function that was declared without any named identifier to refer to it. Anonymous functions can accept inputs and return outputs, just as standard functions do.
**Assigning function to the variable.**
```go
package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))
}
```

### Closures Functions in Golang
Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.

**Anonymous function accessing the variable defined outside body.**
```go
package main

import "fmt"

func main() {
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
```


### Go and Varargs -  variadic function
A variadic function is a function that accepts a variable number of arguments. In Golang, it is possible to pass a varying number of arguments of the same type as referenced in the function signature. To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type. This type of function is useful when you don't know the number of arguments you are passing to the function, the best example is built-in Println function of the fmt package which is a variadic function.

```go


package main

import "fmt"

func main() {
	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 20))
}

func calculation(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

/* output
600
400
*/
```
### Deferred Functions Calls
Go has a special statement called defer that schedules a function call to be run after the function completes. Consider the following example:
```go


package main
import "fmt"
func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func main() {
	defer second()
	first()
}
// This program prints First followed by Second.
```

> and more... https://www.golangprograms.com/go-language/functions.html




## more on data type
- https://www.programiz.com/golang/data-types
- https://www.w3schools.com/go/go_data_types.php


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
 
## Memory Statistics in Go and Golang
The Go standard library has a host of functions to peek at memory statistics runtime. We can use it to investigate what is going on behind the scene as the garbage collection works in the background. The runtime package offers some key struct types that can be used to gather memory info at runtime. One of them is called MemStats. This can be used to get feedback on the statistics of the memory allocator. Some of the key fields of MemStats type and what they refer to are as follows. Note that all of these are declared as 64-bit unsigned int:

```go
type MemStats struct {
	Alloc 		uint64
	TotalAlloc 	uint64
	Mallocs 		uint64
	Frees 		uint64
	Sys			uint64
	...
}

```
* Alloc: It represents bytes of allocated heap objects. The bytes increase as more objects are created and decrease as they are deallocated.
* TotalAlloc: It keeps track of the total number of bytes allocated in the heap objects; however, the number of bytes does not get adjusted as memory gets deallocated through the garbage collector.
* Sys: It represents total bytes of memory obtained from the Operating System.
* Mallocs and Frees: The malloc represents the total count of heap objects allocated and Frees represents the total number of heap objects deallocated. Therefore, the count of live objects is always Mallocs – Frees.
There is also HeapAlloc, HeapSys, HeapIdle, HeapInuse, which represent bytes of allocated heap objects, bytes of heap memory obtained from OS, bytes of unused heap spans, and bytes of used heap span, respectively. Similarly, there are StackAlloc, StackSys, StackIdle, and StackInuse representing stack information.

### Example of Garbage Collection in Go and Golang
Let us write some simple Go code to get the memory statistics of a running program. You can extend it to a bigger program as well. The point here is to illustrate how to extract memory information. Getting a memory snapshot after a certain interval – and then comparing and investigating the result – will reveal how garbage collection works behind the scenes.
```go
File: memprog1.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var ms runtime.MemStats
	printMemStat(ms)

	//----------------------------------
	// you can write any code here
	//----------------------------------
	intArr := make([]int, 900000)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}
	//------------------------------------
	time.Sleep(5 * time.Second)

	printMemStat(ms)

}

func printMemStat(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)
	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total bytes of Heap object: ", ms.TotalAlloc)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
}
```
Output
```go
-------------------------------------------------------
Memory Statistics Reporting time:  2022-04-14 17:43:11.048224903 +0530 IST m=+0.000264317
-------------------------------------------------------
Bytes of allocated heap objects:  89432
Total bytes of Heap object:  89432
Bytes of memory obtained from OS:  8211472
Count of heap objects:  180
Count of heap objects freed:  3
Count of live heap objects 177
NumGC is the number of completed GC cycles:  0
-------------------------------------------------------
-------------------------------------------------------
Memory Statistics Reporting time:  2022-04-14 17:43:16.072656121 +0530 IST m=+5.024695581
-------------------------------------------------------
Bytes of allocated heap objects:  7285832
Total bytes of Heap object:  7301992
Bytes of memory obtained from OS:  17189648
Count of heap objects:  227
Count of heap objects freed:  47
Count of live heap objects 180
NumGC is the number of completed GC cycles:  1
-------------------------------------------------------
```

There is a way to get even more detailed info about the Go garbage collector using the following command while running the program above:
```go
GODEBUG=gctrace=1 go run memprog1.go
```
> https://www.developer.com/languages/garbage-collection-go/
> https://github.com/golang/go/blob/master/src/runtime/mgc.go

### The implementation of Go’s garbage collector
Go’s garbage collector is a non-generational concurrent, tri-color mark and sweep garbage collector. Let’s break these terms down.
The generational hypothesis assumes that short lived objects, like temporary variables, are reclaimed most often. Thus, a generational garbage collector focuses on recently allocated objects. However, as mentioned before, compiler optimisations allow the Go compiler to allocate objects with a known lifetime to the stack. This means fewer objects will be on the heap, so fewer objects will be garbage collected. This means that a generational garbage collector is not necessary in Go. So, Go uses a non-generational garbage collector. Concurrent means that the collector runs at the same time as mutator threads. Therefore, Go uses a non-generational concurrent garbage collector. Mark and sweep is the type of garbage collector and tri-color is the algorithm used to implement this
A mark and sweep garbage collector has two phases, unsurprisingly named mark and sweep. In the mark phase the collector traverses the heap and marks objects that are no longer needed. The follow-up sweep phase removes these objects. Mark and sweep is an indirect algorithm, as it marks live objects, and removes everything else.
> https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8

## Documentation on standard libraries
- https://pkg.go.dev/std

## Go for Java developers
- https://github.com/fstab/go-programming-for-java-developers

## Go tutorials on github
  - https://github.com/topics/go-tutorial
  - https://www.workfall.com/learning/blog/how-to-use-go-modules-for-package-management/ **
  - https://www.golangprograms.com/
  - https://github.com/forhadulislam/learn-go-with-examples
  - https://www.twilio.com/blog/scaling-your-go-application-with-kubernetes
  - https://go.googlesource.com/example
