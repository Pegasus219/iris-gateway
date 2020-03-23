### 一个用iris框架搭建的简单http网关入口

#### 环境变量
| 变量名 | 描述 | 默认值 |
| ----- | ---- | ---- |
| IRIS_HTTP_PORT | 端口 | 80 |
| IRIS_DEBUG | 是否开启debug日志 | 0 |
| IRIS_LOG_PATH | 本地日志文件地址 |  |

#### demo中可测试的请求地址
    GET /index?param= 测试get请求
    POST /index  测试post请求
    GET /panic  测试panic情况下的服务recover