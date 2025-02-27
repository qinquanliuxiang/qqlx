package permissions

import (
	"fmt"
	"os"
	"qqlx/base/conf"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/casbin/casbin/v2/model"
)

func InitCasbin(cabinModelFile string) (e *casbin.Enforcer, err error) {
	_, err = os.Stat(cabinModelFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("casbin model file %s does not exist", cabinModelFile)
	}
	if err != nil {
		return nil, fmt.Errorf("stat casbin model file %s faild. err: %w", cabinModelFile, err)
	}

	dsn := conf.GetCasbinDsn()
	// 加载模型
	m, err := model.NewModelFromFile(cabinModelFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load model, %w", err)
	}
	// 加载策略
	a, err := gormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		return nil, fmt.Errorf("failed to load adapter, %w", err)
	}

	// 初始化casbin
	e, err = casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, fmt.Errorf("failed to init casbin, %w", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("failed to load policy, %w", err)
	}

	return e, nil
}
