FROM golang:alpine

WORKDIR /app

ARG DB_URL
ENV DB_URL=${DB_URL}

COPY go.mod go.sum /app/

RUN go mod download

RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . . 

RUN sqlc generate

EXPOSE 3001

CMD ["sh", "-c", "migrate -path=./sql/migrations -database ${DB_URL} up && air"]
