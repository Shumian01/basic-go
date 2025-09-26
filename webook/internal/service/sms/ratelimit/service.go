package ratelimit

import (
	"context"
	"fmt"
	"webook/internal/service/sms"
	"webook/pkg/ratelimit"
)

type RatelimitSMSService struct {
	svc     sms.Service
	limiter ratelimit.Limiter
}

func NewRatelimitSMSService(svc sms.Service, limiter ratelimit.Limiter) sms.Service {
	return &RatelimitSMSService{
		svc:     svc,
		limiter: limiter,
	}
}
func (s RatelimitSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	//你可以这里加一些代码 新特性
	limited, err := s.limiter.Limite(ctx, "sms:tencent")
	if err != nil {
		return fmt.Errorf("短信服务判断是否限流异常 %w", err)
	}
	if limited {
		return fmt.Errorf("触发了限流")
	}
	err = s.svc.Send(ctx, tpl, args, numbers...)
	//你在这里也可以加代码 新特性
	return err
}
