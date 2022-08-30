FROM golang:1.18 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY /src/ .
RUN go build -o app .

FROM scratch
LABEL authors="Shone"
ENV env=release
COPY --from=builder /build/app /
EXPOSE 80
ENTRYPOINT ["app"]