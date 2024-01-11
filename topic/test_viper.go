package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadIni() {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)
}

func ReadYml() {
	v := viper.New()
	v.AddConfigPath("./config") // 路径
	v.SetConfigName("config")   // 名称
	v.SetConfigType("yaml")     // 类型

	err := v.ReadInConfig() // 读配置
	if err != nil {
		panic(err)
	}

	s := v.GetString("db.username")
	fmt.Printf("s: %v\n", s)
}

func main() {
	ReadIni()
	ReadYml()

}
