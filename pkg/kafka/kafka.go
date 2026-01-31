package kafka

import (
	"MengGoods/config"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"net"

	"github.com/segmentio/kafka-go"
)

func GetConnection() *kafka.Conn {
	conn, err := kafka.Dial("tcp", config.Conf.Kafka.Addr)
	if err != nil {
		logger.Fatal("kafka connection error: ", err)
	}
	return conn
}

func NewReader(topic string, groupID string) (*kafka.Reader, error) {
	if topic == "" || groupID == "" {
		return nil, merror.NewMerror(merror.InternalKafkaErrorCode, "new reader topic or groupID is empty")
	}
	cfg := kafka.ReaderConfig{
		Brokers:     []string{config.Conf.Kafka.Addr},
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    constants.KafkaMinBytes,
		MaxBytes:    constants.KafkaMaxBytes,
		MaxWait:     constants.KafkaMaxWait,
		MaxAttempts: constants.KafkaMaxTries,
		//		Logger: ,
		Dialer: getKafkaDialer(),
	}
	return kafka.NewReader(cfg), nil
}

func NewWriter(topic string) (*kafka.Writer, error) {
	if topic == "" {
		logger.Fatal("new writer topic is empty")
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Conf.Kafka.Addr)
	if err != nil {
		return nil, merror.NewMerror(merror.InternalNetworkErrorCode, "kafka connection error: "+err.Error())
	}
	return &kafka.Writer{
		Addr:                   addr,
		Topic:                  topic,
		Balancer:               &kafka.RoundRobin{},     // 轮询策略
		MaxAttempts:            constants.KafkaMaxTries, // 最大尝试次数
		RequiredAcks:           kafka.RequireOne,        // 至少一个副本都确认
		Async:                  false,                   // 同步写入,保证数据安全
		AllowAutoTopicCreation: true,                    // 自动创建Topic
		Transport:              getKafkaTransport(),
	}, nil
}

func getKafkaDialer() *kafka.Dialer {
	return &kafka.Dialer{
		Timeout:   constants.KafkaReadTimeout,
		DualStack: true,
		/* SASLMechanism: plain.Mechanism{
			Username: config.Conf.Kafka.User,
			Password: config.Conf.Kafka.Password,
		}, */
	}
}

func getKafkaTransport() *kafka.Transport {
	return &kafka.Transport{
		/* SASL: plain.Mechanism{
			Username: config.Conf.Kafka.User,
			Password: config.Conf.Kafka.Password,
		}, */
	}
}
