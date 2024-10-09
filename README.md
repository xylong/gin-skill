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
包
- 大小写转换：https://github.com/iancoleman/strcase
- 英文单复数：https://github.com/gertd/go-pluralize