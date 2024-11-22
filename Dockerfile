## Stage 1: Build stage
#FROM golang:1.22-alpine AS build
#
#WORKDIR /app
#COPY ../go.mod ../go.sum ./
#COPY ../vendor ./vendor
#COPY .. .
#
#RUN go build -o main .
#
## Stage 2: Run stage
#FROM alpine:3.18
#
#WORKDIR /app
#
## Copy binary từ giai đoạn build
#COPY --from=build /app/main .
#
## Sao chép thư mục internal/config vào container
#COPY --from=build /app/internal/config /app/internal/config
#
#EXPOSE 8080
#CMD ["./main"]

# Stage 1: Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod tidy

# Nếu cần thiết, bạn có thể sao chép chỉ các dependencies cần thiết
COPY ../ .

RUN go build -o main .

# Stage 2: Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080
CMD ["./main"]
