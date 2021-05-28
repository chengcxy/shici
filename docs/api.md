
## 获取朝代列表 首页用到 导航栏 朝代 dynastys.json

```
/api/shici/v1/dynastys/
get 请求
参数:
无
```



## 获取诗人列表页面 点击朝代后返回诗人列表 poeters.json

```
/api/shici/v1/dynasty/360287970189639680?currentPage=1&pageSize=10
get 请求
分页参数:
currentPage=1
pageSize=10
```


## 获取作品列表页面 点击朝代后返回诗人列表 poems.json

```
/api/shici/v1/poeter/诗人id?currentPage=1&pageSize=10
get 请求
参数:
currentPage=当前页
pageSize=返回条数
```

## 获取作品详情页 

```
/api/shici/v1/poem/作品id
get 请求
```


