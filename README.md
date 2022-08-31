## OpenGin

[![version](https://img.shields.io/badge/version-1.0.0-blue)]()

<p>OpenGin is a WebApi template based on <strong>Gin</strong> & <strong>OpenApi3.0</strong>.</p>

## Contents

- [OpenGin](#opengin)
  - [Contents](#contents)
  - [Required](#required)
  - [About Air](#about-air)
  - [Docker(Windows)](#dockerwindows)

## Required

|Category|Package
|:---:|:---:|
|Gin|github.com/gin-gonic/gin
|Cors|github.com/gin-contrib/cors
|GORM|gorm.io/gorm
|Air|github.com/cosmtrek/air@latest
|Golang-jwt|github.com/golang-jwt/jwt/v4
|WebSocket|github.com/gorilla/websocket
|RabbitMQ|github.com/rabbitmq/amqp091-go
|Redis|github.com/go-redis/redis/v9
|SwaggerUI|https://github.com/swagger-api/swagger-ui/tree/master/dist

## About Air

Install
```go
go install github.com/cosmtrek/air@latest
```

How to use
```bash
air
```

## Docker(Windows)
```
docker build -f .\docker\Dockerfile -t roll:1 .
```

