package utils

import (
	"os"
	"sync"

	"lonely-monitor/pkg/consts"

	"github.com/cloudwego/hertz/pkg/app"
)

var (
	requestKey string
	once       sync.Once
)

func GetTraceId(c *app.RequestContext) string {
	return c.Request.Header.Get(GetRequestKey())
}

func GetRequestKey() string {
	once.Do(func() {
		requestKey = os.Getenv(consts.RequestId)
		if requestKey == "" {
			requestKey = "X-Request-ID"
		}
	})
	return requestKey
}
