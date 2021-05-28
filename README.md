# 廖雪峰-中华诗词网站抓取入库后 使用go开发api


## 1.1 修改config/test.json


## 1.2 数据库表参见爬虫项目
[中华诗词爬虫](https://github.com/chengcxy/scrapy_spider/tree/master/tangshi)

## 1.3 运行服务

```
go build

注意:
$project_path为项目路径
-c config_path
-e test/dev/prod.json config目录下的配置文件名称

./shici -c $project_path/config -e test

运行成功输出:
2021/05/28 13:19:31 run in port:  :8088 WebMode: release
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

2021/05/28 13:19:31 find mapmode release

```

## 1.4 打开新终端 curl api

```
curl http://localhost:8088/api/shici/v1/dynastys/

response如下:

{
	"success": true,
	"code": 200,
	"msg": "获取数据成功",
	"data": {
		"currentPage": 1,
		"datas": [{
			"dynasty": "先秦",
			"dynasty_id": "72057594037927936"
		}, {
			"dynasty": "汉代",
			"dynasty_id": "144115188075855872"
		}, {
			"dynasty": "三国两晋",
			"dynasty_id": "216172782113783808"
		}, {
			"dynasty": "南北朝",
			"dynasty_id": "288230376151711744"
		}, {
			"dynasty": "隋代",
			"dynasty_id": "360287970189639680"
		}, {
			"dynasty": "唐代",
			"dynasty_id": "432345564227567616"
		}, {
			"dynasty": "宋代",
			"dynasty_id": "504403158265495552"
		}, {
			"dynasty": "元代",
			"dynasty_id": "576460752303423488"
		}, {
			"dynasty": "明代",
			"dynasty_id": "648518346341351424"
		}, {
			"dynasty": "清代",
			"dynasty_id": "720575940379279360"
		}, {
			"dynasty": "近现代",
			"dynasty_id": "792633534417207296"
		}],
		"pageSize": 500000,
		"totalCount": 11
	}
}

```



## 获取朝代列表 首页用到 导航栏 朝代

```
/api/shici/v1/dynastys/
get 请求
参数:
无
```


## api:

#### 获取朝代列表 

```
/api/shici/v1/dynastys/
get 请求
参数:
无
```

#### 获取某个朝代的诗人列表页面  

```
/api/shici/v1/dynasty/360287970189639680?currentPage=1&pageSize=10
get 请求
分页参数:
currentPage=1
pageSize=10
```


#### 获取某个诗人的作品列表页面 

```
/api/shici/v1/poeter/诗人id?currentPage=1&pageSize=10
get 请求
参数:
currentPage=当前页
pageSize=返回条数
```

#### 获取某个作品详情页 

```
/api/shici/v1/poem/作品id
get 请求
```
