package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/transform/transformclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformclient.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return &types.ShortenResp{}, err
	}

	return &types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
}
