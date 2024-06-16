package router

import (
	"fmt"
	"github.com/Dbinggo/HireSphere/server/global"
	_ "github.com/Dbinggo/HireSphere/server/internal/router/api"
	"github.com/Dbinggo/HireSphere/server/internal/router/manager"
	_ "github.com/Dbinggo/HireSphere/server/internal/router/middleware"
	"github.com/Dbinggo/HireSphere/server/log/zlog"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func RunServer() {
	h, err := listen()
	if err != nil {
		zlog.Errorf("Listen error: %v", err)
		panic(err.Error())
	}
	h.Spin()
}

func listen() (*server.Hertz, error) {

	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", global.Config.App.Host, global.Config.App.Port)))
	manager.RouteHandler.Register(h)
	return h, nil
}
