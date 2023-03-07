## ddg-api的docker镜像
一个DuckDuckGo搜索api的docker镜像


### 构建api正式包
在项目根路经执行
```shell
docker build . -t samge/ddg-api-free -f docker/Dockerfile
```

### 上传
```shell
docker push samge/ddg-api-free
```

### 运行docker镜像

如果需要挂载config.json跟run.log文件，在运行docker之前需要先创建文件，否则默认挂载路径会生成文件夹导致挂载失败。

`/home/samge`需要替换为你自己的本地路径

在项目根路径执行：
```shell
mkdir -p /home/samge/docker_data/ddg-api-free
cp config.dev.json /home/samge/docker_data/ddg-api-free/config.json
touch /home/samge/docker_data/ddg-api-free/run.log
```


`第一种：基于环境变量运行，不是部署在国外服务器需要设置PROXY代理`
```shell
docker run -d \
--name ddg-api-free \
-p 8231:8080 \
-v /home/samge/docker_data/ddg-api-free/run.log:/app/run.log \
-e ACCESS_TOKEN=换成你接口请求的token（请求头中的xxx值，Authorization: Bearer xxx） \
-e PROXY="http://127.0.0.1:7890" \
-e ALLOW_ORIGIN=* \
--pull=always \
--restart always \
--memory=0.5G \
samge/ddg-api-free:latest
```

`第二种：基于配置文件挂载运行`
```shell
docker run -d \
--name ddg-api-free \
-v /home/samge/docker_data/ddg-api-free/config.json:/app/config.json \
-v /home/samge/docker_data/ddg-api-free/run.log:/app/run.log \
-p 8231:8080 \
--pull=always \
--restart always \
--memory=0.5G \
samge/ddg-api-free:latest
```