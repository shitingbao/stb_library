package biz

import "context"

type helloRepo interface {
	SayHello(context.Context, *Greeter) (string, error)
}

func (uc *GreeterUsecase) SayHello(ctx context.Context, g *Greeter) (string, error) {
	return uc.repo.SayHello(ctx, g)
}
