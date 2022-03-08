package data

import (
	"context"
	"encoding/json"
	slog "stb-library/api/slog/v1"
	"stb-library/app/storage/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

const address = "124.70.156.31:5000"

var (
	defaultVersion          = "1.0.0"
	defaultErrorLevel int64 = 3.0
	defaultSys              = "stb"

	_ biz.SlogRepo = (*logServerHandle)(nil)
)

type logServerHandle struct {
	data *Data
}

func NewLogServerHandleRepo(d *Data) biz.SlogRepo {
	return &logServerHandle{data: d}
}

func NewSlogServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) slog.LogServerClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("discovery:///stb-library.slog.service"),
		grpc.WithEndpoint(address),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return slog.NewLogServerClient(conn)
}

// SendOneLog，插入标题和错误，标题就是本次错误主题，方便定位
func (l *logServerHandle) SendOneLog(topic string, err error) error {
	if err == nil {
		return err
	}
	da := &slog.RequestLogMessages{
		Sys: defaultSys,
		Msg: &slog.LogMessage{
			Topic:   topic,
			Content: err.Error(),
		},
		Level:   defaultErrorLevel,
		Version: defaultVersion,
		LogTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err = l.data.slog.SendOneLog(context.Background(), da)
	return err
}

// 任意类型的内容
func (l *logServerHandle) SendOneLogMes(topic string, content interface{}) error {
	mes, err := json.Marshal(content)
	if err != nil {
		return err
	}
	da := &slog.RequestLogMessages{
		Sys: defaultSys,
		Msg: &slog.LogMessage{
			Topic:   topic,
			Content: string(mes),
		},
		Level:   defaultErrorLevel,
		Version: defaultVersion,
		LogTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err = l.data.slog.SendOneLog(context.Background(), da)
	return err
}
