FROM golang:1.11-alpine3.8

RUN apk add --no-cache git

# https://github.com/russross/blackfriday/releases
ENV BLACKFRIDAY_VERSION v1.5.1

RUN git clone \
		-b "$BLACKFRIDAY_VERSION" \
		--depth 1 \
		https://github.com/russross/blackfriday.git \
		"$GOPATH/src/github.com/russross/blackfriday"

ENV GOPKG github.com/shurcooL/markdownfmt

RUN go get -v -d "$GOPKG" \
	&& rm -rv "$GOPATH/src/$GOPKG"

WORKDIR $GOPATH/src/$GOPKG
COPY . .

RUN go install -v $GOPKG

CMD ["markdownfmt"]
