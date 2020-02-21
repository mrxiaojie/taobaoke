package tkapi
//淘抢购api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type JuTqg struct {
	ReqParam JuTqgParam
}


//请求参数
type JuTqgParam struct {
	AdZoneId int
	Fields string
	StartTime string
	EndTime string
	PageNo int
	PageSize int
}

//初始化api
func (t *JuTqg) Init()  {
	t.ReqParam.AdZoneId = 0 //推广位id 需要在淘宝联盟后台创建；且属于appkey备案的媒体id（siteid)下的推广位管理中的，例如:mm_400001_1111111_22222 ,这个22222就是推广位ID
	t.ReqParam.Fields = "click_url,pic_url,reserve_price,zk_final_price,total_amount,sold_num,title,category_name,start_time,end_time" //链接形式：1：PC，2：无线，默认：１
	t.ReqParam.StartTime = "" //最早开团时间 例如：2016-08-09 09:00:00
	t.ReqParam.EndTime = "" //最晚开团时间 例如：2016-08-09 16:00:00
	t.ReqParam.PageNo = 1 //第几页，默认1，1~100
	t.ReqParam.PageSize = 40 //页大小，默认40，1~40
}

func (t *JuTqg) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["adzone_id"] = t.ReqParam.AdZoneId
	paramMap["fields"] = t.ReqParam.Fields
	paramMap["start_time"] = t.ReqParam.StartTime
	paramMap["end_time"] = t.ReqParam.EndTime
	paramMap["page_no"] = t.ReqParam.PageNo
	paramMap["page_size"] = t.ReqParam.PageNo
	return paramMap
}

func (t *JuTqg) ApiName() (s string)  {
	return "taobao.tbk.ju.tqg.get"
}


func (t *JuTqg) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.AdZoneId <1 {
		return make(map[string]string) , errors.New("推广位ID未设置")
	}
	if t.ReqParam.Fields == "" {
		return make(map[string]string) , errors.New("返回字段fields未设置")
	}
	if t.ReqParam.StartTime == "" {
		return make(map[string]string) , errors.New("最早开团时间未设置")
	}
	if t.ReqParam.EndTime == "" {
		return make(map[string]string) , errors.New("最晚开团时间未设置")
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