FROM golang:1.18 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY /src/ .
RUN go build -o start_server .

FROM scratch
LABEL authors="Shone"
ENV env=release
WORKDIR /home/app
COPY --from=builder /build .
EXPOSE 80
CMD ["./start_server"]