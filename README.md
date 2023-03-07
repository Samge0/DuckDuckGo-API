# DuckDuckGo-API
一个DuckDuckGo搜索api，可使用docker一键执行。
基于[DuckDuckGo-API](https://github.com/acheong08/DuckDuckGo-API)，感谢[Antonio Cheong（acheong08）](https://github.com/acheong08)

# 快速开始

# 基于源码运行（适合了解go语言编程的同学）

````
# 获取项目
$ git clone https://github.com/samge0/ddg-api-free.git

# 进入项目目录
$ cd ddg-api-free

# 复制配置文件
$ copy config.dev.json config.json

# 启动项目
$ go run main.go
````

# 使用docker运行
你可以使用docker快速运行本项目。[点击这里查看docker方式运行README.md](docker/README.md)
```shell
docker run -d \
--name ddg-api-free \
-p 8231:8080 \
-v `pwd`/docker_data/ddg-api-free/run.log:/app/run.log \
-e ACCESS_TOKEN=换成你接口请求的token（请求头中的xxx值，Authorization: Bearer xxx） \
-e PROXY="http://127.0.0.1:7890" \
-e ALLOW_ORIGIN=* \
--pull=always \
--restart always \
--memory=0.5G \
samge/ddg-api-free:latest
```
其中配置文件参考下边的配置文件说明。


# 配置文件说明

````
{
    "access_token": "自定义的api接口认证token",
    "allow_origin": "*",
    "host": "0.0.0.0",
    "port": "8080",
    "proxy": ""
}

access_token:接口请求的token（请求头中的xxx值，Authorization: Bearer xxx），默认为空，可自定义值
allow_origin:允许访问的域名，多个用英文逗号分隔，默认为*，允许所有
host：主机地址，例如：127.0.0.1、0.0.0.0，默认：0.0.0.0
port: http服务端口，默认8080
proxy: 如果不是部署在国外机子，则需要设置代理访问DuckDuckGo，代理格式：http://127.0.0.1:7890
````

【注意】：环境变量的优先级高于config.json，如果二者同时配置，则优先取环境变量的值。


### 接口访问
- 接口使用请查看`test_main.http`测试文件

### 有疑问请添加微信（备注: ddg-api-free），不定期通过解答

**微信号 SamgeApp **


### 源README.md
API server and module for DuckDuckGo

```go
type Search struct {
	Query     string `json:"query"`
	Region    string `json:"region"`
	TimeRange string `json:"time_range"`
	Limit     int    `json:"limit"`
}

type Result struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}
```

Send requests via GET or POST in JSON
