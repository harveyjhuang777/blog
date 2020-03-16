package auth

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"github.com/jwjhuang/blog/service/app/logger"
	"github.com/jwjhuang/blog/service/model"
	"go.uber.org/dig"
)

var (
	accessInstance *Access
	accessPrefix   = "access"
	packet         authSet
)

type Access struct {
	*casbin.Enforcer
}

type authSet struct {
	dig.In

	DB *gorm.DB
}

func NewAccess(db *gorm.DB) {
	cas, err := casbin.NewEnforcer("./conf.d/rbac_model.conf")
	if err != nil {
		logger.Log().Panic(err)
	}

	adapter, err := gormadapter.NewAdapterByDBUsePrefix(db, accessPrefix+"_")
	if err != nil {
		logger.Log().Panic(err)
	}
	cas.SetAdapter(adapter)
	cas.EnableAutoSave(true)

	accessInstance = &Access{
		cas,
	}

	accessInstance.accessInitData()
}

func GetAccessInstance() *Access {
	return accessInstance
}

func (ac *Access) accessInitData() {
	policy := []*model.Policy{
		{Role: "general account", Path: "/api", Method: ".*"},
		{Role: "trial account", Path: "/api", Method: "GET"},
	}

	_ = ac.Enforcer.LoadPolicy()

	// for admin
	for _, value := range policy {
		ac.AddPolicy(value.Role, value.Path, value.Method)
	}
}
