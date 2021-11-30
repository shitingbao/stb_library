package data

import (
	"stb-library/app/storage/internal/conf"
	"stb-library/lib/ddb"
	"stb-library/lib/rediser"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"gorm.io/gorm"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo, NewCentralGrpcClient)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	rds *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d, err := ddb.OpenMysqlClient(c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	r, err := rediser.OpenRedisClient(c.Redis.Addr, c.Redis.Password, int(c.Redis.Level))
	if err != nil {
		return nil, nil, err
	}
	da := &Data{
		db:  d,
		rds: r,
	}
	return da, da.cleanup, nil
}

// clear close all connect
func (d *Data) cleanup() {
	d.rds.Close()
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
