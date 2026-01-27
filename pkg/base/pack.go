package base

import (
	"MengGoods/kitex_gen/model"
	"MengGoods/pkg/merror"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type DataResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func BuildBaseResp(err error) *model.BaseResp {
	if err == nil {
		return &model.BaseResp{
			Code:    merror.SuccessCode,
			Message: "Success",
		}
	}
	merr := merror.CoverError(err)
	return &model.BaseResp{
		Code:    int64(merr.Code),
		Message: merr.Msg,
	}
}

func ResErr(c *app.RequestContext, err error) {
	Err := merror.CoverError(err)
	c.JSON(http.StatusOK, ErrorResp{
		Code:    int64(Err.Code),
		Message: Err.Msg,
	})
}

func ResData(c *app.RequestContext, data any) {
	c.JSON(http.StatusOK, DataResp{
		Code:    merror.SuccessCode,
		Message: "Success",
		Data:    data,
	})
}

func ResList(c *app.RequestContext, data []any) {
	c.JSON(http.StatusOK, DataResp{
		Code:    merror.SuccessCode,
		Message: "Success",
		Data:    data,
	})
}
