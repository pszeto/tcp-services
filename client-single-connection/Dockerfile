FROM golang:1.23-bookworm AS builder
WORKDIR /
COPY go.* ./
RUN go mod download
COPY main.go ./
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o tcp-client-single

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /tcp-client-single /tcp-client-single
EXPOSE 8000
ENTRYPOINT ["/tcp-client-single"]