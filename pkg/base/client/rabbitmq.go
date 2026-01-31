package client

import (
	"MengGoods/config"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/logger"
	"MengGoods/pkg/merror"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMq struct {
	conn             *amqp.Connection
	ch               *amqp.Channel
	exchangeName     string
	DelayQueueName   string
	ProcessQueueName string
	routingKey       string
}

func NewRabbitMq(exchangeName string, delayQueueName string, processQueueName string, routingKey string) *RabbitMq {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v/", config.Conf.RabbitMQ.User, config.Conf.RabbitMQ.Password, config.Conf.RabbitMQ.Address))
	if err != nil {
		logger.Fatalf("创建rabbitmq连接失败%v", err)
		return nil
	}
	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("创建rabbitmq channel失败%v", err)
		return nil
	}
	//声明业务交换器
	err = ch.ExchangeDeclare(
		exchangeName, // 交换器名称
		"direct",     // 交换器类型
		true,         // 是否持久化
		false,        // 是否自动删除
		false,        // 是否内部使用
		false,        // 是否等待服务器确认
		nil,          // 其他属性
	)
	if err != nil {
		logger.Fatalf("声明业务交换器失败%v", err)
		return nil
	}
	//业务队列,有消费者
	_, err = ch.QueueDeclare(
		processQueueName, // 队列名称
		true,             // 是否持久化
		false,            // 是否自动删除
		false,            // 是否排他
		false,            // 是否等待服务器确认
		nil,
	)
	if err != nil {
		logger.Fatalf("声明死信队列失败%v", err)
		return nil
	}
	//绑定业务队列到交换器
	err = ch.QueueBind(
		processQueueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		logger.Fatalf("绑定死信队列到死信交换器失败%v", err)
		return nil
	}
	//缓冲队列,无消费者
	_, err = ch.QueueDeclare(
		delayQueueName, // 队列名称
		true,           // 是否持久化
		false,          // 是否自动删除
		false,          // 是否排他
		false,          // 是否等待服务器确认
		amqp.Table{
			"x-dead-letter-exchange":    exchangeName,
			"x-dead-letter-routing-key": routingKey,
			"x-message-ttl":             constants.DelayTime, // 消息过期时间,15分钟
		},
	)
	if err != nil {
		logger.Fatalf("声明缓冲队列失败%v", err)
		return nil
	}
	//绑定缓冲队列到死信交换器
	delayRoutingKey := routingKey + "_delay"
	err = ch.QueueBind(
		delayQueueName,  // 队列名称
		delayRoutingKey, // 路由键
		exchangeName,    // 交换器名称
		false,           // 是否等待服务器确认
		nil,             // 其他属性
	)
	if err != nil {
		logger.Fatalf("绑定普通队列到死信交换器失败%v", err)
		return nil
	}
	return &RabbitMq{
		conn:             conn,
		ch:               ch,
		exchangeName:     exchangeName,
		DelayQueueName:   delayQueueName,
		ProcessQueueName: processQueueName,
		routingKey:       delayRoutingKey,
	}
}

func (r *RabbitMq) Publish(msg []byte) error {
	err := r.ch.Publish(
		r.exchangeName,
		r.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return merror.NewMerror(merror.InternalRabbitMqErrorCode, err.Error())
	}
	return nil
}

func (r *RabbitMq) Consume(fn func(msg []byte) error) error {
	msgs, err := r.ch.Consume(
		r.ProcessQueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return merror.NewMerror(merror.InternalRabbitMqErrorCode, err.Error())
	}
	go func() {
		for msg := range msgs {
			if err := fn(msg.Body); err != nil {
				logger.Errorf("处理消息失败%v", err)
				continue
			}
			msg.Ack(false)
		}
	}()
	return nil
}

// 优雅关停
func (r *RabbitMq) Close() error {
	if err := r.ch.Close(); err != nil {
		return merror.NewMerror(merror.InternalRabbitMqErrorCode, err.Error())
	}
	if err := r.conn.Close(); err != nil {
		return merror.NewMerror(merror.InternalRabbitMqErrorCode, err.Error())
	}
	return nil
}
