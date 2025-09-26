package failover

import (
	"context"
	"sync/atomic"
	"webook/internal/service/sms"
)

type TimeoutFailoverSMSService struct {
	svcs []sms.Service //服务商
	cnt  int32         //连续超时的个数
	idx  int32
	//阈值
	//连续 超过这个数字 就要更换
	threshold int32
}

func NewTimeoutFailoverSMSService(svcs []sms.Service) *TimeoutFailoverSMSService {
	return &TimeoutFailoverSMSService{}
}

func (t TimeoutFailoverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	idx := atomic.LoadInt32(&t.idx)
	cnt := atomic.LoadInt32(&t.cnt)
	if cnt > t.threshold {
		//这里切换 新的下标
		newIdx := (idx + 1) % int32(len(t.svcs))
		if atomic.CompareAndSwapInt32(&t.idx, idx, newIdx) {
			//我成功忘后挪了一位
			atomic.StoreInt32(&t.cnt, 0)
		}
		//else 出现并发了 别人换了
		idx = atomic.LoadInt32(&t.idx)
	}
	svc := t.svcs[idx]
	err := svc.Send(ctx, tpl, args, numbers...)
	switch err {
	case context.DeadlineExceeded:
		//这里超时了 +1
		atomic.AddInt32(&t.cnt, 1)
		return err
	case nil:
		//连续超时状态被打断了
		atomic.StoreInt32(&t.cnt, 0)
		return nil
	default:
		//不知道什么错误
		//可以考虑换下一个
		return err
	}
}
