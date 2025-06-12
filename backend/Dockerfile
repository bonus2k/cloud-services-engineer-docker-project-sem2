FROM golang:1.17 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -a -o healthcheck ./cmd/healthcheck

FROM scratch
COPY --chown=1000:1000 --from=builder /app/app /
COPY --chown=1000:1000 --from=builder /app/healthcheck /
USER 1000:1000
EXPOSE 8081
ENTRYPOINT ["/app"]
