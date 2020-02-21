package tkapi
//淘宝客-公用-商品关联推荐

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type ItemRecommend struct {
	ReqParam ItemRecommendParam
}


//请求参数
type ItemRecommendParam struct {
	Fields string
	NumIid string
	Count int
	Platform int
}

//初始化api
func (t *ItemRecommend) Init()  {
	t.ReqParam.Fields = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url" //需返回的字段列表
	t.ReqParam.NumIid = "" //商品Id
	t.ReqParam.Count = 20 //返回数量，默认20，最大值40
	t.ReqParam.Platform = 1 //链接形式：1：PC，2：无线，默认：１
}

func (t *ItemRecommend) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["num_iid"] = t.ReqParam.NumIid
	paramMap["count"] = t.ReqParam.Count
	paramMap["platform"] = t.ReqParam.Platform
	return paramMap
}

func (t *ItemRecommend) ApiName() (s string)  {
	return "taobao.tbk.item.recommend.get"
}


func (t *ItemRecommend) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.Fields =="" {
		return make(map[string]string) , errors.New("返回字符字段未设置")
	}
	if t.ReqParam.NumIid =="" {
		return make(map[string]string) , errors.New("商品ID未设置")
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