package logic

import (
	"context"

	"fim_server/fim_file/file_api/internal/svc"
	"fim_server/fim_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageLogic {
	return &ImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageLogic) Image(req *types.ImageRequest) (resp *types.ImageResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ImageResponse)
	return
}
