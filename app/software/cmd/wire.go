//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	// "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"stb-library/app/software/internal/biz"
	"stb-library/app/software/internal/conf"
	"stb-library/app/software/internal/data"
	"stb-library/app/software/internal/server"
	"stb-library/app/software/internal/sgin"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, server.ProviderSet, sgin.ProviderSet, newApp))
}
