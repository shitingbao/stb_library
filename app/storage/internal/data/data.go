package data

import (
	"stb-library/app/storage/internal/conf"
	"stb-library/lib/ddb"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	d, err := ddb.OpenDb(c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		// no thing to close
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: d,
	}, cleanup, nil
}
