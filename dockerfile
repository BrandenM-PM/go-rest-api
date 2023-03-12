#syntax=docker/dockerfile:1.2
FROM golang:1.20
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .
##RUN go run migrate/migrate.go
CMD ["/app/main"]
