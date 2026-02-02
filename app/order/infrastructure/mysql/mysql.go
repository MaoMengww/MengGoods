package mysql

import (
	"MengGoods/app/order/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"context"
	"time"

	"gorm.io/gorm"
)

type OrderDB struct {
	db *gorm.DB
}

func NewOrderDB(db *gorm.DB) *OrderDB {
	return &OrderDB{db: db}
}

func (d *OrderDB) GetOrderStatus(ctx context.Context, orderId int64) (int, error) {
	var order Orders
	err := d.db.WithContext(ctx).Where("order_id = ?", orderId).First(&order).Error
	if err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, "Get order status failed, err:"+err.Error())
	}
	return order.OrderStatus, nil
}

func (d *OrderDB) UpdateOrderStatus(ctx context.Context, orderId int64, status int) error {
	if err := d.db.WithContext(ctx).Model(&Orders{}).Where("order_id = ?", orderId).Update("order_status", status).Error; err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "Update order status failed, err:"+err.Error())
	}
	return nil
}

func (d *OrderDB) CreateOrder(ctx context.Context, order *model.Order, orderItem []*model.OrderItem) error {
	orderDB := &Orders{
		OrderId:          order.OrderId,
		UserId:           order.UserId,
		TotalPrice:       order.TotalPrice,
		PayPrice:         order.PayPrice,
		ReceiverName:     order.ReceiverName,
		ReceiverEmail:    order.ReceiverEmail,
		ReceiverProvince: order.ReceiverProvince,
		ReceiverCity:     order.ReceiverCity,
		ReceiverDetail:   order.ReceiverDetail,
		OrderStatus:      order.OrderStatus,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
		ExpireTime:       order.ExpireTime,
	}
	mqMsg := &MqMsg{
		OrderId:    order.OrderId,
		CreateTime: time.Now(),
		MsgStatus:  constants.MsgWaited,
	}
	orderItems := make([]*OrderItem, len(orderItem))
	for i, item := range orderItem {
		orderItems[i] = &OrderItem{
			OrderId:           item.OrderId,
			SellerId:          item.UserId,
			ProductId:         item.ProductId,
			ProductName:       item.ProductName,
			ProductImage:      item.ProductImage,
			ProductPrice:      item.ProductPrice,
			ProductNum:        item.ProductNum,
			ProductTotalPrice: item.ProductTotalPrice,
			ProductProperties: item.ProductProperties,
		}
	}
	tx := d.db.WithContext(ctx).Begin()
	if err := tx.Create(orderDB).Error; err != nil {
		tx.Rollback()
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "Create order failed, err:"+err.Error())
	}
	if err := tx.Create(orderItems).Error; err != nil {
		tx.Rollback()
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "Create order items failed, err:"+err.Error())
	}
	if err := tx.Create(mqMsg).Error; err != nil {
		tx.Rollback()
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "Create mq msg failed, err:"+err.Error())
	}
	tx.Commit()
	return nil
}

func (d *OrderDB) ViewOrderById(ctx context.Context, orderId int64) (*model.OrderWithItems, error) {
	var (
		orderDB      Orders
		orderItemsDB []*OrderItem
	)
	if err := d.db.WithContext(ctx).Where("order_id = ?", orderId).First(&orderDB).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "View order by id failed, err:"+err.Error())
	}
	if err := d.db.WithContext(ctx).Where("order_id = ?", orderId).Find(&orderItemsDB).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "View order items by id failed, err:"+err.Error())
	}
	order := &model.Order{
		OrderId:          orderDB.OrderId,
		UserId:           orderDB.UserId,
		TotalPrice:       orderDB.TotalPrice,
		PayPrice:         orderDB.PayPrice,
		ReceiverName:     orderDB.ReceiverName,
		ReceiverEmail:    orderDB.ReceiverEmail,
		ReceiverProvince: orderDB.ReceiverProvince,
		ReceiverCity:     orderDB.ReceiverCity,
		ReceiverDetail:   orderDB.ReceiverDetail,
		OrderStatus:      orderDB.OrderStatus,
		CreateTime:       orderDB.CreateTime,
		UpdateTime:       orderDB.UpdateTime,
		ExpireTime:       orderDB.ExpireTime,
	}
	var orderItems []*model.OrderItem
	for _, itemDB := range orderItemsDB {
		orderItem := &model.OrderItem{
			OrderItemId:       itemDB.OrderItemId,
			OrderId:           itemDB.OrderId,
			ProductId:         itemDB.ProductId,
			ProductName:       itemDB.ProductName,
			ProductImage:      itemDB.ProductImage,
			ProductPrice:      itemDB.ProductPrice,
			ProductNum:        itemDB.ProductNum,
			ProductTotalPrice: itemDB.ProductTotalPrice,
			ProductProperties: itemDB.ProductProperties,
		}
		orderItems = append(orderItems, orderItem)
	}
	return &model.OrderWithItems{
		Order:      order,
		OrderItems: orderItems,
	}, nil
}

