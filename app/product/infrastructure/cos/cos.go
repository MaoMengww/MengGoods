package cos

import (
	"MengGoods/pkg/merror"
	"bytes"
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type ProductCos struct {
	client *cos.Client
}

func NewProductCos(client *cos.Client) *ProductCos {
	return &ProductCos{
		client: client,
	}
}

func (c *ProductCos) UploadSpuImage(ctx context.Context, spuImageData []byte, fileName string) (string, error) {	
	ext := path.Ext(fileName)
	objectKey := fmt.Sprintf("spu/%s/%s%s",
		time.Now().Format("20060102"),
		uuid.New().String(),
		ext,
	)
	reader := bytes.NewReader(spuImageData)
	_, err := c.client.Object.Put(ctx, objectKey, reader, nil)
	if err != nil {
		return "", merror.NewMerror(merror.InternalCosErrorCode, err.Error())
	}
	return fmt.Sprintf("%s/%s", c.client.BaseURL.BucketURL, objectKey), nil
}

func (c *ProductCos) UploadSkuImage(ctx context.Context, spuImageData []byte, fileName string) (string, error) {
	ext := path.Ext(fileName)
	objectKey := fmt.Sprintf("sku/%s/%s%s",
		time.Now().Format("20060102"),
		uuid.New().String(),
		ext,
	)
	reader := bytes.NewReader(spuImageData)
	_, err := c.client.Object.Put(ctx, objectKey, reader, nil)
	if err != nil {
		return "", merror.NewMerror(merror.InternalCosErrorCode, err.Error())
	}
	return fmt.Sprintf("%s/%s", c.client.BaseURL.BucketURL, objectKey), nil
}
