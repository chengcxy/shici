

package shici

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"github.com/chengcxy/shici/apps"
	

)

//分页查询诗词 //默认唐代 李白
func GetPoems(c *gin.Context) {
    currentPage := 1
	pageSize := 50
	dynasty_id := 432345564227567616
	poeter_id := 474470375833468928
	if _currentPage := c.Query("currentPage");_currentPage !=""{
		currentPage = com.StrTo(_currentPage).MustInt()
	}
	if _pageSize := c.Query("pageSize");_pageSize !=""{
		pageSize = com.StrTo(_pageSize).MustInt()
	}
	if _dynasty_id := c.Query("dynastyId");_dynasty_id !=""{
		dynasty_id = com.StrTo(_dynasty_id).MustInt()
	}

	if _poeter_id := c.Query("poeterId");_poeter_id !=""{
		poeter_id = com.StrTo(_poeter_id).MustInt()
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

//查询单个朝代
func GetDynasty(c *gin.Context) {
	dynasty_id := com.StrTo(c.Param("dynasty_id")).MustInt()
	query := fmt.Sprintf(QUERY_DYNASTY_DETAIL,dynasty_id)
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


//分页查询诗人
func GetPoeters(c *gin.Context) {
    currentPage := 1
	pageSize := 100
	dynasty_id := 72057594037927936
	if _currentPage := c.Query("currentPage");_currentPage !=""{
		currentPage = com.StrTo(_currentPage).MustInt()
	}
	if _pageSize := c.Query("pageSize");_pageSize !=""{
		pageSize = com.StrTo(_pageSize).MustInt()
	}
	if _dynasty_id := c.Query("dynastyId");_dynasty_id !=""{
		dynasty_id = com.StrTo(_dynasty_id).MustInt()
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

//查询单个朝代
func GetPoeter(c *gin.Context) {
	poeter_id := com.StrTo(c.Param("poeter_id")).MustInt()
	query := fmt.Sprintf(QUERY_POETER_DETAIL,poeter_id)
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
