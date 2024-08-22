package logic

import (
	"context"
	"runtime"
	"time"

	"github.com/asjard/benchmark/servers/go_zero/internal/svc"
	"github.com/asjard/benchmark/servers/go_zero/internal/types"
	"github.com/asjard/benchmark/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type Go_zeroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGo_zeroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Go_zeroLogic {
	return &Go_zeroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Go_zeroLogic) Go_zero(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	if l.svcCtx.Options.CpuBound {
		utils.Pow(l.svcCtx.Options.Target)
	} else {
		if l.svcCtx.Options.SleepTime > 0 {
			time.Sleep(l.svcCtx.Options.SleepTime)
		} else {
			runtime.Gosched()
		}
	}
	return &types.Response{
		Message: "hello",
	}, nil
}
