package service

import (
	"stb-library/app/software/internal/biz"

	v1 "stb-library/api/software/v1"
)

type Ghealth struct {
	// v1.UnimplementedGreeterServer
	v1.UnimplementedSoftwareServer

	uc *biz.GhealthUseCase
}

func NewGhealthServer(gc *biz.GhealthUseCase) *Ghealth {
	return &Ghealth{
		uc: gc,
	}
}

func (g *Ghealth) SayHello(name string) (string, error) {
	return g.uc.SayHello(name)
}
