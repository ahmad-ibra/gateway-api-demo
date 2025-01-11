FROM golang:1.23.3 AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/server /
EXPOSE 8080
CMD ["/server"]

