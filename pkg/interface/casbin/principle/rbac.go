package principle

import (
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"log"
)

var Rbac *casbin.Enforcer

func SetUp() *casbin.Enforcer {
	a, err := mongodbadapter.NewAdapter("mongodb://localhost:27017/")
	if err != nil {
		log.Fatalln(err)
	}

	r, err := casbin.NewEnforcer("./internal/config/rbac_model.conf", a)
	if err != nil {
		log.Fatalln(err)
	}

	err = r.LoadPolicy()
	if err != nil {
		return nil
	}

	Rbac = r

	return r
}
