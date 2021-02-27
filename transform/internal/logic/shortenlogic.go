package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/hash"
	"shorturl/transform/model"

	"shorturl/transform/internal/svc"
	"shorturl/transform/transform"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {

	// 查询是否已经有了

	exists,_ := l.svcCtx.Model.FindByUrl(in.Url)
	fmt.Println(exists)
	if exists != nil{
		return &transform.ShortenResp{Shorten: "已经存在短连接了"+exists.Shorten},nil
	}

	// 手动代码开始，生成短链接
	key := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})
	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{
		Shorten: key,
	}, nil
	// 手动代码结束
}
