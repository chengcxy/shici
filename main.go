package main

import (
	"net/http"
	"github.com/chengcxy/gotools/configor"
	"flag"
	"log"
	"github.com/chengcxy/shici/apps"
	"github.com/gin-gonic/gin"
	"github.com/chengcxy/shici/apps/shici"

)

// ./shici -c /data/config -e test



//命令行参数
var ConfigPath string
var Env string
var WebPort string
var WebMode string

func SetupRouters(r *gin.Engine,RegisterRouter func(r *gin.Engine)(string,map[string]map[string]gin.HandlerFunc)){
	group_url,map_fns := RegisterRouter(r)
	g := r.Group(group_url)
	for method,fns := range map_fns{
		if method == "POST"{
			for url,fn := range fns{
				g.POST(url,fn)
			}
		}else if method == "GET"{
			for url,fn := range fns{
				g.GET(url,fn)
			}
		}else if  method == "PUT"{
			for url,fn := range fns{
				g.PUT(url,fn)
			}
		}else if  method == "DELETE"{
			for url,fn := range fns{
				g.DELETE(url,fn)
			}
		}else{
			for url,fn := range fns{
				g.GET(url,fn)
			}
		}
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	var ModeMap = map[string]string{
		"test":gin.TestMode,
		"debug":gin.DebugMode,
		"release":gin.ReleaseMode,
	}
	mode,ok := ModeMap[WebMode]
	if !ok{
		mode = gin.DebugMode
	}
	log.Println("find mapmode",mode)
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	SetupRouters(r,shici.RegisterRouter)
	return r
}




func init(){
	const (
        default_conf_path = "/data/"
		default_conf_usage = "配置文件目录"  
		default_env = "test"
		default_env_usage = "json文件前缀"
    )
    flag.StringVar(&ConfigPath, "c", default_conf_path, default_conf_usage)
    flag.StringVar(&Env, "e", default_env, default_env_usage)
	flag.Parse()
	conf := configor.NewConfig(ConfigPath,Env)
	web_conf,_ := conf.Get("web")
	web_config := web_conf.(map[string]interface{})
	WebPort = web_config["port"].(string)
	WebMode = web_config["mode"].(string)
	log.Println("run in port: ",WebPort,"WebMode:",WebMode)
	apps.InitDb(conf)
	apps.InitRoboter(conf)


}

func main(){
	r :=  InitRouter()
	server := &http.Server{
		Addr:WebPort,
		Handler:  r,
	}
	server.ListenAndServe()
}