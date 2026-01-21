package user

import (
	"MengGoods/app/user/controller/rpc"
	"MengGoods/app/user/domain/service"
	"MengGoods/app/user/infrastructrue/cache"
	"MengGoods/app/user/infrastructrue/mysql"
	"MengGoods/app/user/usecase"
	"MengGoods/kitex_gen/user"
	"MengGoods/pkg/base/client"

	"github.com/spf13/viper"
)

func InjectUserServiceImpl() user.UserService {
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
	serviceImpl := rpc.NewUserServiceImpl(usecase)
	return serviceImpl
}