```shell
// install echo libraries
// go get github.com/labstack/echo/{version}
go get github.com/labstack/echo/v4

```

```shell
// build your web project
mkdir example1
cd example1

#creates a whole new module in the current directory.
go mod init example1

#imports missing modules and removes ones not in use.
go mod tidy

#run
go run .
```

```shell
//See results here:
// http://localhost:1323/
curl -s localhost:1323
```
