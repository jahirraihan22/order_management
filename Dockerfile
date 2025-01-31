FROM golang:alpine3.19 as builder
WORKDIR /app
COPY . ./
RUN go mod download && go build -o /order_management

FROM alpine:3.19
RUN apk update &&  apk --no-cache add ca-certificates tzdata

WORKDIR /
COPY --from=builder /order_management /order_management
EXPOSE 8080
ENTRYPOINT ["/order_management"]