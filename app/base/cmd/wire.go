//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	// "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"stb-library/app/base/internal/biz"
	"stb-library/app/base/internal/conf"
	"stb-library/app/base/internal/data"
	"stb-library/app/base/internal/server"
	"stb-library/app/base/internal/sgin"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, server.ProviderSet, sgin.ProviderSet, newApp))
}