func (d *OrderDB) ViewOrderList(ctx context.Context, status int, pageNum int, pageSize int) ([]*model.Order, error) {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var ordersDB []*Orders
	if err := d.db.WithContext(ctx).Where("user_id = ? AND order_status = ?", userId, status).Order("create_time desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&ordersDB).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "View order list failed, err:"+err.Error())
	}
	var orders []*model.Order
	for _, orderDB := range ordersDB {
		order := &model.Order{
			OrderId:          orderDB.OrderId,
			UserId:           orderDB.UserId,
			CreateTime:       orderDB.CreateTime,
			UpdateTime:       orderDB.UpdateTime,
			ExpireTime:       orderDB.ExpireTime,
			CancelTime:       orderDB.CancelTime,
			CancelReason:     orderDB.CancelReason,
			TotalPrice:       orderDB.TotalPrice,
			PayPrice:         orderDB.PayPrice,
			ReceiverName:     orderDB.ReceiverName,
			ReceiverEmail:    orderDB.ReceiverEmail,
			ReceiverProvince: orderDB.ReceiverProvince,
			ReceiverCity:     orderDB.ReceiverCity,
			ReceiverDetail:   orderDB.ReceiverDetail,
			OrderStatus:      orderDB.OrderStatus,
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (d *OrderDB) GetPendingMsgs(ctx context.Context) ([]*model.MqMsg, error) {
	var msgs []*MqMsg
	//获取创建时间超过10s且状态为未处理的消息
	if err := d.db.WithContext(ctx).Where("create_time < DATE_SUB(NOW(), INTERVAL 10 SECOND) AND msg_status = ?", constants.MsgWaited).Find(&msgs).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "Get pending mq msgs failed, err:"+err.Error())
	}
	var modelMsgs []*model.MqMsg
	for _, msg := range msgs {
		modelMsgs = append(modelMsgs, &model.MqMsg{
			MsgId:      msg.MsgId,
			OrderId:    msg.OrderId,
			CouponId:   msg.CouponId,
			CreateTime: msg.CreateTime,
			MsgStatus:  msg.MsgStatus,
		})
	}
	return modelMsgs, nil
}

func (d *OrderDB) MarkMsg(ctx context.Context, orderId int64) error {
	if err := d.db.WithContext(ctx).Model(&MqMsg{}).Where("order_id = ?", orderId).Update("msg_status", constants.MsgSended).Error; err != nil {
		return merror.NewMerror(merror.InternalDatabaseErrorCode, "Mark mq msg failed, err:"+err.Error())
	}
	return nil
}

func (d *OrderDB) GetPayAmount(ctx context.Context, orderId int64) (int64, error) {
	var orderDB Orders
	if err := d.db.WithContext(ctx).Where("order_id = ?", orderId).First(&orderDB).Error; err != nil {
		return 0, merror.NewMerror(merror.InternalDatabaseErrorCode, "Get pay amount failed, err:"+err.Error())
	}
	return orderDB.PayPrice, nil
}

func (d *OrderDB) GetOrderItem(ctx context.Context, orderItemId int64) (*model.OrderItem, error) {
	var itemDB OrderItem
	if err := d.db.WithContext(ctx).Where("order_item_id = ?", orderItemId).First(&itemDB).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "Get order item failed, err:"+err.Error())
	}
	return &model.OrderItem{
		OrderItemId:       itemDB.OrderItemId,
		OrderId:           itemDB.OrderId,
		ProductId:         itemDB.ProductId,
		ProductName:       itemDB.ProductName,
		ProductImage:      itemDB.ProductImage,
		ProductPrice:      itemDB.ProductPrice,
		ProductNum:        itemDB.ProductNum,
		ProductTotalPrice: itemDB.ProductTotalPrice,
		ProductProperties: itemDB.ProductProperties,
	}, nil
}

func (d *OrderDB) GetOrderInfo(ctx context.Context, orderId int64) (*model.Order, error) {
	var orderDB Orders
	if err := d.db.WithContext(ctx).Where("order_id = ?", orderId).First(&orderDB).Error; err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, "Get order info failed, err:"+err.Error())
	}
	return &model.Order{
		OrderId:          orderDB.OrderId,
		UserId:           orderDB.UserId,
		TotalPrice:       orderDB.TotalPrice,
		PayPrice:         orderDB.PayPrice,
		ReceiverName:     orderDB.ReceiverName,
		ReceiverEmail:    orderDB.ReceiverEmail,
		ReceiverProvince: orderDB.ReceiverProvince,
		ReceiverCity:     orderDB.ReceiverCity,
		ReceiverDetail:   orderDB.ReceiverDetail,
		OrderStatus:      orderDB.OrderStatus,
		CreateTime:       orderDB.CreateTime,
		UpdateTime:       orderDB.UpdateTime,
		ExpireTime:       orderDB.ExpireTime,
	}, nil
}
