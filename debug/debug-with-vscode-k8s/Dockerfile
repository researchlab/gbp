# docker build -t webapp-go .
FROM golang:1.10.1
RUN go get -u -v github.com/derekparker/delve/cmd/dlv

ENV APP_PATH=/go/src/github.com/travishaagen/go-delve-remote-debug-example

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
COPY . $APP_PATH

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--accept-multiclient"]