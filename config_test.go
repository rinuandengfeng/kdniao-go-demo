package main

import (
	"fmt"
	"testing"

	"github.com/kdniao-go/config"
)

func TestNewConfig(t *testing.T) {
	con := config.NewConfig()
	fmt.Println("apikey: ", con.KDNiao.APIKEY)
	fmt.Println("EBusinessID: ", con.KDNiao.EBusinessID)
	fmt.Println("RequestType: ", con.KDNiao.RequestType)
	fmt.Println("DataType: ", con.KDNiao.DataType)
}
