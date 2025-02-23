package notice

// 定义通知接口
type Notifier interface {
	Notify(userID int64) error
	GetType() string
}

type BaseNotifier struct {
	Type string
}

func (n *BaseNotifier) GetType() string {
	return n.Type
}

// GetNotifier 获取指定类型的通知器
func GetNotifier(typeName string) (Notifier, bool) {
	regLock.RLock()
	defer regLock.RUnlock()

	notifier, ok := registry[typeName]
	return notifier, ok
}

// GetAllNotifiers 获取所有已注册的通知器
func GetAllNotifiers() map[string]Notifier {
	regLock.RLock()
	defer regLock.RUnlock()

	// 返回一个副本以避免并发修改
	result := make(map[string]Notifier, len(registry))
	for k, v := range registry {
		result[k] = v
	}
	return result
}
