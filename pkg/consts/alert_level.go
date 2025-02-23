package consts

type AlertLevel struct {
	Level    int
	Notifier string
}

var AlertLevels = []AlertLevel{
	{Level: 1, Notifier: "sms"},
	{Level: 2, Notifier: "email"},
}

var alertLevelMap = make(map[int]AlertLevel)

func init() {
	for _, alertLevel := range AlertLevels {
		alertLevelMap[alertLevel.Level] = alertLevel
	}
}

func GetAlertLevel(level int) (AlertLevel, bool) {
	alertLevel, ok := alertLevelMap[level]
	return alertLevel, ok
}
