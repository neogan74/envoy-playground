FROM golang:1.22-alpine as build

WORKDIR /app 
COPY . .
RUN go mod download && go build -o service-uuid 


FROM alpine:latest
WORKDIR /app 
COPY --from=build /app/service-uuid 
EXPOSE 8080
ENTRYPOINT ["./service-uuid"]
