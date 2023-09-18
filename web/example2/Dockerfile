# syntax=docker/dockerfile:1

FROM golang:1.21-alpine
LABEL maintainer="Me"
RUN mkdir /example2
COPY . /example2
WORKDIR /example2/
RUN go get github.com/labstack/echo/v4

EXPOSE 8888
CMD go run main.go todomanager.go
# CMD ["go", "run", "/var/server/main.go", "/var/server/todomanager.go"]

# build with debug
# docker build -t todo .
# docker run -it -p 3000:8888 todo
# curl -s localhost:3000
