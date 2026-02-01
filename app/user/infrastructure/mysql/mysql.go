package mysql

import (
	"MengGoods/app/user/domain/model"
	"MengGoods/pkg/base/mcontext"
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"context"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) IsUserExist(ctx context.Context, username string) (bool, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&User{}).Where("username = ?", username).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
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
	return dbUser.UserId, nil
}

func (u *UserDB) GetUserByID(ctx context.Context, uid int64) (*model.User, error) {
	var user User
	err := u.db.WithContext(ctx).First(&user, uid).Error
	if err == gorm.ErrRecordNotFound {
		return nil, merror.NewMerror(merror.UserNotExist, "user not exist")
	}
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserId:    user.UserId,
		Username:  user.Username,
		Password:  user.Password,
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}

func (u *UserDB) GetUserByName(ctx context.Context, username string) (*model.User, error) {
	var user User
	err := u.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, merror.NewMerror(merror.UserNotExist, "user not exist")
	}
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserId:    user.UserId,
		Username:  user.Username,
		Password:  user.Password,
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}

func (u *UserDB) GetAddress(ctx context.Context, id int64) ([]*model.Address, error) {
	var addrs []*Address
	err := u.db.WithContext(ctx).Where("user_id = ?", id).Find(&addrs).Error
	if err == gorm.ErrRecordNotFound {
		return nil, merror.NewMerror(merror.AddressNotExist, "address not exist")
	}
	if err != nil {
		return nil, err
	}

	// 将数据库模型转换为领域模型
	result := make([]*model.Address, 0, len(addrs))
	for _, addr := range addrs {
		result = append(result, &model.Address{
			ID:       addr.AddressId,
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
		ID:       addr.AddressId,
		Province: addr.Province,
		City:     addr.City,
		Detail:   addr.Detail,
	}, nil
}

func (u *UserDB) AddAddress(ctx context.Context, addr *model.Address) (int64, error) {
	userId, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return 0, err
	}
	if err := u.db.WithContext(ctx).Create(&Address{
		AddressId: addr.ID,
		UserId:    userId,
		Province:  addr.Province,
		City:      addr.City,
		Detail:    addr.Detail,
	}).Error; err != nil {
		return 0, err
	}
	return addr.ID, nil
}

func (u *UserDB) SetUserAdmin(ctx context.Context, uid int64) error {
	err := u.db.WithContext(ctx).Model(&User{}).Where("user_id = ?", uid).Update("role", constants.Admin).Error
	if err == gorm.ErrRecordNotFound {
		return merror.NewMerror(merror.UserNotExist, "user not exist")
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) UpdatePassword(ctx context.Context, password string, uid int64) error {
	err := u.db.WithContext(ctx).Model(&User{}).Where("user_id = ?", uid).Update("password", password).Error
	if err == gorm.ErrRecordNotFound {
		return merror.NewMerror(merror.UserNotExist, "user not exist")
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDB) UploadAvatar(ctx context.Context, avatarURL string) error {
	uid, err := mcontext.GetUserIDFromContext(ctx)
	if err != nil {
		return err
	}
	err = u.db.WithContext(ctx).Model(&User{}).Where("user_id = ?", uid).Update("avatar", avatarURL).Error
	if err == gorm.ErrRecordNotFound {
		return merror.NewMerror(merror.UserNotExist, "user not exist")
	}
	if err != nil {
		return err
	}
	return nil
}
