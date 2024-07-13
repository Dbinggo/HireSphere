package databases

// todo log抽象类接口
import (
	"context"
	"github.com/Dbinggo/HireSphere/server/common/log/zlog"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"gorm.io/gorm"
)

type DataBase interface {
	getDsn(config configs.Config) string
	initDataBases(config configs.Config) (*gorm.DB, error)
}

func InitDataBases(base DataBase, config configs.Config) {
	var err error
	global.DB, err = base.initDataBases(config)
	if err != nil {
		zlog.Fatalf("无法初始化数据库 %v", err)
		return
	}
	zlog.Infof("初始化数据库成功！")
	return
}

type rdbFindFunc func() error
type dbFindFunc func() error
type rdbSetFunc func() error
type showValueFunc interface{}

func FindInRedisOrDB(ctx context.Context, findRdb rdbFindFunc, findDB dbFindFunc, setRdb rdbSetFunc, value showValueFunc) error {
	err := findRdb()
	if err != nil {
		zlog.ErrorfCtx(ctx, "redis中未找到数据 %v", err)
		err = findDB()
		if err != nil {
			zlog.ErrorfCtx(ctx, "数据库中未找到数据 %v", err)
			return err
		} else {
			zlog.InfofCtx(ctx, "数据库中找到数据 %v", value)
			err = setRdb()
			if err != nil {
				zlog.ErrorfCtx(ctx, "redis中设置数据失败 %v", err)
				return err
			} else {
				zlog.InfofCtx(ctx, "redis中设置数据成功")
				return nil
			}
		}
	} else {
		zlog.InfofCtx(ctx, "redis中找到数据 %v", value)
		return nil
	}
}
