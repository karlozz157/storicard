FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT=8080

EXPOSE $PORT

CMD ["./main"]
