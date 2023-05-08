FROM golang:1.19-alpine as build

ARG version=main

RUN apk add git make ncurses && \
    git clone https://github.com/waldirborbajr/ngtools.git $GOPATH/src/github.com/waldirborbajr/ngtools && \
    cd $GOPATH/src/github.com/waldirborbajr/ngtools && \
    git checkout $version

ENV GOPROXY=https://proxy.golang.org,direct
ENV GO111MODULE=on
ENV GOSUMDB=off

WORKDIR $GOPATH/src/github.com/waldirborbajr/ngtools

ENV PATH=$PATH:./bin

RUN make build

FROM alpine

COPY --from=build /go/src/github.com/waldirborbajr/ngtools/bin/ngtools /usr/local/bin/
RUN adduser -h /config -DG users -u 20000 ngt

USER ngt
ENTRYPOINT ["ngtools"]
