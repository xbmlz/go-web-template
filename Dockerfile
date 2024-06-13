FROM golang:1.21.11-alpine AS builder

RUN apk --no-cache --no-progress add build-base git bash

WORKDIR /src

COPY . .

RUN make build

RUN chmod 755 ./bin/go-web-template

FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/go-web-template .

EXPOSE 3000

CMD ["./go-web-template"]

