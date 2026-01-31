package mysql

import "time"

type Stock struct {
	ID          int64     `gorm:"primaryKey;column:sku_id"`
	Stock       int32     `gorm:"column:stock"`
	LockedStock int32     `gorm:"column:locked_stock"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at"`
}

type StockJournal struct {
	ID        int64     `gorm:"primaryKey;column:journal_id"`
	SkuID     int64     `gorm:"column:sku_id"`
	OrderID   int64     `gorm:"column:order_id"`
	Count     int32     `gorm:"column:change_num"`
	Type      int32     `gorm:"column:change_type"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
