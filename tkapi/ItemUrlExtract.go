package tkapi
//淘宝客-公用-链接解析出商品id

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type ItemClickExtract struct {
	ReqParam ItemClickExtractParam
}


//请求参数
type ItemClickExtractParam struct {
	ClickUrl string
}

//初始化api
func (t *ItemClickExtract) Init()  {
	t.ReqParam.ClickUrl = "" //长链接或短链接 例如：https://s.click.taobao.com/***
}

func (t *ItemClickExtract) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["click_url"] = t.ReqParam.ClickUrl
	return paramMap
}

func (t *ItemClickExtract) ApiName() (s string)  {
	return "taobao.tbk.item.click.extract"
}


func (t *ItemClickExtract) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.ClickUrl == "" {
		return make(map[string]string) , errors.New("长链接或短链接未设置")
	}

	TopClient := taobaoke.TopClient{}
	TopClient.SetConf(appKey,appSecret)
	respByte,err :=TopClient.Exec(t)
	if resIsMap {
		mapDate := make(map[string]interface{})
		json.NewDecoder(bytes.NewBuffer(respByte)).Decode(&mapDate)
		return mapDate,err
	}
	return respByte,err
}