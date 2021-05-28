
## 获取朝代列表 首页用到 导航栏 朝代 dynastys.json

```
/api/shici/v1/dynastys/
get 请求
参数:
无
```



## 获取诗人列表页面 点击朝代后返回诗人列表 poeters.json

```
/api/shici/v1/poeters/
get 请求
参数:
currentPage=1
pageSize=10
dynastyId=朝代id
```


## 获取作品列表页面 点击朝代后返回诗人列表 poems.json

```
/api/shici/v1/poems/
get 请求
参数:
currentPage=1
pageSize=10
dynastyId=朝代id
poeterId=诗人id
```

