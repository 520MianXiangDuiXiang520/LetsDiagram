# LetsDiagram
基于 Web 的多人协同实时绘图系统

## How To Run

### Linux

请确保您的计算机安装了以下应用：
1. git
2. docker (版本大于 17)
3. docker-compose (如果没有可以手动编译 dockerfile)


克隆本项目

```cmd
git clone git@github.com:520MianXiangDuiXiang520/LetsDiagram.git
```

修改配置文件：
* 后端： 根据您的需要修改 `let-diagram-api/setting.json` 配置 MySQL，CORS 等信息
* 前端：
  * 在 `lets-diagram/main.js` 中修改后端服务地址
  * 修改 `lets-diagram/nginx/conf` 中的 nginx 配置文件

构建镜像并运行项目

```cmd
docker-compose up -d
```

### Windows

windows 下您可以分别运行前后端项目，这需要 `go` 和 `node` 环境

```sh
# 后端
cd let-diagram-api
go mod download
go run .
```

```sh
# 前端
cd lets-diagram
npm install
npm run serve
```

> 记得修改配置文件