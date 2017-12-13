FROM golang:1.9-alpine3.7

RUN apk add --no-cache git

# https://github.com/russross/blackfriday/releases
ENV BLACKFRIDAY_VERSION v1.5

RUN git clone \
		-b $BLACKFRIDAY_VERSION \
		--depth 1 \
		https://github.com/russross/blackfriday.git \
		$GOPATH/src/github.com/russross/blackfriday

ENV GOPKG github.com/shurcooL/markdownfmt

RUN go get -v -d $GOPKG \
	&& rm -rv "$GOPATH/src/$GOPKG"

WORKDIR $GOPATH/src/$GOPKG
COPY . $GOPATH/src/$GOPKG

RUN go install -v $GOPKG

CMD ["markdownfmt"]
