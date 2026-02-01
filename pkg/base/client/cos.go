package client

import (
	"MengGoods/config"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func NewCosClient() *cos.Client {
	u := "https://" + config.Conf.Cos.Bucket + ".cos." + config.Conf.Cos.Region + ".myqcloud.com/"
	url, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	c := cos.NewClient(&cos.BaseURL{BucketURL: url}, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Conf.Cos.SecretId,
			SecretKey: config.Conf.Cos.SecretKey,
		},
	})
	return c
}
