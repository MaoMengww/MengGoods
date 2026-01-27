package mysql

import (
	"MengGoods/app/user/domain/model"
	"MengGoods/pkg/constants"
	"context"

	"gorm.io/gorm"
)

type User struct {
	Uid      int64  `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Role     int64  `gorm:"column:role"`
}

type Address struct {
	ID       int64  `gorm:"primaryKey"`
	Province string `gorm:"column:province"`
	City     string `gorm:"column:city"`
	Detail   string `gorm:"column:detail"`
}

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) IsUserExist(ctx context.Context, username string) (bool, error) {
	var count int64
	if err := u.db.WithContext(ctx).Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *UserDB) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	dbUser := &User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}
	if err := u.db.WithContext(ctx).Create(dbUser).Error; err != nil {
		return 0, err
	}
	return dbUser.Uid, nil
}

func (u *UserDB) GetUserByID(ctx context.Context, uid int64) (*model.User, error) {
	var user User
	if err := u.db.WithContext(ctx).First(&user, uid).Error; err != nil {
		return nil, err
	}
	return &model.User{
		Uid:      user.Uid,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (u *UserDB) GetUserByName(ctx context.Context, username string) (*model.User, error) {
	var user User
	if err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &model.User{
		Uid:      user.Uid,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (u *UserDB) GetAddress(ctx context.Context, id int64) ([]*model.Address, error) {
	var addrs []*Address
	if err := u.db.WithContext(ctx).Where("uid = ?", id).Find(&addrs).Error; err != nil {
		return nil, err
	}

	// 将数据库模型转换为领域模型
	result := make([]*model.Address, 0, len(addrs))
	for _, addr := range addrs {
		result = append(result, &model.Address{
			ID:       addr.ID,
			Province: addr.Province,
			City:     addr.City,
			Detail:   addr.Detail,
		})
	}
	return result, nil
}

func (u *UserDB) GetAddressByID(ctx context.Context, addressId int64) (*model.Address, error) {
	var addr Address
	if err := u.db.WithContext(ctx).First(&addr, addressId).Error; err != nil {
		return nil, err
	}
	return &model.Address{
		ID:       addr.ID,
		Province: addr.Province,
		City:     addr.City,
		Detail:   addr.Detail,
	}, nil
}

func (u *UserDB) AddAddress(ctx context.Context, addr *model.Address) (int64, error) {
	if err := u.db.WithContext(ctx).Create(&Address{
		ID:       addr.ID,
		Province: addr.Province,
		City:     addr.City,
		Detail:   addr.Detail,
	}).Error; err != nil {
		return 0, err
	}
	return addr.ID, nil
}

func (u *UserDB) SetUserAdmin(ctx context.Context, uid int64) error {
	if err := u.db.WithContext(ctx).Model(&User{}).Where("uid = ?", uid).Update("role", constants.Admin).Error; err != nil {
		return err
	}
	return nil
}
