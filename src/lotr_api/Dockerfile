FROM golang:1.24.6

WORKDIR /app/src/lotr_api

COPY . .

RUN go build main.go

CMD ["./main"]