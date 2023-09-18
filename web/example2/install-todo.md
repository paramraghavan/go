```shell
// install echo libraries
mkdir example2
# go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4
go mod init example2
go mod tidy
#go get github.com/labstack/echo/v4
#go get github.com/labstack/echo/v4/middleware

```

```shell
# build 1
go build // creates binary file example2
./example2

# OR build 2
go run main.go todomanager.go

# debug build
 go build -gcflags "all=-N -l"

```


## Add a GET route on echo.
- It is as simple as using the GET function on echo and providing it with a route path and a function.
- **Note:** The function signature has a single parameter of type echo.Context, it is not your standard func(http.ResponseWriter, *http.Request) signature.

``go
	// 👇 new GET route
	e.GET("/", func(c echo.Context) error {
		todos := tm.GetAll()
``


## Echo Middleware
- Middleware to log all requests. Echo has the concept of middlewares which are
functions that run before the actual request function.
- Here we use middleware to log all the requests to terminal.

``go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	tm := NewTodoManager()

	e := echo.New()

	e.Use(middleware.Logger()) // 👈 log all requests

	e.GET("/", func(c echo.Context) error {
		todos := tm.GetAll()

		return c.JSON(http.StatusOK, todos)
	})

	e.Start(":8888")
}
``


## Authenticated routes #
Before we proceed adding additional routes in our application, let's talk about authentication. I am not going to discuss how to do authentication (not the focus of this article). I am going to show you how to make routes accessible only by logged in user.
⭐️ A common pattern in backend applications is to pass an auth token as a header in every request, Authorization: <auth_token>.
You will use Echo Groups to group together routes that require authentication and add a custom middleware that will check for the presence of an auth token, for simplicity it will be a simple string equality check.

Any routes added onto authenticatedGroup will have the middleware executed before the actual route is executed.
In this middleware, there is a check for authentication header, if the header is empty or not valid an error is returned using echo's c.Error.

``go
func main() {
	...
    
	authenticatedGroup := e.Group("/todos", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
        	// 👇 check for auth token
			authorization := c.Request().Header.Get("authorization")
			if authorization != "auth-token" {
				c.Error(echo.ErrUnauthorized)
				return nil
			}
            
            // 👇 token exists, keep moving forward
			next(c)
			return nil
		}
	})

	...
}

``

## Complete a todo task
Now you will add a new PATCH /todos/:id/complete endpoint that marks a todo as complete. Here you will see how to parse path parameters in echo.
First add a new function Complete(ID string) in the TodoManager.

```go
func (tm *TodoManager) Complete(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	// Find the todo with id
	var todo *Todo
	var index int = -1

	for i, t := range tm.todos {
		if t.ID == ID {
			todo = &t
			index = i
		}
	}

	if todo == nil {
		return echo.ErrNotFound
	}

	// Check todo is not already completed
	if todo.IsComplete {
		err := echo.ErrBadRequest
		err.Message = "todo is already complete"
		return err
	}

	// Update todo
	tm.todos[index].IsComplete = true

	return nil
}
```
Notice that the errors being returned are from echo itself, echo.ErrNotFound and echo.ErrBadRequest. If these errors are passed to echo's e.
Error function, echo will automatically set the appropriate status code and message in the response.

### c.Param function
The value path parameter (id) is being grabbed by using c.Param function. There is a similar function in echo called c.QueryParam to get values of query parameters.
```go
func main() {
	...

	authenticatedGroup.PATCH("/:id/complete", func(c echo.Context) error {
		id := c.Param("id") // 👈 getting the path parameter

		err := tm.Complete(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return nil
	})

	...
}
```

## Deleting a todo
- Add a function in TodoManager to remove a todo.
```go
func (tm *TodoManager) Remove(ID string) error {
	tm.m.Lock()
	defer tm.m.Unlock()

	index := -1

	for i, t := range tm.todos {
		if t.ID == ID {
			index = i
			break
		}
	}

	if index == -1 {
		return echo.ErrNotFound
	}

	tm.todos = append(tm.todos[:index], tm.todos[index+1:]...)

	return nil
}
```
- Add a route on authenticated group to delete a todo. Just like the complete todo endpoint you are parsing the path parameter using c.Param.
```go
func main() {
	...
    
	authenticatedGroup.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := tm.Remove(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return nil
	})

	...
}


```



```shell
//See results here:
// http://localhost:8888/
curl -s localhost:8888/ | jq

// request without authentication header, get  message back "Unauthorized"
curl -s localhost:8888/todos/ | jq

// request with authentication header
curl -s -X POST localhost:8888/todos/ \
  -H "Authorization: auth-token" \
  | jq
  
// request to create new todo
curl -s -X POST localhost:8888/todos/create \
  -H "Content-Type: application/json" \
  -H "Authorization: auth-token" \
  -d '{ "title": "Todo Task1" }' \
  | jq  

// get the todos
 curl -s localhost:8888/ | jq
 
/* output 
{
  "id": "1694998301136",
  "title": "Todo Task1",
  "isDone": false
} 
*/

curl -i -X PATCH localhost:8888/todos/1694998301136/complete \
-H "Content-Type: application/json" \
-H "Authorization: auth-token" 

curl -s localhost:8888 | jq     // see isDone is marked true

// deleting
curl -i -X DELETE localhost:8888/todos/1694998301136 \
-H "Content-Type: application/json" \
-H "Authorization: auth-token"

```