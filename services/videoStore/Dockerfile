FROM golang:1.13.0-stretch

RUN mkdir -p /go/src/videoStore

WORKDIR /go/src/videoStore

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]