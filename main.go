package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/kdniao-go/config"
)

func main() {

	expressData := ExpressData{
		ShipperCode:  "YTO",
		LogisticCode: "YT7454990157013",
	}

	// 获取配置参数
	auth := config.NewConfig()

	// 查询快递信息
	ExpressQuery(&expressData, auth)
}

// ExpressData 快递单号请求相关数据
type ExpressData struct {
	ShipperCode  string `json:"ShipperCode"`  // 快递公司编码
	LogisticCode string `json:"LogisticCode"` // 快递单号
	OrderCode    string `json:"OrderCode"`    // 订单编号

}

// ExpressQuery 查询快递信息
func ExpressQuery(expressData *ExpressData, auth *config.Conf) {

	// 快递单号数据编码成json格式
	rdata, err := json.Marshal(expressData)
	if err != nil {
		fmt.Println("JSON marshal fail", err)
	}

	// 获取DataSign
	DataSign := SignData(string(rdata), auth.KDNiao.APIKEY)

	// 请求数据
	data := url.Values{}
	data.Add("EBusinessID", auth.KDNiao.EBusinessID)
	data.Add("RequestType", auth.KDNiao.RequestType)
	data.Add("DataSign", DataSign)
	data.Add("DataType", auth.KDNiao.DataType)
	data.Add("RequestData", string(rdata))
	// 进行url编码
	//data.Encode()

	// 发送请求
	resp, err := http.PostForm("https://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx", data)
	if err != nil {
		fmt.Println("Request Fail:", err)
	}

	defer resp.Body.Close()

	// 读取所有响应数据
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", all)
}

// SignData 生成签名
func SignData(requestData string, key string) string {
	// 拼接数据
	data := requestData + key

	// 进行md5加密
	md5d := md5.Sum([]byte(data))
	// 将md5加密后的数据转换为32位字符串
	md5Data := fmt.Sprintf("%x", md5d)

	// 进行Base64编码
	encoding := base64.StdEncoding.EncodeToString([]byte(md5Data))

	// 进行url编码并返回
	//return url.QueryEscape(encoding)
	return encoding
}
