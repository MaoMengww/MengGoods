package es

import (
	"MengGoods/app/product/domain/model"
	"MengGoods/kitex_gen/product"
	"MengGoods/pkg/merror"
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

type ProductEs struct {
	es *elastic.Client
}

func NewProductEs(es *elastic.Client) *ProductEs {
	return &ProductEs{es: es}
}

func (p *ProductEs) AddSpuItem(ctx context.Context, spu *model.SpuEs) error {
	_, err := p.es.Index().
		Index(viper.GetString("elasticsearch.productIndex")).
		Id(fmt.Sprintf("%v", spu.Id)).
		BodyJson(spu).
		Do(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("index error: %v", err))
	}
	return nil
}

func (p *ProductEs) UptateSpuItem(ctx context.Context, spu *model.SpuEs) error {
	_, err := p.es.Update().
		Index(viper.GetString("elasticsearch.productIndex")).
		Id(fmt.Sprintf("%v", spu.Id)).
		Doc(spu).
		Do(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("update error: %v", err))
	}
	return nil
}

func (p *ProductEs) DeleteSpuItem(ctx context.Context, spuId int64) error {
	_, err := p.es.Delete().
		Index(viper.GetString("elasticsearch.productIndex")).
		Id(fmt.Sprintf("%v", spuId)).
		Do(ctx)
	if err != nil {
		return merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("delete error: %v", err))
	}
	return nil
}

func (p *ProductEs) SearchSpu(ctx context.Context, req *product.GetSpuReq) ([]*model.SpuEs, int64, error) {
	boolQuery := elastic.NewBoolQuery()
	if req.Name != "" {
		boolQuery.Must(elastic.NewMatchQuery("name", req.Name))
	}
	if req.CategoryId != 0 {
		boolQuery.Must(elastic.NewTermQuery("category_id", req.CategoryId))
	}
	if req.MinPrice > req.MaxPrice {
		return nil, 0, merror.NewMerror(merror.QuerySettingInvalid, "minPrice must less than maxPrice")
	}
	if req.MinPrice > 0 {
		boolQuery.Must(elastic.NewRangeQuery("price").Gte(req.MinPrice))
	}
	if req.MaxPrice > 0 {
		boolQuery.Must(elastic.NewRangeQuery("price").Lte(req.MaxPrice))
	}
	searchResult, err := p.es.Search().
		Index(viper.GetString("elasticsearch.productIndex")).
		Query(boolQuery).                            // 放入查询条件
		Sort("price", true).                         // 按价格升序 (true=asc, false=desc)
		From(int((req.PageNum - 1) * req.PageSize)). // 分页：跳过多少条
		Size(int(req.PageSize)).                     // 分页：取多少条
		Do(ctx)
	if err != nil {
		return nil, 0, merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("search error: %v", err))
	}
	var spuList []*model.SpuEs
	for _, hit := range searchResult.Hits.Hits {
		var item model.SpuEs
		err := json.Unmarshal(hit.Source, &item)
		if err != nil {
			return nil, 0, merror.NewMerror(merror.InternalESErrorCode, fmt.Sprintf("unmarshal error: %v", err))
		}
		spuList = append(spuList, &item)
	}
	return spuList, searchResult.Hits.TotalHits.Value, nil
}
