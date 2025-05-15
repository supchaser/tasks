FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go build -o /echo-server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /echo-server /app/echo-server
EXPOSE 8000
ENV AUTHOR="Valentin Stremin (supchaser)"
CMD ["/app/echo-server"]
