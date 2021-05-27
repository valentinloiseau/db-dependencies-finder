FROM golang:1.15

ENV GO111MODULE=on
WORKDIR $GOPATH/src

COPY . .

ENV GOBIN /go/bin

CMD [ "app" ]