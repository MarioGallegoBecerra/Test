# Production stage
FROM alpine:latest AS production
FROM golang:1.19 AS build
ARG APP_NAME
WORKDIR /root/app/$APP_NAME
COPY --from=build /$APP_NAME ./
CMD ./$APP_NAME
