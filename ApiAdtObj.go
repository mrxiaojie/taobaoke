package taobaoke


type ApiAdtObj interface {
	Init()
	ApiName() string
	Run(appKey string,appSecret string,resIsMap bool) (interface{},error)
	GetParam() map[string]interface{}
}

