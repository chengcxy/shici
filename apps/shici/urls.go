package shici

import (
	"github.com/gin-gonic/gin"
	
)



//注册路由
func RegisterRouter(r *gin.Engine)(string,map[string]map[string]gin.HandlerFunc){
	urlmap := make(map[string]map[string]gin.HandlerFunc)
	gets := make(map[string]gin.HandlerFunc)
	gets["/dynastys/"] = GetDynastys //获取所有朝代
	gets["/dynasty/:dynasty_id"] = GetDynastyPoeters //获取某个朝代下的所有诗人
	gets["/poeter/:poeter_id"] = GetPoeterPoems //获取诗人下的作品
	gets["/poem/:poem_id"] = GetPoemDetail  //获取诗词详情
	urlmap["GET"] = gets
	return GroupUrl,urlmap
}






	
