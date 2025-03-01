package utils

import (
	"os"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
)

var (
	requestKey string
	once       sync.Once
)

func getTraceId(c *app.RequestContext) string {
	return c.Request.Header.Get(getRequestKey())
}

func getRequestKey() string {
	once.Do(func() {
		requestKey = os.Getenv("APP_REQUEST_KEY")
		if requestKey == "" {
			requestKey = "Request-Id"
		}
	})
	return requestKey
}
