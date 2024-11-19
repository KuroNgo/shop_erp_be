# Dockerfile
# Use Build Multistage
# Stage 1: Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app
COPY ../go.mod ../go.sum ./
COPY ../vendor ./vendor
COPY .. .

RUN go build -o main .

# Stage 2: Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080
CMD ["./main"]

# docker build -t shoperp .
# docker run -v "$(pwd)/app.env:/app/app.env" -p 8080:8080 shoperp
