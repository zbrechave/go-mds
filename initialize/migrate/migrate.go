package migrate

import (
	"github.com/sirupsen/logrus"
	"github.com/zbrechave/go-mds/initialize/db"
	M "github.com/zbrechave/go-mds/model"
)

var err error

func Init() error {
	err = db.Engine.Sync2(new(M.Message))
	if err != nil {
		logrus.Error("数据库迁移失败")
		return err
	} else {
		logrus.Info("数据库迁移成功")
	}
	return nil
}
