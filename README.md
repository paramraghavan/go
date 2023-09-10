# java-go
## Go for Java developers
- https://github.com/fstab/go-programming-for-java-developers

# documentation on standard libraries
- https://pkg.go.dev/std

## Is go compiled or interpreted. Is it VM based
- Go is a compiled language.
- This means we must run our source code files through a compiler, which reads source code and generates a binary, or executable, file that is used to run the program
- It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

## Concurrency and Go
Supports two styles
  - Communicating sequential processes (CSP) use communication as synchronization primitive
  - Shared memory multithreading uses locks

## Setup
- Intall Go toolchain. I has - build, dependencies - like third party libraries, profile code, application tracng/debugging, test, documentation
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




## Go tutorials on github
  - https://github.com/topics/go-tutorial
  - https://www.workfall.com/learning/blog/how-to-use-go-modules-for-package-management/ **
