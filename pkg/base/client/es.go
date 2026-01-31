package client

import (
	"MengGoods/config"
	"MengGoods/pkg/merror"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
)

func NewEsClient() (*elastic.Client, error) {
	esConn := fmt.Sprintf("http://%s", config.Conf.Elasticsearch.Address)
	client, err := elastic.NewClient(
		elastic.SetURL(esConn),
		elastic.SetSniff(false),                              // 禁用节点嗅探，适用于单节点或Docker环境
		elastic.SetHealthcheck(true),                         // 启用健康检查
		elastic.SetHealthcheckTimeoutStartup(15*time.Second), // 增加启动超时时间到15秒
		elastic.SetHealthcheckTimeout(10*time.Second),        // 增加健康检查超时时间                  // 增加最大重试次数
		elastic.SetGzip(false),
	)
	if err != nil {
		return nil, merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("es clint failed,error: %v", err))
	}
	return client, nil
}
