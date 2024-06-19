package router

import (
	"fmt"
	"github.com/Dbinggo/HireSphere/server/configs"
	_ "github.com/Dbinggo/HireSphere/server/internal/router/api"
	"github.com/Dbinggo/HireSphere/server/internal/router/manager"
	_ "github.com/Dbinggo/HireSphere/server/internal/router/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

var hServer *server.Hertz

func RunServer(config configs.Config) {
	hServer = server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)))
	manager.RouteHandler.Register(hServer)

	hServer.Spin()
}
