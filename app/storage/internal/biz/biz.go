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
	DefaultFilePath string
}
