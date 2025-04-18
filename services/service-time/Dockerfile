FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
RUN go mod download && go build -o service-time

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/service-time . 
EXPOSE 8082
ENTRYPOINT ["./service-time"]

