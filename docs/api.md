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