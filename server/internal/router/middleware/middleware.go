package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/Dbinggo/HireSphere/server/internal/handler/response"
	"github.com/Dbinggo/HireSphere/server/internal/model"
	"github.com/Dbinggo/HireSphere/server/internal/router/manager"
	"github.com/Dbinggo/HireSphere/server/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

func init() {
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_GLOBAL, AddTraceId, false)
	manager.RouteHandler.RegisterMiddleware(manager.LEVEL_V1, AuthToken, false)
}

// AddTraceId
//
//	@Description: add traced in logger
//	@return app.HandlerFunc
func AddTraceId() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		traceID := ctx.Request.Header.Get("X-Request-ID")
		if traceID == global.COMMON_EMPTY_STRING {
			traceID = uuid.New().String()
		}
		ctx.Response.Header.Set("X-Response-ID", traceID)
		c = zlog.NewContext(c, zap.String(global.LOGGER_KEY_TRACEID, traceID))
		ctx.Next(c)
	}
}

// AuthToken
//
//	@Description: auth user
//	@return app.HandlerFunc
func AuthToken() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		r := response.NewResponse(ctx)
		token := ctx.GetHeader("Authorization")
		if token == nil {
			// 无权限
			r.Ctx.Header("WWW-Authenticate", "Basic")
			r.Ctx.Status(401)
			zlog.WarnfCtx(c, "no permission !")
			ctx.Abort()
			return
		}

		/** 验证token是否合法与过期 */
		claims, err := utils.TokenVal(string(token))
		if err != nil {
			zlog.WarnfCtx(c, "token illegal ! err: %s", err.Error())
			r.Error(c, response.USER_ACCOUNT_ILLEGAL)

			ctx.Abort()
			return
		}
		if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
			zlog.WarnfCtx(c, "token is expired !")
			r.Error(c, response.TOKEN_IS_EXPIRED)
			ctx.Abort()
			return
		}
		// 假设 redis可以拿到用户信息
		cmd := global.Rdb.Get(c, fmt.Sprintf(global.REDIS_KEY_USER_FORMATE, claims.UserId))
		if cmd.Err() != nil {
			if errors.Is(cmd.Err(), redis.Nil) {
				zlog.WarnfCtx(c, "redis get token error ! err: %s", cmd.Err().Error())
				r.Error(c, response.TOKEN_IS_EXPIRED)
				ctx.Abort()
				return
			}
			ctx.Abort()
			return
		}
		user := &model.User{}
		err = utils.JsonToStruct(cmd.Val(), user)
		if err != nil {
			zlog.ErrorfCtx(c, "JsonToStruct failed ,msg: %s", err.Error())
			return
		}

		// 将 user 放到 context
		ctx.Set(global.LOGGER_KEY_USER, user)
	}
}
