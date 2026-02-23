![log](./docs/image/main.png)       

### 🛠️技术栈

- **后端框架**：Kitex 、Hertz 
- **消息队列**：Kafka、RabbitMQ
- **安全鉴权**：JWT
- **配置管理**：Viper
- **日志服务**:  Elasticsearch（同时用于搜索功能的实现）、Filebeat、Kibana、Zap
- **数据库**：MySQL、Redis、GORM
- **可观测性**：Prometheus、Grafana、Jaeger、OpenTelemetry
- **对象储存**：腾讯云COS
- **服务注册与发现**：Etcd
- **限流**：sentinel-golang
- **other**：lua、布隆过滤器、docker-compose、dockerfile、singleflight、Snowflake


### 🚀 快速开始
```
1、克隆本项目
2、填写config.yaml文件
3、docker-compose up -d
4、根据config/sql文件创建表格
5、make start_all
```


### ✨优点

- 使用布隆过滤器解决缓冲穿透、使用singleflight解决缓存击穿、使用添加随机偏移量以及限流解决缓存雪崩
- 采用本地消息表模式实现分布式事务的最终一致性
- OpenTelemetry + EFK 实现可观测性和日志搜索, 减少问题排查难度
- 使用redis lua脚本防止超卖和高并发查写一致
- 使用kafka解决大数据高吞吐的数据流、使用rabbitMq解决例如订单等延时任务
- 采用整洁架构对服务进行深度解耦
- 使用雪花算法实现分布式唯一ID
- 使用容器化部署，方便快捷


      

### ⛏️架构图

![architecture](./docs/image/architecture.png)

### 📖接口文档

[接口文档](https://4721v9dymm.apifox.cn)


            

### 📌单服务架构（整洁架构）

使用接口使依赖倒置， 耦合度低（如更换消息队列只需重写infrastructure里的方法）

```
app
└─  coupon
   ├─ controller
   │  └─ api
   │     ├─ handler.go
   │     └─ pack
   ├─ domain
   │  ├─ model
   │  ├─ repository
   │  └─ service
   ├─ infrastructure
   │  ├─ cache
   │  ├─ mq
   │  └─ mysql
   ├─ inject.go
   └─ usecase
```


