FROM golang

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

ENV GOPATH /opt/go:$GOPATH
ENV PATH /opt/go/bin:$PATH
ADD . /opt/go/src/local/myorg/myapp
WORKDIR /opt/go/src/local/myorg/myapp


RUN go get github.com/go-delve/delve/cmd/dlv
RUN go build -o main main.go
CMD ["./main"]
