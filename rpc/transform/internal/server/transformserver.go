// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package server

import (
	"context"

	"shorturl/rpc/transform/internal/logic"
	"shorturl/rpc/transform/internal/svc"
	"shorturl/rpc/transform/transform"
)

type TransformServer struct {
	svcCtx *svc.ServiceContext
}

func NewTransformServer(svcCtx *svc.ServiceContext) *TransformServer {
	return &TransformServer{
		svcCtx: svcCtx,
	}
}

func (s *TransformServer) Expand(ctx context.Context, in *transform.ExpandReq) (*transform.ExpandResp, error) {
	l := logic.NewExpandLogic(ctx, s.svcCtx)
	return l.Expand(in)
}

func (s *TransformServer) Shorten(ctx context.Context, in *transform.ShortenReq) (*transform.ShortenResp, error) {
	l := logic.NewShortenLogic(ctx, s.svcCtx)
	return l.Shorten(in)
}
