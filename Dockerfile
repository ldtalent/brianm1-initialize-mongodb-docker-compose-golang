FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

COPY init.js /docker-entrypoint-initdb.d/
RUN chmod +x /wait-for-it.sh

EXPOSE 8800
CMD ["/app/main"]