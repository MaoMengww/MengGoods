package router

import (
	api "MengGoods/app/gateway/handler/api/user"
	"MengGoods/app/gateway/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitUserRouter(h *server.Hertz) {
	h.Use(mw.Sentinel())
	user := h.Group("/api/v1/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
	}
	user.Use(mw.AuthMiddleware())
	{
		user.POST("/addAddress", api.AddAddress)
		user.GET("/getAddress", api.GetAddressList)
		user.POST("/banUser", api.BanUser)
		user.POST("/unBanUser", api.UnBanUser)
		user.POST("/setAdmin", api.SetAdmin)
		user.GET("/getUserInfo", api.GetUserInfo)
		user.POST("/logOut", api.Logout)
		user.POST("/sendCode", api.SendCode)
		user.POST("/resetPassword", api.ResetPassword)
	}
}
