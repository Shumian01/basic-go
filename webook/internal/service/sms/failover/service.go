package failover

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
	"webook/internal/service/sms"
)

type FailOverSMSService struct {
	svcs []sms.Service
	idx  uint64
}

func NewFailOverSMSService(svcs []sms.Service) *FailOverSMSService {
	return &FailOverSMSService{
		svcs: svcs,
	}
}

func (f FailOverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	for _, svc := range f.svcs {
		err := svc.Send(ctx, tpl, args, numbers...)
		//发送成功
		if err == nil {
			return nil
		}

		//正常这边 输出日志
		//要做好监控
		log.Println(err)
	}
	return errors.New("发送失败, 所有服务商都尝试过")
}
func (f FailOverSMSService) SendV1(ctx context.Context, tpl string, args []string, numbers ...string) error {
	//我取下一个节点
	idx := atomic.AddUint64(&f.idx, 1)
	length := uint64(len(f.svcs))
	for i := idx; i < idx+length; i++ {
		err := f.svcs[int(i%length)].Send(ctx, tpl, args, numbers...)
		switch err {
		case nil:
			return nil
		case context.DeadlineExceeded, context.Canceled:
			return err
		default:
			//输出日志
		}
	}
	return errors.New("发送失败, 所有服务商都尝试过")
}
