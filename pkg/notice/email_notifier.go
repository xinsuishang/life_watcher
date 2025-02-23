package notice

type EmailNotifier struct {
	BaseNotifier
}

func (n *EmailNotifier) Notify(userID int64) error {
	// 实现邮件通知逻辑
	return nil
}

func init() {
	defaultRegister.Auto(&EmailNotifier{
		BaseNotifier: BaseNotifier{
			Type: "email",
		},
	})
}
