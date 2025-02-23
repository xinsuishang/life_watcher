package config

import "fmt"

const (
	// MySqlDSNFormat 支持MySQL/TiDB的DSN格式
	// 当使用TiDB时，可以通过额外参数进行调整
	MySqlDSNFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True%s%s"
)

func MySqlDSNFormatUtil(user, password, host string, port int, dbName string, tls string, extraParams string) string {
	tlsConfig := ""
	if tls != "" {
		tlsConfig = fmt.Sprintf("&tls=%s", tls)
	}

	if extraParams != "" && extraParams[0] != '&' {
		extraParams = "&" + extraParams
	}

	return fmt.Sprintf(MySqlDSNFormat, user, password, host, port, dbName, tlsConfig, extraParams)
}
