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
var GoodsId string
c.Ctx.Input.Bind(&GoodsId,"id")

var obj = tkapi.ItemRecommend{} //实例接口
obj.Init()//初始化默认参数
	
obj.ReqParam.NumIid = GoodsId//设置参数值

d,err := obj.Run(tbappkey,tbappSecret,true)//执行api请求淘宝数据

```


##taobao.tbk.item.info.get( 淘宝客-公用-淘宝客商品详情查询(简版) )


##taobao.tbk.uatm.favorites.item.get( 淘宝客-推广者-选品库宝贝信息 )


##taobao.tbk.uatm.favorites.get( 淘宝客-推广者-选品库宝贝列表 )


##taobao.tbk.ju.tqg.get( 淘抢购api )


##taobao.tbk.item.click.extract( 淘宝客-公用-链接解析出商品id )


##taobao.tbk.tpwd.create( 淘宝客-公用-淘口令生成 )


##taobao.tbk.shop.get( 淘宝客-推广者-店铺搜索 )


##taobao.tbk.shop.recommend.get( 淘宝客-公用-店铺关联推荐 )


##taobao.tbk.coupon.get( 淘宝客-公用-阿里妈妈推广券详情查询 )