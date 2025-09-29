package retryable

import (
	"context"
	"errors"
	"webook/internal/service/sms"
)

type service struct {
	svc      sms.Service
	retryMax int
}

func NewService(svc sms.Service, retryMax int) sms.Service {
	return &service{
		svc:      svc,
		retryMax: retryMax,
	}
}
func (s service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	err := s.svc.Send(ctx, tpl, args, numbers...)
	cnt := 1
	for err != nil && cnt < s.retryMax {
		err = s.svc.Send(ctx, tpl, args, numbers...)
		if err == nil {
			return nil
		}
		cnt++
	}
	return errors.New("重试都失败了")
}

//设计并实现了一个高可用的短信服务平台
//1. 提高可用性：引入了重试机制 客户端限流 failover(轮询，实时监测)
//1.1实时检测:
//1.1.1 基于超时的实时检测(连续超时)
//1.1.2 基于响应时间的实时检测 (比如说 平均响应时间上升20%)
//1.1.3 基于长尾请求的实时检测 (比如说 响应时间超过1s的请求占比超过了10%)
//1.1.4 错误率
//2. 提高安全性
// 2.1 完整的资源申请与审批流程
// 2.2 鉴权
// 2.2.1 静态token
// 2.2.2 动态token
//3 提高可观测性: 日志 metrics tracing 丰富完善的排查手段
