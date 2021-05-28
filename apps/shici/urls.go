package shici

import (
	"github.com/gin-gonic/gin"
	
)



//注册路由
func RegisterRouter(r *gin.Engine)(string,map[string]map[string]gin.HandlerFunc){
	urlmap := make(map[string]map[string]gin.HandlerFunc)
	gets := make(map[string]gin.HandlerFunc)
	gets["/poems/"] = GetPoems  //获取所有诗词
	gets["/poeters/"] = GetPoeters //获取所有诗人
	gets["/dynastys/"] = GetDynastys //获取所有朝代
	urlmap["GET"] = gets
	return GroupUrl,urlmap
}





	
