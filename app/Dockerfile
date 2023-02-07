FROM golang:1.19-alpine3.17 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main /app/main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD [ "/app/main" ]