package plugin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	pb "github.com/team4yf/fpm-go-plugin-grpc/biz"
)

func TestExecute(t *testing.T) {
	// fpmApp := fpm.New()

	// fpmApp.Init()

	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		t.Logf("err: %#v", err)
		return
	}
	defer conn.Close()
	client := pb.NewBizServiceClient(conn)

	data, err := client.Execute(context.Background(), &pb.BizRequest{
		Name:  "foo.bar",
		Param: "{}",
	})

	assert.Nil(t, err, "should not err")

	assert.NotNil(t, data, "should not nil")
}
