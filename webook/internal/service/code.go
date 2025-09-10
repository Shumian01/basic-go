package service

import (
	"context"
	"fmt"
	"math/rand"
	"webook/internal/repository"
	"webook/internal/service/sms"
)

const codeTpliId = "1877556"

type CodeService struct {
	repo   *repository.CodeRepository
	smsSvc sms.Service
}

func NewCodeService(repo *repository.CodeRepository,
	smsSvc sms.Service) *CodeService {
	return &CodeService{
		repo:   repo,
		smsSvc: smsSvc,
	}
}

//Send 发验证码 我需要什么参数
func (svc *CodeService) Send(ctx context.Context,
	biz string, phone string) error {
	//phone_code:login:150xxxxx
	//code:$biz:150xxxxx
	//$biz:code:150xxxxx
	//两个步骤:1.生成验证码 2.发出去
	code := svc.geneerateCode()
	//塞进去redis
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	//前面成功了
	//发出去
	err = svc.smsSvc.Send(ctx, codeTpliId, []string{code}, phone)

	//if err != nil {
	//	//这个地方怎么办？
	//	//这意味着 Redis有这个验证码 但是没发出去给用户
	//	//能不能删掉验证码
	//	//你这个err 可能是超时的err	你都不知道发出去了没
	//	//在这里重试
	//	//要重试的话 初始化的时候 传入一个会重试的smsSvc
	//}

	return err
}

func (svc *CodeService) Verify(ctx context.Context, biz string,
	phone string, inputCode string) (bool, error) {
	return svc.repo.Verify(ctx, biz, phone, inputCode)
}

func (svc *CodeService) geneerateCode() string {
	//六位数. num在0,99999之间
	num := rand.Intn(1000000)
	//不够6位 加前导0
	return fmt.Sprintf("%06d", num)
}
