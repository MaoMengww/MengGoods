package mcontext

import (
	"MengGoods/pkg/merror"
	"context"
	"fmt"
	"strconv"

	"github.com/bytedance/gopkg/cloud/metainfo"
)

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	userInfo, ok := metainfo.GetPersistentValue(ctx, "uid")
	if !ok {
		return 0, merror.NewMerror(
			merror.ParamFromContextFailed,
			"failed to get user info",
		)
	}
	me, err := strconv.ParseInt(userInfo, 10, 64)
	if err != nil {
		return 0, merror.NewMerror(
			merror.ParamFromContextFailed,
			fmt.Sprintf("failed to parse uid, err:%v", err),
		)
	}
	return me, nil
}

func WithUserIDInContext(ctx context.Context, uid int64) context.Context {
	return metainfo.WithPersistentValue(ctx, "uid", strconv.FormatInt(uid, 10))
}

func WithStreamUserIDInContext(ctx context.Context, uid int64) context.Context {
	return metainfo.WithPersistentValue(ctx, "stream_uid", strconv.FormatInt(uid, 10))
}


