package images

import (
	"image"
	"image/jpeg"
	"os"
	"path"

	"github.com/nfnt/resize"
	"github.com/pborman/uuid"
)

type InterpolationFunction int

const (
	NearestNeighbor InterpolationFunction = iota
	Bilinear
	Bicubic
	MitchellNetravali
	Lanczos2
	Lanczos3
)

type option struct {
	URL           string
	Width         uint
	Height        uint
	Quality       int
	Interpolation int
}

type Option func(*option)

type Images struct {
	Width         uint
	Height        uint
	Quality       int
	Interpolation int
}

func NewImages(opt ...Option) *Images {
	op := &option{
		Width:   352,
		Height:  440,
		Quality: 100,
	}
	for _, o := range opt {
		o(op)
	}
	return &Images{
		Width:   op.Width,
		Height:  op.Height,
		Quality: op.Quality,
	}
}

func WithHeight(height uint) Option {
	return func(o *option) {
		o.Height = height
	}
}

func WithWidth(width uint) Option {
	return func(o *option) {
		o.Width = width
	}
}

func WithQuality(quality int) Option {
	return func(o *option) {
		o.Quality = quality
	}
}

func WithInterpolation(interpolation int) Option {
	return func(o *option) {
		o.Interpolation = interpolation
	}
}

func (i *Images) ImageZoom(url, assetsURL string) (string, error) {
	fl, err := os.Open(url)
	if err != nil {
		return "", err
	}
	info, err := fl.Stat()
	if err != nil {
		return "", err
	}
	img, _, err := image.Decode(fl)
	if err != nil {
		return "", err
	}

	m := resize.Resize(i.Width, i.Height, img, resize.InterpolationFunction(i.Interpolation))
	fname := path.Join(assetsURL, uuid.NewUUID().String()+info.Name())
	f, err := os.Create(fname) //"aa.jpeg"
	if err != nil {
		return "", err
	}
	return info.Name(), jpeg.Encode(f, m, &jpeg.Options{Quality: i.Quality})
}
