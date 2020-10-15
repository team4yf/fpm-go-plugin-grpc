package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/team4yf/fpm-go-pkg/utils"
	pb "github.com/team4yf/fpm-go-plugin-grpc/biz"
	"github.com/team4yf/yf-fpm-server-go/fpm"
	"google.golang.org/grpc"
)

type grpcConfig struct {
	Port int16
}

type bizExecutorServer struct {
	pb.UnimplementedBizServiceServer
}

func (s *bizExecutorServer) Execute(ctx context.Context, req *pb.BizRequest) (rsp *pb.BizResponse, err error) {
	bizParam := fpm.BizParam{}
	if err = json.Unmarshal(([]byte)(req.Param), &bizParam); err != nil {
		return
	}
	data, err := fpm.Default().Execute(req.Name, &bizParam)
	if err != nil {
		return nil, err
	}
	rsp.Data = utils.JSON2String(data)
	return
}

func init() {
	fpm.RegisterByPlugin(&fpm.Plugin{
		Name: "fpm-plugin-grpc",
		V:    "0.0.1",
		Handler: func(fpmApp *fpm.Fpm) {
			config := grpcConfig{
				Port: 9091,
			}
			if fpmApp.HasConfig("grpc") {
				if err := fpmApp.FetchConfig("grpc", &config); err != nil {
					panic(err)
				}
			}

			fpmApp.Logger.Debugf("Startup : %s, config: %v", "grpc", config)

			// should run in another routine
			go func() {
				// grpc server
				lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
				if err != nil {
					panic(err)
				}
				grpcServer := grpc.NewServer()

				pb.RegisterBizServiceServer(grpcServer, &bizExecutorServer{})
				grpcServer.Serve(lis)
			}()

		},
	})
}
