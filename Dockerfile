FROM golang:1.16.5-buster as builder

RUN  mkdir -p $GOPATH/src/project

COPY . $GOPATH/src/project

EXPOSE 5001

CMD ["rock3"]
