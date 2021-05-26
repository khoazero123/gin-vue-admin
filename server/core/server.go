package core

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// Initialization Redis service
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// Ensure the text order output
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	welcome Gin-Vue-Admin
	current version:V2.4.1
    Plus mode: micro signal：shouzi_1994 QQ group：622360840
	Default automation document address:http://127.0.0.1%s/swagger/index.html
	Default front end file running address:http://127.0.0.1:8080
	If the project gives you a benefit, I hope you can invite the team to drink Cup.:https://www.gin-vue-admin.com/docs/coffee
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
