ARG GO_VERSION=1

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .


FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /run-app /usr/local/bin/
COPY --from=builder /usr/src/app/.ssh/id_ed25519 /usr/src/app/.ssh/
COPY --from=builder /usr/src/app/web /usr/src/app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["run-app"]
