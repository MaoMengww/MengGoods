package user

import (
	"MengGoods/app/user/controller/api"
	"MengGoods/app/user/domain/service"
	"MengGoods/app/user/infrastructrue/cache"
	"MengGoods/app/user/infrastructrue/mysql"
	"MengGoods/app/user/usecase"
	"MengGoods/pkg/base/client"

	"github.com/spf13/viper"
)

func InjectUserServiceImpl() *api.UserServiceImpl {
	gormDB, err := client.NewMySQLClient(viper.GetString("mysql.userdb"))
	if err != nil {
		panic(err)
	}
	userDB := mysql.NewUserDB(gormDB)
	cacheClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	userCache := cache.NewUserCache(cacheClient)
	svc := service.NewUserService(userDB, userCache)
	usecase := usecase.NewUserUsecase(userDB, userCache, svc)
	serviceImpl := api.NewUserServiceImpl(usecase)
	return serviceImpl
}
