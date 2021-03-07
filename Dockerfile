FROM golang:1.16.0

WORKDIR /go/src/github.com/viniciusramosdefaria/zerohashchallenge/

COPY . .

RUN GOCACHE='/tmp/.cache' CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/pictionary -v ./cmd/pictionary/main.go

FROM alpine:3.10.2
RUN apk update && apk add ca-certificates

COPY --from=0 /go/src/github.com/viniciusramosdefaria/zerohashchallenge/bin/pictionary /server/
COPY --from=0 /go/src/github.com/viniciusramosdefaria/zerohashchallenge/static /server/static

WORKDIR /server

RUN chmod +x /server/pictionary

ENTRYPOINT ["/server/pictionary"]
