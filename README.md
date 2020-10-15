# fpm-go-plugin-grpc

### Install

`$ go get -u github.com/team4yf/fpm-go-plugin-grpc`

```golang

import _ "github.com/team4yf/fpm-go-plugin-grpc/plugin"

```

`$ sudo apt install -y protobuf-compiler`

cd `biz`

`protoc --go_out=. --go-grpc_out=. *.proto`


### Config

`conf/config.local.yaml`

```yaml
grpc:
    port: 9091
```

### Usage

```golang
import (
    "google.golang.org/grpc"
    pb "github.com/team4yf/fpm-go-plugin-grpc/biz"
)

conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
if err != nil {
    panic(err)
}
defer conn.Close()
client := pb.NewBizServiceClient(conn)

data, err := client.Execute(context.Background(), &pb.BizRequest{
    Name:  "foo.bar",
    Param: "{}",
})

```

