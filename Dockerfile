FROM golang:1.22 AS builder

WORKDIR /app

# Copy mod files first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Small runtime image
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]