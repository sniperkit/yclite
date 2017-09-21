FROM golang:latest

RUN mkdir -p /go/src/github.com/shohi/yclite
ADD . /go/src/github.com/shohi/yclite/

WORKDIR /go/src/github.com/shohi/yclite
EXPOSE 8080
CMD ["go run main.go"]
