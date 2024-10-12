# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Thiết lập thư mục làm việc là /app
WORKDIR /app

# Sao chép file go.mod và go.sum để tối ưu caching
COPY go.mod go.sum ./

# Tải và cài đặt các phụ thuộc
RUN go mod download

# Sao chép mã nguồn của ứng dụng Go vào thư mục /app
COPY . .

# Biên dịch ứng dụng Go
RUN go build -o main .

# Stage 2: Run stage
FROM alpine:3.18

# Thiết lập thư mục làm việc là /app
WORKDIR /app

# Sao chép file thực thi từ stage build
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Chạy ứng dụng Go khi container được khởi chạy
CMD ["./main"]
