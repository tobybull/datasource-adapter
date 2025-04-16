# Stage 1: Builder
FROM golang:1.24.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's caching
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -ldflags '-s -w' -o /out/datasource-adapter ./cmd/main.go

# Stage 2: Final (Minimal) Image
FROM scratch
WORKDIR /app
COPY --from=builder /out/datasource-adapter /app/datasource-adapter
EXPOSE 8080
CMD ["/app/datasource-adapter"]
