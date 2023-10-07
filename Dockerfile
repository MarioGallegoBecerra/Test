ARG APP_NAME=Gopies.go

# Build stage
FROM golang:1.19 AS build
ARG APP_NAME
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /$APP_NAME
RUN chmod +x /$APP_NAME

# Production stage
FROM alpine:latest AS production
ARG APP_NAME
WORKDIR /root/
COPY --from=build /$APP_NAME ./
RUN chmod +x ./$APP_NAME
CMD ./$APP_NAME
