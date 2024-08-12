ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt update && apt install -y ca-certificates openssl

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/.ssh/id_ed25519 ./.ssh/id_ed25519
CMD ["run-app"]
