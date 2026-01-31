package pack

import (
	domainModel "MengGoods/app/user/domain/model"
	"MengGoods/kitex_gen/model"
)

func BuildUserInfo(user *domainModel.User) *model.UserInfo {
	return &model.UserInfo{
		Id:       user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}

func BuildAddressList(address []*domainModel.Address) []*model.AddressInfo {
	addressList := make([]*model.AddressInfo, 0)
	for _, addr := range address {
		addressList = append(addressList, &model.AddressInfo{
			AddressID: addr.ID,
			Province:  addr.Province,
			City:      addr.City,
			Detail:    addr.Detail,
		})
	}
	return addressList
}

func BuildAddress(address *domainModel.Address) *model.AddressInfo {
	return &model.AddressInfo{
		AddressID: address.ID,
		Province:  address.Province,
		City:      address.City,
		Detail:    address.Detail,
	}
}
