FROM golang:1.21-alpine3.19 AS build

WORKDIR /markdownfmt

COPY go.mod go.sum ./
RUN set -eux; go mod download; go mod verify

COPY . .

RUN set -eux; go build -v -trimpath -o markdownfmt ./; ./markdownfmt -h

FROM alpine:3.19

COPY --from=build /markdownfmt/markdownfmt /usr/local/bin/

CMD ["markdownfmt"]
