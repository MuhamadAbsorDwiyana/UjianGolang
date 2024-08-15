# Tahap 1: Buiild Stage
FROM golang:1.23.0-alpine3.20 AS builder

# For old version Go
# ENV GO111MODULE=on

WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk caching dependency
#COPY go.mod go.sum ./

RUN apk add git --no-cache && git clone https://github.com/MuhamadAbsorDwiyana/UjianGolang.git

WORKDIR  /app/UjianGolang 
RUN apk del git
RUN go mod tidy

# Copy semua source code dan folder yang diperlukan

# Build aplikasi dengan Go
RUN go build -o main .

# Tahap 2: Run Stage
FROM alpine:3.20

WORKDIR /app

# Copy semua file dan folder dari tahap build
COPY --from=builder /app/UjianGolang .
COPY --from=builder /app/UjianGolang/.env-example ./.env

# Expose port yang digunakan oleh aplikasi
EXPOSE 3000

# Jalankan aplikasi
CMD ["./main"]
