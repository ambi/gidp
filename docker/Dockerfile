ARG GO_VERSION=1.14

# 1 prebuild-app
FROM golang:${GO_VERSION} AS prebuild-app

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

# 2a.1 build-echoapp
FROM prebuild-echoapp AS build-echoapp

COPY . .

RUN go build -o /go/bin/echoapp cmd/echoapp/main.go

# 2a.2 exec-echoapp (binary)
FROM debian:10-slim AS exec-echoapp

WORKDIR /app

COPY --from=build-echoapp /go/bin/echoapp .

ENTRYPOINT ["./echoapp"]

# 2b.1 build-grpcapp
FROM prebuild-grpcapp AS build-grpcapp

COPY . .

RUN go build -o /go/bin/grpcapp cmd/grpcapp/main.go

# 2b.2 exec-grpcapp (binary)
FROM debian:10-slim AS exec-grpcapp

WORKDIR /app

COPY --from=build-grpcapp /go/bin/grpcapp .

ENTRYPOINT ["./grpcapp"]
