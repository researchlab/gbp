
client 

post 
```
go run client.go entity.go --update 
```

get 

```
go run client.go entity.go
```

server 

```
go run server.go entity.go &
```


解决问题

有一个google的protobuf包一直下载不了

go get google.golang.org/protobuf/proto
package google.golang.org/protobuf/proto: unrecognized import path "google.golang.org/protobuf/proto": https fetch: Get "https://google.golang.org/protobuf/proto?go-get=1": dial tcp 216.58.200.241:443: i/o timeout


解决方案

在 $GOPATH/src/ 目录下 mkdir google.golang.org  并且 cd google.golang.org/  
然后 

git clone https://github.com/protocolbuffers/protobuf-go.git

mv  protobuf-go  protobuf 就可以了
