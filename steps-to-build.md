# Steps to build in go

## Create a directory

```go
mkdir simple
cd simple
```

## Create a new module

```go
go mod init github.com/username/simple
# Here, the module name is: github.com/username/simple.
# You're free to choose any module name.
# It doesn't matter as long as it's unique.
# It's better to be a URL: so it can be go-gettable.
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
./simple     # if you're on xnix

# or, just:
simple       # if you're on Windows
```
>more details... https://go.dev/blog/using-go-modules
