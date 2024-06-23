package elastic_search

import (
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/elastic/go-elasticsearch/v8"
)

func GetESClient(config configs.Config) (*elasticsearch.Client, error) {
	if !config.ES.Enable {
		return nil, nil
	}

	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		zlog.Fatalf("elasticsearch.NewTypedClient failed, err:%v\n", err)
	}
	global.ESClient = client
	zlog.Infof("es connect success.")
	return client, nil
}
