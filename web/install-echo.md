``go
// install echo libraries
// go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4

``

```shell
// build your web project
mkdir web

#creates a whole new module in the current directory.
go mod init

#imports missing modules and removes ones not in use.
go mod tidy 
```

```shell
//See results here:
// http://localhost:1323/
```