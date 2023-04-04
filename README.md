# MyCloudDisk

我的 go 练手项目

感谢项目：[xiaogao67/gin-cloud-storage: 使用Go语言开发的云存储网盘项目 (github.com)](https://github.com/xiaogao67/gin-cloud-storage)

在他的基础上想要修改为前后端分离的项目，奈何我并不会 js，所以进度缓慢



## 项目结构

- api/v1：后端接口，包括 login/logout/index 
- config：配置文件，使用 viper 解析 yaml ，初始化 Redis/MySQL 等
- controller：前端路由信息，因为之前的前端用的是 `c.HTML()` 模板渲染，想要改成前后端分离。改好后会直接用静态文件，删除这个目录
- global：全局变量，包含 Mysql/Redis/JWT/dir 等变量信息
- middleware：中间件，含有 jwt filter/ content-type filter 
- model：模型，包括 api 响应模型和 db 数据库模型
- router：gin 绑定路由和函数
- service：和数据库交互部分
- static：css/js 等文件
- storage：初始化方法
- tmp：临时文件，这里用于作为文件上传的路径
- utils：一些工具类，经常使用或者不需要知道具体实现的方法
- view：前端html



## 技术栈

- mysql
- redis（用于jwt）
- viper
- jwt



慢慢来写，细细地学
