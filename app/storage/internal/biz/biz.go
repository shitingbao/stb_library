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
	NewSlogUseCase,
	NewGhealthUseCase,
	NewImageZoomCase,
)

type DefaultFileDir struct {
	DefaultAssetsPath string // 资源目录
	DefaultDirPath    string // 执行基本目录
}
