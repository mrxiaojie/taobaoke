package tkapi
//淘宝客-公用-店铺关联推荐

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type ShopRecommend struct {
	ReqParam ShopRecommendParam
}


//请求参数
type ShopRecommendParam struct {
	Fields string
	UserId string
	Count int
	Platform int
}

//初始化api
func (t *ShopRecommend) Init()  {
	t.ReqParam.Fields = "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url" //需返回的字段列表
	t.ReqParam.UserId = "" //卖家Id,例如：123
	t.ReqParam.Platform = 1 //链接形式：1：PC，2：无线，默认：１
	t.ReqParam.Count = 20 //返回数量，默认20，最大值40
}

func (t *ShopRecommend) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["user_id"] = t.ReqParam.UserId
	paramMap["platform"] = t.ReqParam.Platform
	paramMap["count"] = t.ReqParam.Count
	return paramMap
}

func (t *ShopRecommend) ApiName() (s string)  {
	return "taobao.tbk.shop.recommend.get"
}


func (t *ShopRecommend) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.Fields =="" {
		return make(map[string]string) , errors.New("返回字段列表未设置")
	}
	if t.ReqParam.UserId =="" {
		return make(map[string]string) , errors.New("卖家ID未设置")
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