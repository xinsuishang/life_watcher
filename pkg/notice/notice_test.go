package notice

import (
	"testing"
)

// 测试用的通知器实现
type TestNotifier struct{}

func (n *TestNotifier) GetType() string {
	return "test"
}

func (n *TestNotifier) Notify(userID int64) error {
	return nil
}

func TestNotifierRegistration(t *testing.T) {
	// 测试正常注册
	testNotifier := &TestNotifier{}
	if err := register(testNotifier.GetType(), testNotifier); err != nil {
		t.Errorf("Failed to register test notifier: %v", err)
	}

	// 测试重复注册
	if err := register(testNotifier.GetType(), testNotifier); err == nil {
		t.Error("Expected error for duplicate registration, got nil")
	}

	// 测试获取已注册的通知器
	if notifier, ok := GetNotifier("test"); !ok {
		t.Error("Failed to get registered notifier")
	} else if notifier != testNotifier {
		t.Error("Got wrong notifier instance")
	}

	// 测试获取所有通知器
	allNotifiers := GetAllNotifiers()
	if len(allNotifiers) < 1 {
		t.Error("Expected at least one registered notifier")
	}
}
