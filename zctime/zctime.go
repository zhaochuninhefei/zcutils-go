package zctime

type ZcTimeFormat string

func (ztf ZcTimeFormat) String() string {
	return string(ztf)
}

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	TIME_FORMAT_SIMPLE       ZcTimeFormat = "20060102150405"      // 简单格式: yyyyMMddHHmmss
	TIME_FORMAT_SLASH        ZcTimeFormat = "2006/01/02 15:04:05" // 斜杠格式: yyyy/MM/dd HH:mm:ss
	TIME_FORMAT_DASH         ZcTimeFormat = "2006-01-02 15:04:05" // 横杠格式: yyyy-MM-dd HH:mm:ss
	TIME_FORMAT_SIMPLE_MILLI ZcTimeFormat = "20060102150405.000"  // 带毫秒的简单格式: yyyyMMddHHmmss.SSS
)
