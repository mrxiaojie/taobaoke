package tkapi
// 淘宝客-公用-阿里妈妈推广券详情查询
//该接口一般用于查询从淘宝联盟上每天下载的excl中优惠券商品的优惠券信息

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type CouponInfo struct {
	ReqParam CouponInfoParam
}


//请求参数
type CouponInfoParam struct {
	Me string
	ItemId string
	ActivityId string
}

//初始化api
func (t *CouponInfo) Init()  {
	t.ReqParam.Me = "" //带券ID与商品ID的加密串,例如：nfr%2BYTo2k1PX18gaNN%2BIPkIG2PadNYbBnwEsv6mRavWieOoOE3L9OdmbDSSyHbGxBAXjHpLKvZbL1320ML%2BCF5FRtW7N7yJ056Lgym4X01A%3D
	t.ReqParam.ItemId = "" //商品ID,例如：123
	t.ReqParam.ActivityId = "" //券ID,例如：sdfwe3eefsdf
}

func (t *CouponInfo) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["me"] = t.ReqParam.Me
	paramMap["item_id"] = t.ReqParam.ItemId
	paramMap["activity_id"] = t.ReqParam.ActivityId
	return paramMap
}

func (t *CouponInfo) ApiName() (s string)  {
	return "taobao.tbk.coupon.get"
}


func (t *CouponInfo) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.Me =="" && t.ReqParam.ItemId=="" && t.ReqParam.ActivityId =="" {
		return make(map[string]string) , errors.New("加密字符串或者商品ID和券ID未设置")
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