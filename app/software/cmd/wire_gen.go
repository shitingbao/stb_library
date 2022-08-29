// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"go.opentelemetry.io/otel/sdk/trace"
	"stb-library/app/software/internal/biz"
	"stb-library/app/software/internal/conf"
	"stb-library/app/software/internal/data"
	"stb-library/app/software/internal/server"
	"stb-library/app/software/internal/sgin"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	engine := sgin.NewGinEngine()
	hub := sgin.NewChatSocketfunc()
	httpServer := server.NewHTTPServer(confServer, engine, hub)
	defaultFileDir, err := sgin.ConstructorDefaultDir()
	if err != nil {
		return nil, nil, err
	}
	discovery := data.NewDiscovery(registry)
	logServerClient := data.NewSlogServiceClient(discovery, tracerProvider)
	centralClient := data.NewCentralGrpcClient(discovery, tracerProvider)
	dataData, cleanup, err := data.NewData(confData, logServerClient, centralClient)
	if err != nil {
		return nil, nil, err
	}
	centralRepo := data.NewCentralRepo(dataData)
	slogRepo := data.NewLogServerHandleRepo(dataData)
	slogUseCase := biz.NewSlogUseCase(defaultFileDir, slogRepo)
	centralUseCase := biz.NewCentralUseCase(centralRepo, slogUseCase)
	userRepo := data.NewUserRepo(dataData)
	userUseCase := biz.NewUserCase(userRepo, slogUseCase)
	sginSgin := sgin.NewSgin(defaultFileDir, engine, centralUseCase, userUseCase)
	grpcServer := server.NewGRPCServer(confServer, tracerProvider, sginSgin)
	registrar := data.NewRegistrar(registry)
	app := newApp(httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
