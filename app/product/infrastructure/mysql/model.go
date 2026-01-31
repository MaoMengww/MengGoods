package mysql

import "time"

// Spu结构体 - 字段改为大写导出
type Spu struct {
	ID              int64      `gorm:"column:spu_id;primaryKey"`
	Name            string     `gorm:"column:name"`
	Price           int64      `gorm:"column:price"`
	Description     string     `gorm:"column:description"`
	MainImageURL    string     `gorm:"column:main_image_url"`
	SliderImageURLs string     `gorm:"column:slider_image_urls"`
	Creator         int64      `gorm:"column:creator"`
	Category        int        `gorm:"column:category"`
	CreateAt        time.Time  `gorm:"column:created_at"`
	UpdateAt        time.Time  `gorm:"column:updated_at"`
	DeleteAt        *time.Time `gorm:"column:deleted_at"`
	Status          int32      `gorm:"column:status"`
}

// Sku结构体 - 字段改为大写导出
type Sku struct {
	ID          int64      `gorm:"primaryKey;column:id"`
	Name        string     `gorm:"column:name"`
	Price       int64      `gorm:"column:price"`
	Description string     `gorm:"column:description"`
	ImageURL    string     `gorm:"column:image_url"`
	Properties  string     `gorm:"column:properties"`
	Sale        int64      `gorm:"column:sale"`
	SpuID       int64      `gorm:"column:spu_id"`
	CreateAt    time.Time  `gorm:"column:created_at"`
	UpdateAt    time.Time  `gorm:"column:updated_at"`
	DeleteAt    *time.Time `gorm:"column:deleted_at"`
}

// 需要修复的结构体字段
type Category struct {
	ID       int64     `gorm:"primaryKey;column:id"`
	Name     string    `gorm:"column:name"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
	DeleteAt time.Time `gorm:"column:deleted_at"`
}
