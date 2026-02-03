![log](./docs/image/main.png)       

### 🛠️技术栈

- **后端框架**：Kitex 、Hertz 
- **消息队列**：Kafka、RabbitMQ
- **配置管理**：Viper
- **日志服务**:  Elasticsearch（同时用于搜索功能的实现）、Filebeat、Kibana、Zap
- **数据库**：MySQL、Redis、GORM
- **可观测性**：Prometheus、Grafana、Jaeger、OpenTelemetry
- **对象储存**：腾讯云COS
- **服务注册与发现**：Etcd
- **限流**：sentinel-golang
- **other**：lua、布隆过滤器、docker-compose

### ⛏️架构图

![architecture](./docs/image/architecture.png)

### 📖接口文档

[接口文档](https://4721v9dymm.apifox.cn)



### 📌单服务框架（整洁架构）

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

