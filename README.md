项目结构

```dockerfile
.
├── app
│   ├── common              # 公共模块（请求、响应结构体等
│   ├── controllers         # 控制器
│   ├── dao                 # curd
│   ├── middlewares         # 中间件
│   ├── models              # db模型
│   └── services            # 逻辑层
├── bootstrap                     # 初始化
├── config                        # 配置
├── dto
├── global                        # 全局变量
├── logs                          # 日志
├── pkg
│   └── response
├── routes
├── static                        # 静态资源
├── storage                       # 存储
└── utils
```