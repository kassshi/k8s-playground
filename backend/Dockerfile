FROM golang:1.26.1 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
