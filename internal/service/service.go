package service

import (
	"context"
)

type Service struct {
	ctx context.Context
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	return svc
}

func (svc *Service) Ctx() context.Context {
	return svc.ctx
}
