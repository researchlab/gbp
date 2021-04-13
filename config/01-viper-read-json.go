package main

import (
	"fmt"

	"github.com/spf13/viper"
)

//定义config结构体
type Config struct {
	AppId  string
	Secret string
	Host   Host
}

//json中的嵌套对应结构体的嵌套
type Host struct {
	Address string
	Port    int
}

func main() {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("appId"))
	fmt.Println(config.GetString("secret"))
	fmt.Println(config.GetString("host.address"))
	fmt.Println(config.GetString("host.port"))

	//直接反序列化为Struct
	var configjson Config
	if err := config.Unmarshal(&configjson); err != nil {
		fmt.Println(err)
	}

	fmt.Println(configjson.Host)
	fmt.Println(configjson.AppId)
	fmt.Println(configjson.Secret)

}

//output
//-----------
//123456789
//maple123456
//localhost
//5799
//{localhost 5799}
//123456789
//maple123456
