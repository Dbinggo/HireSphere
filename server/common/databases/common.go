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

type Find interface {
	KeyName() string
	Where() *gorm.DB
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

func FindInRedisOrDB(ctx context.Context, f Find) error {
	err := global.Rdb.HGetAll(ctx, f.KeyName()).Scan(f)
	if err != nil {
		zlog.InfofCtx(ctx, "redis中没有数据，从数据库中查找")
	} else {
		zlog.DebugfCtx(ctx, "从redis中查找到数据：%v", f)
		return nil
	}
	err = f.Where().Find(f).Error
	if err != nil {
		zlog.InfofCtx(ctx, "数据库中没有数据")
		return err
	}
	zlog.DebugfCtx(ctx, "从数据库中查找到数据：%v", f)
	err = global.Rdb.HSet(ctx, f.KeyName(), f).Err()
	if err != nil {
		zlog.InfofCtx(ctx, "redis中存储数据失败 :%v", err)
	}
	return nil
}
