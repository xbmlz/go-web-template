FROM golang:1.21.11-alpine AS builder

RUN apk --no-cache --no-progress add \
  bash \
  ca-certificates \
  make

WORKDIR /src

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/go-web-template .

EXPOSE 3000

CMD ["./go-web-template"]

