# Stage 1: Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

# Copy các file module trước
COPY ../go.mod ../go.sum ./

# Download module dependencies (cache hiệu quả hơn)
RUN go mod download

# Copy toàn bộ mã nguồn
COPY .. .

# Build ứng dụng
RUN go build -o main .

# Stage 2: Run stage
FROM alpine:3.18

WORKDIR /app

# Copy binary từ giai đoạn build
COPY --from=build /app/main .

# Sao chép thư mục internal/config vào container
COPY --from=build /app/internal/config /app/internal/config

EXPOSE 8080

CMD ["./main"]
