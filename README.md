## 廖雪峰-中华诗词网站抓取入库后 使用go开发api


## 修改config/test.json


## 运行服务

```
go build

注意:
$project_path为项目路径
-c config_path
-e test/dev/prod.json config目录下的配置文件名称

./shici -c $project_path/config -e test
```



