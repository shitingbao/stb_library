package data

import (
	"stb-library/app/central/internal/biz"

	// "stb-library/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type centraRepo struct {
	data *Data
	log  *log.Helper
}

// NewCentraRepo .
func NewCentraRepo(data *Data, logger log.Logger) biz.CentralRepo {
	return &centraRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *centraRepo) SayHelloData(name string) (string, error) {
	return name + ":hello nice!", nil
}

func (r *centraRepo) HealthyData(name string) (string, error) {
	return name + ":healthy body", nil
}
