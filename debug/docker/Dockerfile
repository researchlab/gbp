FROM golang:1.15-alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache ca-certificates \
        gcc \
        git \
        musl-dev \
        openssh-client \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && go get github.com/derekparker/delve/cmd/dlv

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 5813

CMD [ "sh", "-c", "dlv debug --headless --continue --listen=:5813 --accept-multiclient --api-version=2 $CONFIG_FILE" ]
