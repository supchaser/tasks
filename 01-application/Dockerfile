FROM golang:1.24-alpine AS builder
LABEL stage=builder
WORKDIR /src

RUN apk add --no-cache git

COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /echo-server ./cmd/main

FROM scratch AS runtime

LABEL maintainer="Valentin Stremin" \
      org.opencontainers.image.source="https://github.com/supchaser/tasks" \
      org.opencontainers.image.license="MIT"

COPY --from=builder /echo-server /echo-server

ENV AUTHOR="Unknown Author"
EXPOSE 8000

ENTRYPOINT ["/echo-server"]
