package retryable

//type service struct {
//	svc      sms.Service
//	retryCnt int
//}
//
//func (s service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
//	err := s.svc.Send(ctx, tpl, args, numbers...)
//	for i := 0; i < s.retryCnt; i++ {
//		err = s.svc.Send(ctx, tpl, args, numbers...)
//		if err == nil {
//			return nil
//		}
//	}
//	return err
//}
