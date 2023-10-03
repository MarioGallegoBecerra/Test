# Build stage
FROM alpine:latest AS production
ARG APP_NAME
WORKDIR /root/app/$APP_NAME
COPY --from=build /$APP_NAME ./
CMD ./$APP_NAME
