package model

import "time"


type Spu struct {
	Id int64
	UserId int64
	Name string
	Description string
	CategoryId int64
	MainImageURL string
	SliderImageURLs string
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime *time.Time
	Status int32
	Price int64  //sku里最低的价格
	Skus []*Sku
}

type SpuEs struct {
	Id int64 `json:"id"`
	UserId int64 `json:"user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	CategoryId int64 `json:"category_id"`
	Price int64 `json:"price"`
	MainImageURL string `json:"main_image_url"`
}

type Sku struct {
	Id int64
	SpuId int64
	Name string
	Description string
	Properties string
	ImageURL string
	Price int64
	Sale int64
	CreateTime int64
	UpdateTime int64
	DeleteTime int64
}

type SkuEs struct {
	Id int64 `json:"id"`
	SpuId int64 `json:"spu_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Properties string `json:"properties"`
	ImageURL string `json:"image_url"`
	Price int64 `json:"price"`
	Sale int64 `json:"sale"`
}

type Category struct {
	Id int64
	Name string
	CreateTime int64
	UpdateTime int64
	DeleteTime int64
}
