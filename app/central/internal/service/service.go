package service

import "github.com/google/wire"

// ProviderSet is service providers.
// 唯一的服务，注册入 http、grpc
var ProviderSet = wire.NewSet(NewCentralService)
