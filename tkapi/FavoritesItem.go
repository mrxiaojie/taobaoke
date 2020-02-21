package tkapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type FavoritesItem struct {
	ReqParam FavoritesItemParam
}


//请求参数
type FavoritesItemParam struct {
	pageNo int
	pageSize int
	Fields string
	Platform int
	AdZoneId int
	UnId string
	FavoritesId int
}

//初始化api
func (t *FavoritesItem) Init()  {
	t.ReqParam.pageNo = 1 //第几页，默认：1，从1开始计数
	t.ReqParam.pageSize = 20
	t.ReqParam.Fields = "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick,shop_title,zk_final_price_wap,event_start_time,event_end_time,tk_rate,status,type"
	t.ReqParam.Platform = 1 //链接形式：1：PC，2：无线，默认：１
	t.ReqParam.AdZoneId = 0 //推广位id，需要在淘宝联盟后台创建；且属于appkey备案的媒体id（siteid)下的推广位管理中的，例如:mm_400001_1111111_22222 ,这个22222就是推广位ID
	t.ReqParam.UnId = "0" //自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	t.ReqParam.FavoritesId = 0 //选品库的id
}

func (t *FavoritesItem) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["platform"] = t.ReqParam.Platform
	paramMap["page_no"] = t.ReqParam.pageNo
	paramMap["page_size"] = t.ReqParam.pageSize
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["adzone_id"] = t.ReqParam.AdZoneId
	paramMap["unid"] = t.ReqParam.UnId
	paramMap["favorites_id"] = t.ReqParam.FavoritesId
	return paramMap
}

func (t *FavoritesItem) ApiName() (s string)  {
	return "taobao.tbk.uatm.favorites.item.get"
}


func (t *FavoritesItem) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.FavoritesId <1 {
		return make(map[string]string) , errors.New("选品库ID未设置")
	}
	if t.ReqParam.AdZoneId <1 {
		return make(map[string]string) , errors.New("推广位id未设置")
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