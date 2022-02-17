package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewCentralUseCase,
	NewExportCase,
	NewFileComparisonCase,
	NewTransformCase,
	NewImageToWordCase,
	NewQrcodeCase,
	NewUserCase,
)

type DefaultFileDir struct {
	DefaultFilePath     string // 资源目录
	DefaultFileBasePath string // 执行基本目录
}
