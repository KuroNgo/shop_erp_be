package principle

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"log"
	"shop_erp_mono/internal/config"
)

var Rbac *casbin.Enforcer

func SetUp(env *config.Database) *casbin.Enforcer {
	a, err := mongodbadapter.NewAdapter(fmt.Sprintf("mongodb+srv://%s:%s@andrew.8ulkv.mongodb.net/?retryWrites=true&w=majority", env.DBUser, env.DBPassword))
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
