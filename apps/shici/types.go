package shici

//公共结构体
type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//api前缀
var GroupUrl string = "/api/shici/v1/"

//获取所有朝代
var QUERY_DYNASTYS = `
select dynasty_id,dynasty
from spider.dynastys
`

//获取某个朝代下的诗人
var QUERY_DYNASTY_POETERS = `
select dynasty_id,poeter_id,poeter_name,poeter_desc
from spider.poeters
where dynasty_id=%d
`

//获取某个诗人下的作品
var QUERY_POETER_POEMS = `
select a.*,b.poeter_name,b.poeter_desc,c.dynasty
from (
	select poem_id,poem_name,dynasty_id,poeter_id,contents
	from spider.china_poems 
	where poeter_id=%d
) as a
left join  spider.poeters as b on a.poeter_id=b.poeter_id
left join spider.dynastys as c on a.dynasty_id=c.dynasty_id
`

//获取单个作品
var QUERY_POEM_DETAIL = `
select a.*,b.poeter_name,b.poeter_desc,c.dynasty
from (
	select poem_id,poem_name,dynasty_id,poeter_id,contents
	from spider.china_poems 
	where poem_id=%d
) as a
left join  spider.poeters as b on a.poeter_id=b.poeter_id
left join spider.dynastys as c on a.dynasty_id=c.dynasty_id
`
