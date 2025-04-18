FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
RUN go mod download && go build -o service-health

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/service-health .
EXPOSE 8081
ENTRYPOINT ["./service-health"]
