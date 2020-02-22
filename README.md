# taobaoke SDK
golang版的淘宝客SDK，长期更新.

- [x] [taobao.tbk.item.recommend.get( 淘宝客-公用-商品关联推荐 )](#taobaotbkitemrecommendget----)
- [x] [taobao.tbk.item.info.get( 淘宝客-公用-淘宝客商品详情查询(简版) )](#taobaotbkiteminfoget----)
- [x] [taobao.tbk.uatm.favorites.item.get( 淘宝客-推广者-选品库宝贝信息 )](#taobaotbkuatmfavoritesitemget----)
- [x] [taobao.tbk.uatm.favorites.get( 淘宝客-推广者-选品库宝贝列表 )](#taobaotbkuatmfavoritesget----)
- [x] [taobao.tbk.ju.tqg.get( 淘抢购api )](#taobaotbkjutqgget-api-)
- [x] [taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )](#taobaotbkitemclickextract---id-)
- [x] [taobao.tbk.tpwd.create( 淘宝客-公用-淘口令生成 )](#taobaotbktpwdcreate----)
- [x] [taobao.tbk.shop.get( 淘宝客-推广者-店铺搜索 )](#taobaotbkshopget----)
- [x] [taobao.tbk.shop.recommend.get( 淘宝客-公用-店铺关联推荐 )](#taobaotbkshoprecommendget----)
- [x] [taobao.tbk.coupon.get( 淘宝客-公用-阿里妈妈推广券详情查询 )](#taobaotbkcouponget----)


#API使用演示

###公共说明：
本SDK中所有的api调用都需要有以下三个步骤，其中Init()方法，必须要在实例接口后调用了，再操作其他的参数赋值
```
var obj = tkapi.ItemRecommend{}
obj.Init()
d,err := obj.Run(tbappkey,tbappSecret,true)
//Run(参数1，参数2，参数3)
//参数1，为淘宝客网站应用 appkey
//参数2,为淘宝客网站应用 appSecret
//参数3，设置Run接口返回数据结构是map还是[]byte,该参数值：true =>返回 map格式数据，false=>返回[]byte格式数据
```


##taobao.tbk.item.recommend.get( 淘宝客-公用-商品关联推荐 )
```
var obj = tkapi.ItemRecommend{} //实例接口
obj.Init()//初始化默认参数
	
obj.ReqParam.NumIid = 23222324 //设置参数值

d,err := obj.Run(tbappkey,tbappSecret,true)//执行api请求淘宝数据

```


##taobao.tbk.item.info.get( 淘宝客-公用-淘宝客商品详情查询(简版) )
```
var obj = tkapi.ItemInfo{}
obj.Init()
obj.ReqParam.NumIids = 1212323
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.uatm.favorites.item.get( 淘宝客-推广者-选品库宝贝信息 )
```
var obj = tkapi.FavoritesItem{}
obj.Init()
obj.ReqParam.FavoritesId = 2011111
obj.ReqParam.AdZoneId = 1100111111
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.uatm.favorites.get( 淘宝客-推广者-选品库宝贝列表 )
```
obj := &tkapi.Favorites{}
obj.Init()
d,err := obj.Run(tbappkey,tbappSecret,true)

```

##taobao.tbk.ju.tqg.get( 淘抢购api )
```
var obj = tkapi.JuTqg{}
obj.Init()
obj.ReqParam.AdZoneId = 1101111111
obj.ReqParam.StartTime = 2020-01-01 00:00:00
obj.ReqParam.EndTime = 2020-03-15 23:59:59
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )
```
var obj = tkapi.ItemClickExtract{}
obj.Init()
obj.ReqParam.ClickUrl = "http://xxxxx"
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.tpwd.create( 淘宝客-公用-淘口令生成 )
```
var obj = tkapi.TpwdCreate{}
obj.Init()
obj.ReqParam.Text = "这个是测试淘口令推广文字"
obj.ReqParam.Url = "https://uland.taobao.com/"
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.shop.get( 淘宝客-推广者-店铺搜索 )
```
var obj = tkapi.ShopSearch{}
obj.Init()
obj.ReqParam.KeyWorkQ = "男装"
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.shop.recommend.get( 淘宝客-公用-店铺关联推荐 )
```
var obj = tkapi.ShopRecommend{}
obj.Init()
obj.ReqParam.UserId = 2735800758
d,err := obj.Run(tbappkey,tbappSecret,true)
```

##taobao.tbk.coupon.get( 淘宝客-公用-阿里妈妈推广券详情查询 )
```
var obj = tkapi.CouponInfo{}
obj.Init()
obj.ReqParam.ItemId = "587636577617"
obj.ReqParam.ActivityId = "0f39649892b340eca7f7abc4b9ddffe3"
d,err := obj.Run(tbappkey,tbappSecret,true)
```