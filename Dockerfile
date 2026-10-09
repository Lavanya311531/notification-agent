FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o notification-agent ./cmd

EXPOSE 50051

CMD ["./notification-agent"]
