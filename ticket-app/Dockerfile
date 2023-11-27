FROM golang:bullseye
EXPOSE 6000

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["go", "run", "main.go"]