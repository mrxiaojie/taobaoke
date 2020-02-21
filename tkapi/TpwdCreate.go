package tkapi
//淘宝客-公用-淘口令生成
//提供淘客生成淘口令接口，淘客提交口令内容、logo、url等参数，生成淘口令关键key如：￥SADadW￥，后续进行文案包装组装用于传播

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/mrxiaojie/taobaoke"
)

type TpwdCreate struct {
	ReqParam TpwdCreateParam
}


//请求参数
type TpwdCreateParam struct {
	UserId string
	Text string
	Url string
	Logo string
	Ext string
}

//初始化api
func (t *TpwdCreate) Init()  {
	t.ReqParam.UserId = "" //生成口令的淘宝用户ID 例如：123
	t.ReqParam.Text = "" //口令弹框内容，长度大于5个字符
	t.ReqParam.Url = "" //口令跳转目标页,例如：https://uland.taobao.com/
	t.ReqParam.Logo = "" //口令弹框logoURL,例如：https://uland.taobao.com/
	t.ReqParam.Ext = "{}" //扩展字段JSON格式,默认为 {}
}

func (t *TpwdCreate) GetParam() map[string]interface{}  {
	paramMap := make(map[string]interface{})
	paramMap["user_id"] = t.ReqParam.UserId
	paramMap["text"] = t.ReqParam.Text
	paramMap["url"] = t.ReqParam.Url
	paramMap["logo"] = t.ReqParam.Logo
	paramMap["ext"] = t.ReqParam.Ext
	return paramMap
}

func (t *TpwdCreate) ApiName() (s string)  {
	return "taobao.tbk.tpwd.create"
}


func (t *TpwdCreate) Run(appKey string,appSecret string,resIsMap bool) (interface{},error)  {
	if t.ReqParam.Text == "" {
		return make(map[string]string) , errors.New("口令弹框内容未设置")
	}
	if t.ReqParam.Url == "" {
		return make(map[string]string) , errors.New("口令跳转目标页未设置")
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