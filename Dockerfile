FROM golang:1.4

ENV GOPKG github.com/shurcooL/markdownfmt

RUN go get -v -d $GOPKG

WORKDIR /go/src/$GOPKG
COPY . /go/src/$GOPKG

RUN go install -v $GOPKG

CMD ["markdownfmt"]
