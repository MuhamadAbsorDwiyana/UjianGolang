# Tahap 1: Build Stage
FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk caching dependency
COPY go.mod go.sum ./

RUN go mod tidy

# Copy semua source code dan folder yang diperlukan ke dalam container
COPY . .

# Build aplikasi dengan Go
RUN go build -o main .

# Tahap 2: Run Stage
FROM alpine:latest

WORKDIR /app

# Copy semua file dan folder dari tahap build
COPY --from=builder /app .

# Expose port yang digunakan oleh aplikasi
EXPOSE 3000

# Jalankan aplikasi
CMD ["./main"]
