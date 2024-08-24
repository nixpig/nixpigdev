ARG GO_VERSION=1

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN make build


FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/tmp/bin/nixpigdev /usr/local/bin/
COPY --from=builder /usr/src/app/.ssh/id_ed25519 /usr/src/app/.ssh/id_ed25519
COPY --from=builder /usr/src/app/web/index.html /usr/src/app/web/index.html
COPY --from=builder /usr/src/app/.env /usr/src/app/.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080
EXPOSE 23234

CMD ["nixpigdev"]
