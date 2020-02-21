package tkapi

import (
	"bytes"
	"encoding/json"
	"github.com/mrxiaojie/taobaoke"
)

type Favorites struct {
	ReqParam FavoritesParam
}


//请求参数
type FavoritesParam struct {
	pageNo int
	pageSize int
	Fields string
	FavoritesType int
}

//初始化api
func (t *Favorites) Init()  {
	t.ReqParam.pageNo = 1
	t.ReqParam.pageSize = 20
	t.ReqParam.Fields = "favorites_title,favorites_id,type"
	t.ReqParam.FavoritesType = -1 //默认值-1；选品库类型，1：普通选品组，2：高佣选品组，-1，同时输出所有类型的选品组
}

func (t *Favorites) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["page_no"] = t.ReqParam.pageNo
	paramMap["page_size"] = t.ReqParam.pageSize
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["type"] = t.ReqParam.FavoritesType
	return paramMap
}

func (t *Favorites) ApiName() (s string)  {
	return "taobao.tbk.uatm.favorites.get"
}


func (t *Favorites) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
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