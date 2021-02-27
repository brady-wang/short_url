package logic

import (
	"context"
	"shorturl/transform/transformer"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortLogic {
	return ShortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortLogic) Shorten(req types.ShortReq) (types.ShortResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return types.ShortResp{}, err
	}

	return types.ShortResp{
		Short: resp.Shorten,
	}, nil
	// 手动代码结束
}