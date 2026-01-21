package mq

import (
	"MengGoods/app/stock/domain/model"
	"MengGoods/pkg/base/client"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)


type StockMq struct {
	*client.Kafka
}

func NewStockMq(client *client.Kafka) *StockMq {
	return &StockMq{
		Kafka: client,
	}
}

func (p *StockMq) SendLockStock(ctx context.Context, orderID int64, items []*model.StockItem) error {
	msg, err := json.Marshal(items)
	if err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json marshal error: %v", err))
	}
	return p.Publish(ctx, "stock_lock_stock", &kafka.Message{
		Key:   []byte(strconv.FormatInt(orderID, 10)),
		Value: msg,
	})
}

func (p *StockMq) SendUnlockStock(ctx context.Context, orderID int64, items []*model.StockItem) error {
	msg, err := json.Marshal(items)
	if err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json marshal error: %v", err))
	}
	return p.Publish(ctx, "stock_unlock_stock", &kafka.Message{
		Key:   []byte(strconv.FormatInt(orderID, 10)),
		Value: msg,
	})
}

func (p *StockMq) SendDeductStock(ctx context.Context, orderID int64, items []*model.StockItem) error {
	msg, err := json.Marshal(items)
	if err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json marshal error: %v", err))
	}
	return p.Publish(ctx, "stock_deduct_stock", &kafka.Message{
		Key:   []byte(strconv.FormatInt(orderID, 10)),
		Value: msg,
	})
}

func (p *StockMq) ConsumeLockStock(ctx context.Context, fn func(ctx context.Context, orderID int64, items []*model.StockItem) error) error {
	if err := p.Consumer("stock_lock_stock", "stock_lock_stock_group", func(ctx context.Context, msg *kafka.Message) error {
		var items []*model.StockItem
		if err := json.Unmarshal(msg.Value, &items); err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json unmarshal error: %v", err))
		}
		orderID, err := strconv.ParseInt(string(msg.Key), 10, 64)
		if err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("parse id error: %v", err))
		}
		return fn(ctx, orderID, items)	
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}

func (p *StockMq) ConsumeUnlockStock(ctx context.Context, fn func(ctx context.Context, orderID int64, items []*model.StockItem) error) error {
	if err := p.Consumer("stock_unlock_stock", "stock_unlock_stock_group", func(ctx context.Context, msg *kafka.Message) error {
		var items []*model.StockItem
		if err := json.Unmarshal(msg.Value, &items); err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json unmarshal error: %v", err))
		}
		orderID, err := strconv.ParseInt(string(msg.Key), 10, 64)
		if err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("parse id error: %v", err))
		}
		return fn(ctx, orderID, items)	
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}

func (p *StockMq) ConsumeDeductStock(ctx context.Context, fn func(ctx context.Context, orderID int64, items []*model.StockItem) error) error {
	if err := p.Consumer("stock_deduct_stock", "stock_deduct_stock_group", func(ctx context.Context, msg *kafka.Message) error {
		var items []*model.StockItem
		if err := json.Unmarshal(msg.Value, &items); err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("json unmarshal error: %v", err))
		}
		orderID, err := strconv.ParseInt(string(msg.Key), 10, 64)
		if err != nil {
			return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("parse id error: %v", err))
		}
		return fn(ctx, orderID, items)	
	}, ctx); err != nil {
		return merror.NewMerror(merror.InternalMqErrorCode, fmt.Sprintf("consumer error: %v", err))
	}
	return nil
}