# Sử dụng hình ảnh Alpine Go
FROM golang:1.22-alpine

# Thiết lập thư mục làm việc là /app
WORKDIR /app

# Sao chép file go.mod và go.sum trước để tối ưu hóa việc caching
COPY go.mod go.sum ./

# Tải và cài đặt các phụ thuộc
RUN go mod download

# Sao chép toàn bộ mã nguồn của ứng dụng Go vào thư mục /app trong container
COPY . .

# Biên dịch ứng dụng Go
RUN go build -o main .

# Expose port 8080 cho ứng dụng Go
EXPOSE 8080

# Chạy ứng dụng Go khi container được khởi chạy
CMD ["./main"]