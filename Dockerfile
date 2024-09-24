FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . . 

RUN sqlc generate

EXPOSE 3001

CMD ["air"]
