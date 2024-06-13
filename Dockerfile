FROM golang:1.21.11-alpine AS builder

USER root

RUN apk --no-cache --no-progress add build-base git bash

WORKDIR /src

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/go-web-template .

RUN chmod 755 .go-web-template

EXPOSE 3000

CMD ["./go-web-template"]

