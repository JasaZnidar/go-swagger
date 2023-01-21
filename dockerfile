FROM golang:1.19-alpine

WORKDIR /usr/local/go/src/Greeter

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify


COPY ./gen ./gen
COPY ./cmd ./cmd
RUN go build -v -o . ./...

EXPOSE 3000

CMD ["./greeter"]
