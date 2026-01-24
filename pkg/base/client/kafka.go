package client

import (
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"context"
	"fmt"
	"sync"

	mkafka "MengGoods/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	readers map[string]*kafka.Reader //string: topic+groupID
	writers map[string]*kafka.Writer //string: topic

	//优雅关停
	mu         sync.RWMutex
	consumerWg sync.WaitGroup
}

func NewKafka() *Kafka {
	return &Kafka{
		readers: make(map[string]*kafka.Reader),
		writers: make(map[string]*kafka.Writer),
	}
}

// 消费者处理函数, 业务逻辑需要实现这个接口
type KafkaHandle func(ctx context.Context, msg *kafka.Message) error

func (k *Kafka) Consumer(topic string, groupID string, handle KafkaHandle, ctx context.Context) error {
	k.mu.RLock()
	r, ok := k.readers[topic+groupID]
	k.mu.RUnlock()
	if !ok {
		var err error
		r, err = mkafka.NewReader(topic, groupID)
		if err != nil {
			logger.Errorf("创建reader失败%v", err)
			return err
		}
		k.mu.Lock()
		k.readers[topic+groupID] = r
		k.mu.Unlock()
	}
	k.consumerWg.Add(1)
	go func() {
		defer k.consumerWg.Done()
		for {
			msg, err := r.FetchMessage(ctx)
			if err != nil {
				logger.Errorf("读取失败%v", err)
				continue
			}
			//执行业务逻辑
			func() {
				err := handle(context.Background(), &msg)
				if err != nil {
					logger.Errorf("处理消息失败%v", err)
					return
				}
				r.CommitMessages(context.Background(), msg)
			}()
		}
	}()
	return nil
}

func (k *Kafka) Publish(ctx context.Context, topic string, msg *kafka.Message) error {
	k.mu.RLock()
	w, ok := k.writers[topic]
	k.mu.RUnlock()
	if !ok {
		var err error
		w, err = mkafka.NewWriter(topic)
		if err != nil {
			logger.Errorf("创建writer失败%v", err)
			return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("create writer error: %v", err))
		}
		k.mu.Lock()
		k.writers[topic] = w
		k.mu.Unlock()
	}
	if err := w.WriteMessages(ctx, *msg); err != nil {
		logger.Errorf("写入失败%v", err)
		return merror.NewMerror(merror.InternalKafkaErrorCode, fmt.Sprintf("write message error: %v", err))
	}
	return nil
}

// 优雅关停
func (k *Kafka) Close() {
	k.mu.Lock()
	defer k.mu.Unlock()
	for _, r := range k.readers {
		r.Close()
	}
	k.consumerWg.Wait()
	for _, w := range k.writers {
		w.Close()
	}
}
