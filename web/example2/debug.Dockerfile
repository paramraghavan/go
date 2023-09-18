# Compile stage
FROM golang:1.21.0-bullseye AS build-env

# Build Delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /dockerdev
WORKDIR /dockerdev

RUN go build -gcflags="all=-N -l" -o /server
# Final stage
FROM debian:buster
EXPOSE 8888 40000

WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]

# build with debug
# docker build --tag debug-image --file debug.Dockerfile .
# docker run --publish 8888:8888 --publish 40000:40000 --name debug-server debug-image
# remote debug on goland with 40000 and connect to todo service with localhost:8888
# curl -s localhost:8888
