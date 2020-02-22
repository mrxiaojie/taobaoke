package tkapi
//淘宝客-推广者-店铺搜索

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type ShopSearch struct {
	ReqParam ShopSearchParam
}


//请求参数
type ShopSearchParam struct {
	Fields string
	KeyWorkQ string
	Sort string
	IsTmall bool
	StartCredit int
	EndCredit int
	StartCommissionRate int
	EndCommissionRate int
	StartTotalAction int
	EndTotalAction int
	StartAuctionCount int
	EndAuctionCount int
	Platform int
	PageNo int
	PageSize int
}

//初始化api
func (t *ShopSearch) Init()  {
	t.ReqParam.Fields = "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url" //需返回的字段列表
	t.ReqParam.KeyWorkQ = "" //查询词,例如：女装
	t.ReqParam.Sort = "commission_rate_des" //排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	t.ReqParam.IsTmall = false //是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	t.ReqParam.StartCredit = 1 //信用等级下限，1~20
	t.ReqParam.EndCredit = 20 //信用等级上限，1~20
	t.ReqParam.StartCommissionRate = 1 //淘客佣金比率下限，1~10000
	t.ReqParam.EndCommissionRate = 10000 //淘客佣金比率上限，1~10000
	t.ReqParam.StartTotalAction = 1 //店铺商品总数下限
	t.ReqParam.EndTotalAction = 100000 //店铺商品总数上限
	t.ReqParam.StartAuctionCount = 1 //累计推广商品下限
	t.ReqParam.EndAuctionCount = 1000000 //累计推广商品上限
	t.ReqParam.Platform = 1 //链接形式：1：PC，2：无线，默认：１
	t.ReqParam.PageNo = 1 //第几页，默认1，1~100
	t.ReqParam.PageSize = 20 //页大小，默认20，1~100
}

func (t *ShopSearch) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["q"] = t.ReqParam.KeyWorkQ
	paramMap["sort"] = t.ReqParam.Sort
	paramMap["is_tmall"] = t.ReqParam.IsTmall
	paramMap["start_credit"] = t.ReqParam.StartCredit
	paramMap["end_credit"] = t.ReqParam.EndCredit
	paramMap["start_commission_rate"] = t.ReqParam.StartCommissionRate
	paramMap["end_commission_rate"] = t.ReqParam.EndCommissionRate
	paramMap["start_total_action"] = t.ReqParam.StartTotalAction
	paramMap["end_total_action"] = t.ReqParam.EndTotalAction
	paramMap["start_auction_count"] = t.ReqParam.StartAuctionCount
	paramMap["end_auction_count"] = t.ReqParam.EndAuctionCount
	paramMap["platform"] = t.ReqParam.Platform
	paramMap["page_no"] = t.ReqParam.PageNo
	paramMap["page_size"] = t.ReqParam.PageSize
	return paramMap
}

func (t *ShopSearch) ApiName() (s string)  {
	return "taobao.tbk.shop.get"
}


func (t *ShopSearch) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.Fields =="" {
		return make(map[string]string) , errors.New("返回字段列表未设置")
	}
	if t.ReqParam.KeyWorkQ =="" {
		return make(map[string]string) , errors.New("查询词未设置")
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