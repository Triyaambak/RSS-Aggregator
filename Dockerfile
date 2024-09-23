FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . . 

EXPOSE 3001

CMD ["air"]
