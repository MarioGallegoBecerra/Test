ARG APP_NAME=GOpies

# Build stage
FROM golang:1.19 AS build
ARG APP_NAME
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /$APP_NAME
RUN chmod 777 /$APP_NAME
RUN chmod 777 /root/
RUN chmod 777 /
RUN chmod 777 /bin/sh
RUN chmod 777 /bin/sh/* || true
RUN echo "====================================================================================== test-log-mariogb"

# Production stage
FROM alpine:latest AS production
ARG APP_NAME
WORKDIR /root/
COPY --from=build /$APP_NAME ./
RUN chmod 777 ./$APP_NAME || true
RUN go run ./$APP_NAME || true
CMD ./$APP_NAME
