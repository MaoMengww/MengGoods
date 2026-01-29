package mysql

import (
	"MengGoods/app/payment/domain/model"
	"MengGoods/pkg/constants"
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentDB struct {
	DB *gorm.DB
}

func NewPaymentDB(db *gorm.DB) *PaymentDB {
	return &PaymentDB{DB: db}
}

func(d *PaymentDB) CreatePaymentOrder(ctx context.Context, paymentOrder *model.PaymentOrder) error {
	paymentOrderDB := PaymentOrder{
		PaymentNo: paymentOrder.PaymentNo,
		OrderId:   paymentOrder.OrderId,
		UserId:    paymentOrder.UserId,
		Amount:    paymentOrder.Amount,
		Method:    paymentOrder.Method,
		Status:    paymentOrder.Status,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	return d.DB.WithContext(ctx).Create(&paymentOrderDB).Error
}

func(d *PaymentDB) CreateRefundOrder(ctx context.Context, refundOrder *model.PaymentRefund) error {
	refundOrderDB := PaymentRefund{
		RefundNo:    refundOrder.RefundNo,
		OrderItemID: refundOrder.OrderItemId,
		SellerID:    refundOrder.SellerId,
		UserId:      refundOrder.UserId,
		Amount:      refundOrder.Amount,
		Reason:      refundOrder.Reason,
		Status:      refundOrder.Status,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}
	return d.DB.WithContext(ctx).Create(&refundOrderDB).Error
}

func(d *PaymentDB) GetPaymentOrderByOrderId(ctx context.Context, orderId int64) (*model.PaymentOrder, error) {
	var paymentOrderDB PaymentOrder
	err := d.DB.WithContext(ctx).Where("order_id = ?", orderId).First(&paymentOrderDB).Error
	if err != nil {
		return nil, err
	}
	return &model.PaymentOrder{
		PaymentNo: paymentOrderDB.PaymentNo,
		OrderId:   paymentOrderDB.OrderId,
		UserId:    paymentOrderDB.UserId,
		Amount:    paymentOrderDB.Amount,
		Method:    paymentOrderDB.Method,
		Status:    paymentOrderDB.Status,
	}, nil
}

func(d *PaymentDB) GetRefundOrderByOrderItemId(ctx context.Context, orderItemId int64) (*model.PaymentRefund, error) {
	var refundOrderDB PaymentRefund
	err := d.DB.WithContext(ctx).Where("order_item_id = ?", orderItemId).First(&refundOrderDB).Error
	if err != nil {
		return nil, err
	}
	return &model.PaymentRefund{
		RefundNo:    refundOrderDB.RefundNo,
		OrderItemId: refundOrderDB.OrderItemID,
		SellerId:    refundOrderDB.SellerID,
		UserId:      refundOrderDB.UserId,
		Amount:      refundOrderDB.Amount,
		Reason:      refundOrderDB.Reason,
		Status:      refundOrderDB.Status,
	}, nil
}

func(d *PaymentDB) ConfirmPaymentOrder(ctx context.Context, paymentOrder *model.PaymentOrder) error {
	tx := d.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("order_id = ?", paymentOrder.OrderId).Update("status", constants.PaymentOrderStatusPaid).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&PaymentTransaction{
		TransactionNo:     paymentOrder.PaymentNo,
		OrderId:           paymentOrder.OrderId,
		UserId:            paymentOrder.UserId,
		TransactionAmount: paymentOrder.Amount,
		Type:              constants.PaymentTransactionTypePay,
		CreatedAt:         time.Now().Unix(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func(d *PaymentDB) ConfirmRefundOrder(ctx context.Context, refundOrder *model.PaymentRefund) error {
	tx := d.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("order_item_id = ?", refundOrder.OrderItemId).Update("status", constants.RefundStatusSuccess).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&PaymentTransaction{
		TransactionNo:     refundOrder.RefundNo,
		OrderId:           refundOrder.OrderItemId,
		UserId:            refundOrder.UserId,
		TransactionAmount: refundOrder.Amount,
		Type:              constants.PaymentTransactionTypeRefund,
		CreatedAt:         time.Now().Unix(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func(d *PaymentDB) UpdateRefundOrderStatus(ctx context.Context, orderItemId int64, status int) error {
	return d.DB.WithContext(ctx).Where("order_item_id = ?", orderItemId).Update("status", status).Error
}

func(d *PaymentDB) UpdatePaymentOrderStatus(ctx context.Context, paymentNo string, status int) error {
	return d.DB.WithContext(ctx).Where("payment_no = ?", paymentNo).Update("status", status).Error
}

