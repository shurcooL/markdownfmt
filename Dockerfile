FROM golang:1.17-alpine3.15

WORKDIR /usr/src/markdownfmt

COPY go.mod go.sum ./
RUN set -eux; go mod download; go mod verify

COPY . .

RUN go install -v

CMD ["markdownfmt"]
