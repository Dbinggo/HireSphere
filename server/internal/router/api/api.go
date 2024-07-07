package api

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/internal/handler/response"
	"github.com/Dbinggo/HireSphere/server/internal/router/manager"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

// only for test
func init() {
	manager.RouteHandler.RegisterRouter(manager.LEVEL_GLOBAL, func(r *route.RouterGroup) {
		r.GET("/test01", Test)
	})
}
func Test(ctx context.Context, c *app.RequestContext) {
	zlog.Infof("load - test")
	zlog.InfofCtx(ctx, "load ctx info - test")
	zlog.WarnfCtx(ctx, "load ctx warn - test")
	zlog.ErrorfCtx(ctx, "load ctx error - test")
	r := response.NewResponse(c)
	r.Error(ctx, response.MsgCode{
		Code: 500,
		Msg:  "test error",
	})
}
