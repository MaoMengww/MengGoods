package mysql

type User struct {
	UserId   int64  `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Role     int64  `gorm:"column:role"`
}

type Address struct {
	AddressId int64  `gorm:"primaryKey"`
	UserId   int64  `gorm:"column:user_id"`
	Province string `gorm:"column:province"`
	City     string `gorm:"column:city"`
	Detail   string `gorm:"column:detail"`
}
