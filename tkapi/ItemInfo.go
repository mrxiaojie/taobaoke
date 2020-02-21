package tkapi


import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type ItemInfo struct {
	ReqParam ItemInfoParam
}


//请求参数
type ItemInfoParam struct {
	NumIids string
	Platform int
	Ip string
}

//初始化api
func (t *ItemInfo) Init()  {
	t.ReqParam.NumIids = "" //商品ID串，用,分割，最大40个,例如：123,456
	t.ReqParam.Platform = 1 //链接形式：1：PC，2：无线，默认：１
	t.ReqParam.Ip = "" //ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供,例如：11.22.33.43
}

func (t *ItemInfo) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["num_iids"] = t.ReqParam.NumIids
	paramMap["platform"] = t.ReqParam.Platform
	paramMap["ip"] = t.ReqParam.Ip
	return paramMap
}

func (t *ItemInfo) ApiName() (s string)  {
	return "taobao.tbk.item.info.get"
}


func (t *ItemInfo) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.NumIids =="" {
		return make(map[string]string) , errors.New("商品ID串未设置")
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