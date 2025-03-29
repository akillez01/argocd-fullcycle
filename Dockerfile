FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]