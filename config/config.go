package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Conf struct {
	KDNiao *KDNiao `yaml:"KDNiao"`
}

type KDNiao struct {
	EBusinessID string `yaml:"EBusinessID"` // 用户ID
	APIKEY      string `yaml:"APIKEY"`      // api-key
	RequestType string `yaml:"RequestType"` // 接口类型
	DataType    string `yaml:"DataType"`    // 响应类型

}

func NewConfig() *Conf {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("err %s ", err.Error())
	}

	var conf Conf

	// 解析配置
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("err %s ", err.Error())
	}

	return &conf
}
