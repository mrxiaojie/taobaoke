package taobaoke

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type TopClient struct {
	TbAppKey string
	TbAppSecret string
	ParamMap map[string]interface{}
	IsTTLs bool
	IsSanBox bool
}


const (
	OpenRouterHttpAPI = "http://gw.api.taobao.com/router/rest"//正式版API路由地址 http 版本
	OpenRouterHttpsAPI = "https://eco.taobao.com/router/rest"//正式版API路由地址 https 版本
	OpenSanBoxRouterHttpAPI = "http://gw.api.tbsandbox.com/router/rest"//沙盒版API路由地址 http 版本
	OpenSanBoxRouterHttpsAPI = "https://gw.api.tbsandbox.com/router/rest"//沙盒版API路由地址 https 版本
	)

//设置淘宝客 appkey 和 appsecret
func (t *TopClient)SetConf(AppKey string,AppSecret string) {
	t.TbAppKey = AppKey
	t.TbAppSecret = AppSecret
}


//运行请求
func (t *TopClient) Exec(Api ApiAdtObj) ([]byte,error)  {
	var routerApi string
	var commonParam  = make(map[string]interface{})
	commonParam["app_key"] = t.TbAppKey
	commonParam["sign_method"] = "md5"
	commonParam["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	commonParam["format"] = "json"
	commonParam["v"] = "2.0"
	commonParam["method"] = Api.ApiName()
	for key,val:=range Api.GetParam(){
		commonParam[key] = val
	}
	commonParam["sign"] = sign(commonParam,t.TbAppSecret)

	if t.IsSanBox {
		if t.IsTTLs{
			routerApi = OpenSanBoxRouterHttpsAPI
		}else {
			routerApi = OpenSanBoxRouterHttpAPI
		}
	}else {
		if t.IsTTLs {
			routerApi = OpenRouterHttpsAPI
		}else {
			routerApi = OpenRouterHttpAPI
		}
	}

	respByte,err := request(routerApi,commonParam)

	return respByte,err;
}

//签名
func sign(param map[string]interface{},appSecret string)(str string)  {
	strArr := []string{}
	for key,_ :=range param{
		strArr = append(strArr,key)
	}
	sort.Strings(strArr)
	for _,key := range strArr{
		str = fmt.Sprintf("%s%s%v", str,string(key),param[key])
	}
	str = fmt.Sprintf("%s%s%s", appSecret,str,appSecret)
	s := md5.New()
	s.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(s.Sum(nil)))
}

//网络请求
func request(reqUrl string,param map[string]interface{}) (respo []byte,err error){
	var s string
	for key,_ := range param{
		s = fmt.Sprintf("%s%s%s%v%s",s,string(key),"=",param[key],"&")
	}
	s = strings.TrimRight(s,"&")//去除字符串中最右侧的 & 符号
	paramBody := strings.NewReader(s)

	resp,err := http.Post(reqUrl,"application/x-www-form-urlencoded; charset=utf-8",paramBody)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	respo,err = ioutil.ReadAll(resp.Body)
	return respo,err
}

