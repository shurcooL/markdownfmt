FROM golang:1.4

ENV BLACKFRIDAY_VERSION v1.4

RUN git clone --single-branch \
		-b $BLACKFRIDAY_VERSION \
		https://github.com/russross/blackfriday.git \
		$GOPATH/src/github.com/russross/blackfriday

ENV GOPKG github.com/shurcooL/markdownfmt

RUN go get -v -d $GOPKG

WORKDIR $GOPATH/src/$GOPKG
COPY . $GOPATH/src/$GOPKG

RUN go install -v $GOPKG

CMD ["markdownfmt"]
