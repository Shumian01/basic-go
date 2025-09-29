package tencent

import (
	"context"
	"fmt"
	"webook/pkg/ratelimit"

	"github.com/ecodeclub/ekit"
	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId     *string
	signaName *string
	client    *sms.Client
	limiter   ratelimit.Limiter
}

func NewService(client *sms.Client, appId string, signName string, limiter ratelimit.Limiter) *Service {
	return &Service{
		client:    client,
		appId:     ekit.ToPtr[string](appId),
		signaName: ekit.ToPtr[string](signName),
		limiter:   limiter,
	}
}

//biz 直接代表的就是tplid

func (s Service) Send(ctx context.Context, biz string, args []string, numbers ...string) error {
	limited, err := s.limiter.Limite(ctx, "sms:tencent")
	if err != nil {
		return fmt.Errorf("短信服务判断是否限流异常 %w", err)
	}
	if limited {
		return fmt.Errorf("触发了限流")
	}
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = s.appId
	req.SignName = s.signaName
	req.TemplateId = ekit.ToPtr[string](biz)
	req.PhoneNumberSet = s.toStringPtrSlice(numbers)
	req.TemplateParamSet = s.toStringPtrSlice(args)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败 %s %s", *status.Code, *status.Message)
		}
	}
	return nil
}

func (s *Service) toStringPtrSlice(src []string) []*string {
	return slice.Map[string, *string](src, func(idx int, src string) *string {
		return &src
	})
}
