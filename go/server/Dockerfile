FROM golang:1.23-bookworm AS builder
WORKDIR /
COPY go.* ./
RUN go mod download
COPY main.go ./
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tcp-server

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /tcp-server /tcp-server
EXPOSE 8000
ENTRYPOINT ["/tcp-server"]