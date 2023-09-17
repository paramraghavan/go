# Steps to build in go

## Create a directory

```go
mkdir sample
cd sample
```

## Create a new module

```go  
go mod init github.com/username/sample                    // or --> go mod init sample
# Here, the module name is: github.com/username/sample.
# You're free to choose any module name.
# It doesn't matter as long as it's unique.
# It's better to be a URL: so it can be go-gettable.

# imports missing modules and removes ones not in use.
go mod tidy 
```
## Put all your files in that directory.
## Finally, run:

```go
go run .
```

- Alternatively, you can create an executable program by building it:

```go
go build .

# then:
./sample     # if you're on xnix

# or, just:
sample       # if you're on Windows
```
>more details... https://go.dev/blog/using-go-modules
