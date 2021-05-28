package apps

import (
	"errors"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"bytes"
	"time"
	"io/ioutil"
	"strings"
	"github.com/chengcxy/gotools/backends"
	"github.com/chengcxy/gotools/configor"
)


type Roboter struct {
	Config *configor.Config
}

func NewRoboter(config *configor.Config)(*Roboter){
	r := &Roboter{
		Config:config,
	}
	return r
}


func (r *Roboter) GetPayload(content string,mobile string) (string,map[string]interface{}){
	data := make(map[string]interface{})
	text := make(map[string]string)
	v,ok := r.Config.Get("roboter")
	if !ok {
		panic ("roboter is not exists")
	}
	rc := v.(map[string]interface{})
	hook_keyword := rc["hook_keyword"].(string)
	text["content"] = hook_keyword + content
	at := make(map[string]interface{})
	token := rc["token"].(string)
	if mobile == ""{
		at["atMobiles"] = rc["atMobiles"].([]interface{})
	}else{
		mobiles := strings.Split(mobile,",")
		at["atMobiles"] = mobiles
	}
	at["isAtAll"] = rc["isAtAll"].(bool)
	data["msgtype"] = "text"
	data["text"] = text
	data["at"] = at
	return token,data
}


func (r *Roboter) SendMsg(content string,mobile string)string{
	contentType := "application/json;charset=utf-8"
	token,data := r.GetPayload(content,mobile)
	api := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s",token)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
    jsonStr, _ := json.Marshal(data)
    resp, err := client.Post(api, contentType, bytes.NewBuffer(jsonStr))
    if err != nil {
        panic("send dingtalk msg err")
    }
    defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
    return string(result)

}


var Db *backends.MysqlClient
var Robot *Roboter



//钉钉报警机器人 后面做任务通知
func InitRoboter(c *configor.Config){
	Robot = NewRoboter(c)
}

//初始化数据库连接池 redis后面加白名单限制用
func InitDb(c *configor.Config) (*backends.MysqlClient){
	mc_interface,ok := c.Get("mysql")
	if !ok{
		panic(errors.New("config obj no mysql key"))
	}
	mc := backends.NewMysqlConfig(mc_interface)
	Db = backends.NewMysqlClient(mc)
	// redis_conf,_ := c.Get("redis")
	// fmt.Println(redis_conf)
	// pool_size := 20
	// redis_host := redis_conf.(map[string]interface{})["host"].(string)
	// redis_port := int(redis_conf.(map[string]interface{})["port"].(float64))
	// redis_password := redis_conf.(map[string]interface{})["password"].(string)
	// redis_db := int(redis_conf.(map[string]interface{})["db"].(float64))
    // RedisPool = redigo.NewPool(func() (redigo.Conn, error) {
	// 	c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d",redis_host,redis_port))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if _, err := c.Do("AUTH", redis_password); err != nil {
	// 		c.Close()
	// 		return nil, err
	// 	  }
	// 	  if _, err := c.Do("SELECT", redis_db); err != nil {
	// 		c.Close()
	// 		return nil, err
	// 	}
	// 	return c, nil
	// 	}, pool_size)
	return Db
}



func Query(sql,order_by string,currentPage,pageSize int) (map[string]interface{}){
	total_query := fmt.Sprintf("select count(1) as total from (%s) as a",sql)
	totals,_,_ := Db.Query(total_query)
	total := totals[0]["total"]
	query := fmt.Sprintf("%s %s limit %d,%d",sql,order_by,(currentPage-1)*pageSize,pageSize)
	log.Println("query-sql:\n",query)
	datas,_,_ := Db.Query(query)
	result := make(map[string]interface{})
	result["datas"] =  datas
	result["totalCount"] = total
	result["currentPage"] = currentPage
	result["pageSize"] = pageSize
	return result
}

