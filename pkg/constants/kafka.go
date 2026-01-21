package constants

import "time"

const (
	KafkaMaxBytes = 10 * 1024 * 1024       // 10MB
	KafkaMinBytes = 10 * 1024              // 10KB
	KafkaMaxTries  = 3                      // 最大重试次数
	KafkaMaxWait  = 500 * time.Millisecond // 最大等待时间
	KafkaReadTimeout = 5 * time.Second // 读取超时时间

)