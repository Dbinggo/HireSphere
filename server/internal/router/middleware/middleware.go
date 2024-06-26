package middleware

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/internal/router/manager"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func init() {
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_GLOBAL, AddTraceId, false)
}

// AddTraceId
//
//	@Description: add traced in logger
//	@return app.HandlerFunc
func AddTraceId() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 假设 Trace ID 存在于 HTTP Header "X-Trace-ID" 中
		traceID := ctx.Request.Header.Get("X-Request-ID")
		if traceID == global.COMMON_EMPTY_STRING {
			traceID = uuid.New().String()
		}
		c = zlog.NewContext(c, zap.String(global.LOGGER_KEY_TRACEID, traceID))
		ctx.Next(c)
	}
}
