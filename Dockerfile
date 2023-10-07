ARG APP_NAME=GOpies

# Build stage
FROM golang:1.19 AS build
ARG APP_NAME
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /$APP_NAME
CMD ./$APP_NAME
RUN echo "====================================================================================== test-log-mariogb"

# Production stage
FROM alpine:latest AS production
ARG APP_NAME
WORKDIR /workspace/
COPY --from=build /$APP_NAME ./
COPY --from=build /$APP_NAME ./workspace
CMD ./$APP_NAME
