FROM golang:1.20

WORKDIR /app
COPY . .

RUN go build -o app
RUN go run /app/cmd/createdb/

EXPOSE 8080

CMD ["./app"]