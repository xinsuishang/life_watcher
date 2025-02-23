package notice

// 短信通知实现
type smsNotifier struct {
	BaseNotifier
}

func init() {
	defaultRegister.Auto(&smsNotifier{
		BaseNotifier: BaseNotifier{
			Type: "sms",
		},
	})
}

func (n *smsNotifier) Notify(userID int64) error {

	return nil
}
