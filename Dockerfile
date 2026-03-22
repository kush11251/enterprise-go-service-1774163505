FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go test ./... && go build -o /usr/local/bin/server ./cmd/server

FROM gcr.io/distroless/base
COPY --from=builder /usr/local/bin/server /usr/local/bin/server
ENTRYPOINT ["/usr/local/bin/server"]