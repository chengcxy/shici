package shici



//公共结构体
type Response struct {
	Success bool `json:"success"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}


//api前缀
var GroupUrl string = "/api/shici/v1/"

//获取朝代
var BASE_DYNASTY_QUERY = `
select dynasty_id,dynasty
from spider.dynastys
`

//获取诗人
var BASE_POETER_QUERY = `
select dynasty_id,dynasty
from spider.poeters
where dynasty_id=%d
`

//获取作品
var BASE_POEM_QUERY = `
select a.*,b.poeter_name,b.poeter_desc,c.dynasty
from (
	select poem_id,poem_name,dynasty_id,poeter_id,contents
	from spider.china_poems 
	where dynasty_id=%d and poeter_id=%d
) as a
left join  spider.poeters as b on a.poeter_id=b.poeter_id
left join spider.dynastys as c on a.dynasty_id=c.dynasty_id
`