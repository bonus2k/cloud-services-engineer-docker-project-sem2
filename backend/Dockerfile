FROM golang:1.17 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

FROM scratch
COPY --from=builder /app/app /
EXPOSE 8081
ENTRYPOINT ["/app"]
