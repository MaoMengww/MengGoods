package user

import (
	"MengGoods/app/user/controller/api"
	"MengGoods/app/user/domain/service"
	"MengGoods/app/user/infrastructure/cache"
	"MengGoods/app/user/infrastructure/cos"
	"MengGoods/app/user/infrastructure/mysql"
	"MengGoods/app/user/usecase"
	"MengGoods/config"
	"MengGoods/pkg/base/client"
)

func InjectUserServiceImpl() *api.UserServiceImpl {
	gormDB, err := client.NewMySQLClient(config.Conf.MySQL.UserDB)
	if err != nil {
		panic(err)
	}
	userDB := mysql.NewUserDB(gormDB)
	cacheClient, err := client.NewRedisClient()
	if err != nil {
		panic(err)
	}
	userCos := cos.NewUserCos(client.NewCosClient())
	userCache := cache.NewUserCache(cacheClient)
	svc := service.NewUserService(userDB, userCache, userCos)
	usecase := usecase.NewUserUsecase(userDB, userCache, svc, userCos)
	serviceImpl := api.NewUserServiceImpl(usecase)
	return serviceImpl
}
