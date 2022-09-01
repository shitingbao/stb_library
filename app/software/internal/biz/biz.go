package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewCentralUseCase,
	NewSlogUseCase,
	NewUserCase,
	NewCodeCase,
)

type DefaultFileDir struct {
	DefaultAssetsPath string // 资源目录
	DefaultDirPath    string // 执行基本目录
}
