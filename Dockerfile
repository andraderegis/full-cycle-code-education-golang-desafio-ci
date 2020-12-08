FROM golang:1.14-alpine as builder
RUN apk add --no-cache upx
WORKDIR /go
COPY . ./src/golang-fullcycle-desafio-ci
RUN GOOS=linux go build -ldflags "-s -w" -o desafio_ci -i ./src/golang-fullcycle-desafio-ci/cmd/main.go
RUN upx --brute desafio_ci

# FROM busybox:latest
# WORKDIR /golang-fullcycle-desafio-ci
# COPY --from=builder /go/desafio_ci ./
ENTRYPOINT ["./desafio_ci"]

