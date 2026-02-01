package cos

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type UserCos struct {
	client *cos.Client
}

func NewUserCos(client *cos.Client) *UserCos {
	return &UserCos{
		client: client,
	}
}

func (c *UserCos) UploadAvatar(ctx context.Context, avatarData []byte, fileName string) (string, error) {
	// 1. 生成唯一路径: avatar/20260201/uuid.jpg
	ext := path.Ext(fileName)
	objectKey := fmt.Sprintf("avatar/%s/%s%s",
		time.Now().Format("20060102"),
		uuid.New().String(),
		ext,
	)

	// 2. 上传文件
	reader := bytes.NewReader(avatarData)
	_, err := c.client.Object.Put(ctx, objectKey, reader, nil)
	if err != nil {
		return "", err
	}
	// 3. 拼接 URL (仿照 DomTok GetImageUrl)
	return fmt.Sprintf("%s/%s", c.client.BaseURL.BucketURL, objectKey), nil
}
