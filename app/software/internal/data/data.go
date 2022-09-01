package data

import (
	"context"
	centralV1 "stb-library/api/central/v1"
	"stb-library/app/software/internal/conf"
	"stb-library/lib/ddb"
	"stb-library/lib/mongodb"
	"stb-library/lib/rediser"

	slog "stb-library/api/slog/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"gorm.io/gorm"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewCentralGrpcClient,
	NewCentralRepo,
	NewSlogServiceClient,
	NewLogServerHandleRepo,
	NewHealthRepo,
	NewUserRepo,
	NewCodeRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db          *gorm.DB
	rds         *redis.Client
	ce          centralV1.CentralClient
	slog        slog.LogServerClient
	mongoClient *mongodb.Mongodb
}

// NewData .
func NewData(c *conf.Data, l slog.LogServerClient, central centralV1.CentralClient) (*Data, func(), error) {
	d, err := ddb.OpenMysqlClient(c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	r, err := rediser.OpenRedisClient(c.Redis.Addr, c.Redis.Password, int(c.Redis.Level))
	if err != nil {
		return nil, nil, err
	}

	m, err := mongodb.OpenMongoDb(c.Mongo.Driver, "software")
	if err != nil {
		return nil, nil, err
	}
	da := &Data{
		db:          d,
		rds:         r,
		ce:          central,
		slog:        l,
		mongoClient: m,
	}

	return da, da.cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

// NewRegistrar 服务注册，对应了 main 里面的 server name
func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

// clear close all connect
func (d *Data) cleanup() {
	d.rds.Close()
}

// NewCentralGrpcClient
func NewCentralGrpcClient(r registry.Discovery, tp *tracesdk.TracerProvider) centralV1.CentralClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///central.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return centralV1.NewCentralClient(conn)
}

func NewMongoClient(conf *conf.Mongo) (*mongodb.Mongodb, error) {
	return mongodb.OpenMongoDb(conf.Driver, "software")
}
