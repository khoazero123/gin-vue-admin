package core

import (
	"flag"
	"fmt"
	"gin-vue-admin/global"
	_ "gin-vue-admin/packfile"
	"gin-vue-admin/utils"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // Priority: Command Line> Environment Variable> Default
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("You are using config default, the path path of Config is%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("You are using the GVA_CONFIG environment variable, the path of Config is%v\n", config)
			}
		} else {
			fmt.Printf("You are using the value of the -c parameter passed by the command line, and the path of Config is%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("You are using FUNC VIPER (), the path of Config is%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
