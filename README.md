# 简介

欢迎使用wechat_message_push服务

# 目录介绍

````
├── common      --服务初始化的具体的配置文件信息 
├── controller  --服务接口的控制层
├── job         --服务的定时器等任务存放处
├── model       --服务初始化参数信息和数据模型
├── router      --服务接口路由
├── service     --接口的服务层，接口的主要逻辑就写在这里
├── tools       --公共的通用方法
├── main.go     --服务的启动文件
````

# 依赖环境

1. Go 1.7 版本及以上（如使用 go mod 需要 Go 1.17）

# Git提交规范

1. feat：新增功能（feature）
2. fix：修复补丁（bug）
3. docs：修订文档，如 Readme, Change Log, Contribute 等
4. refactor：代码重构，未新增任何功能和修复任何 bug
5. style： 仅调整空格、格式缩进等（不改变代码逻辑的变动）
6. perf：优化相关，改善性能和体验的修改
7. test：测试用例的增加/修改
8. chore：非 src 和 test 的修改
9. merge：合并分支或冲突等
10. revert： 回滚到上一个版本
11. build：改变构建流程，新增依赖库、工具等（例如 webpack、maven 修改）
12. ci：自动化流程配置修改

# 快速开始
运行项目根目录下main.go的main函数（或是在项目根目录处使用命令go run main.go运行）

