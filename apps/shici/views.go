

package shici

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"github.com/chengcxy/shici/apps"
	

)




//查询所有朝代
func GetDynastys(c *gin.Context) {
    currentPage := 1
	pageSize := 500000
	order_by := " order by dynasty_id "
	result := apps.Query(BASE_DYNASTY_QUERY,order_by,currentPage,pageSize)
	result["totalCount"] = com.StrTo(result["totalCount"].(string)).MustInt()
	c.JSON(http.StatusOK, Response{
		Success:true,
		Code: 200,
		Msg: "获取数据成功",
		Data: result,
	})
	return 
}


//分页查询某个朝代下的诗人列表
func GetDynasty(c *gin.Context) {
    currentPage := 1
	pageSize := 100
	dynasty_id := com.StrTo(c.Param("dynasty_id")).MustInt()
	if _currentPage := c.query("currentPage");_currentPage !=""{
		currentPage = com.StrTo(_currentPage).MustInt()
	}
	if _pageSize := c.Query("pageSize");_pageSize !=""{
		pageSize = com.StrTo(_pageSize).MustInt()
	}
	order_by := " order by poeter_id"
	query := fmt.Sprintf(BASE_POETER_QUERY,dynasty_id)
	result := apps.Query(query,order_by,currentPage,pageSize)
	result["totalCount"] = com.StrTo(result["totalCount"].(string)).MustInt()
	c.JSON(http.StatusOK, Response{
		Success:true,
		Code: 200,
		Msg: "获取数据成功",
		Data: result,
	})
	return 
}


//分页查询诗词 查询某个诗人的作品
func GetPoeter(c *gin.Context) {
    currentPage := 1
	pageSize := 50
	poeter_id := com.StrTo(c.Param("poeterId")).MustInt()
	if _currentPage := c.Query("currentPage");_currentPage !=""{
		currentPage = com.StrTo(_currentPage).MustInt()
	}
	if _pageSize := c.Query("pageSize");_pageSize !=""{
		pageSize = com.StrTo(_pageSize).MustInt()
	}
	order_by := " order by a.poem_id "
	query := fmt.Sprintf(BASE_POEM_QUERY,dynasty_id,poeter_id)
	result := apps.Query(query,order_by,currentPage,pageSize)
	result["totalCount"] = com.StrTo(result["totalCount"].(string)).MustInt()
	c.JSON(http.StatusOK, Response{
		Success:true,
		Code: 200,
		Msg: "获取数据成功",
		Data: result,
	})
	return 
}

//查询单个作品
func GetPoem(c *gin.Context) {
	poem_id := com.StrTo(c.Param("poem_id")).MustInt()
	query := fmt.Sprintf(QUERY_POEM_DETAIL,poem_id)
	datas,_,_ := apps.Db.Query(query)
	result := datas[0]
	c.JSON(http.StatusOK, Response{
		Success:true,
		Code: 200,
		Msg: "获取数据成功",
		Data: result,
	})
	return 
}
