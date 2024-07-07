package response

import (
	"context"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type JsonMsgResponse struct {
	Ctx *app.RequestContext
}

type JsonMsgResult struct {
	Code    int
	Message string
	Data    interface{}
}
type nilStruct struct{}

const SUCCESS_CODE = 200
const SUCCESS_MSG = "成功"
const ERROR_MSG = "错误"

func NewResponse(c *app.RequestContext) *JsonMsgResponse {
	return &JsonMsgResponse{Ctx: c}
}

func (r *JsonMsgResponse) Success(data interface{}) {
	res := JsonMsgResult{}
	res.Code = SUCCESS_CODE
	res.Message = SUCCESS_MSG
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

func (r *JsonMsgResponse) Error(ctx context.Context, mc MsgCode) {
	r.error(ctx, mc.Code, mc.Msg)
}

func (r *JsonMsgResponse) error(ctx context.Context, code int, message string) {
	if message == "" {
		message = ERROR_MSG
	}
	logData := ctx.Value(global.LOGGER_KEY_LOG)
	if logData == nil {
		logData = nilStruct{}
	}
	res := JsonMsgResult{}
	res.Code = code
	res.Message = message
	res.Data = logData
	r.Ctx.JSON(http.StatusOK, res)
}
