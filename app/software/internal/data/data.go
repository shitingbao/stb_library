package data

import (
	"stb-library/app/software/internal/conf"
	"stb-library/lib/mongodb"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewCodeRepo,
)

// Data .
type Data struct {
	mongoClient *mongodb.Mongodb
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {

	m, err := mongodb.OpenMongoDb(c.Mongo.Driver, "software")
	if err != nil {
		return nil, nil, err
	}
	da := &Data{
		mongoClient: m,
	}

	return da, da.cleanup, nil
}

// clear close all connect
func (d *Data) cleanup() {
}

func NewMongoClient(conf *conf.Mongo) (*mongodb.Mongodb, error) {
	return mongodb.OpenMongoDb(conf.Driver, "software")
}
