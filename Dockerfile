FROM m.daocloud.io/docker.io/library/golang:1.21.11-alpine AS builder


# Move to working directory (/build).
WORKDIR /src

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

RUN make build

FROM m.daocloud.io/docker.io/library/alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/go-web-template .

EXPOSE 8080

CMD ["./go-web-template"]

