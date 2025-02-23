package notice

import (
	"fmt"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 注册表，存储类型名到接口实现的映射
var (
	registry = make(map[string]Notifier)
	regLock  sync.RWMutex
)

// register 注册函数，将类型名和对应的接口实现添加到注册表
func register(name string, impl Notifier) error {
	regLock.Lock()
	defer regLock.Unlock()

	hlog.Infof("register notifier: %s", name)

	// 检查是否已存在同名实现
	if existing, exists := registry[name]; exists {
		return fmt.Errorf("duplicate notifier type registration: %s (existing: %T, new: %T)",
			name, existing, impl)
	}

	registry[name] = impl
	return nil
}

// 注册器，用于注册通知器实现
type Register struct{}

var defaultRegister = &Register{}

// Auto 自动注册通知器实现
func (r *Register) Auto(notifier Notifier) {
	if err := register(notifier.GetType(), notifier); err != nil {
		// 在启动时就暴露问题
		panic(err)
	}
}
