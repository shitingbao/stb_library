package data

import (
	"context"
	centralV1 "stb-library/api/central/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewCentralGrpcClient 这里嵌入 consule 连接，待定
func NewCentralGrpcClient() centralV1.GreeterClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("discovery:///beer.user.service"),
		// grpc.WithDiscovery(r),
		// grpc.WithMiddleware(
		// 	tracing.Client(tracing.WithTracerProvider(tp)),
		// 	recovery.Recovery(),
		// ),
	)
	if err != nil {
		panic(err)
	}
	return centralV1.NewGreeterClient(conn)
}
