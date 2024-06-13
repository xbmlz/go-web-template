FROM m.daocloud.io/docker.io/library/golang:1.21.11-alpine AS builder

WORKDIR /src

COPY . .

RUN make build

FROM m.daocloud.io/docker.io/library/alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/go-web-template .

EXPOSE 3000

CMD ["./go-web-template"]

