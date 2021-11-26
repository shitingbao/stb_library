package data

import (
	"stb-library/app/storage/internal/conf"
	"stb-library/lib/ddb"
	"stb-library/lib/rediser"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"gorm.io/gorm"
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
