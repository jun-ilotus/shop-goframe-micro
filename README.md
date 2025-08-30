## 项目介绍

- 项目名称：GoFrame微服务电商平台（合内容社区）
- 核心定位：基于GoFrame最新微服务架构(GoFramev2.9+)，实现高并发、易扩展的全栈电商系统，覆盖多端协同(H5用户端、Admin管理端、微服务后端)。

## 项目功能梳理
微服务拆分（按业务域）

| 服务名          | 包含模块                         | 通信协议 |
| --------------- | -------------------------------- | -------- |
| user-service    | 用户管理、登录注册、收货地址     | gRPC     |
| product-service | 商品管理、分类、收藏、评论、点赞 | gRPC     |
| order-service   | 购物车、订单、退款、优惠券       | gRPC     |
| content-service | 文章管理、轮播图、手工位图       | gRPC     |
| admin-service   | 角色权限、城市地址、数据统计     | gRPC     |

技术栈

| 层级       | 技术选型                                        |
| ---------- | ----------------------------------------------- |
| 基础设施   | Docker+Kubernetes(本地开发用docker-compose)     |
| 微服务框架 | GoFramev2.9+(内置服务注册/配置中心/链路追踪)    |
| 通信       | gRPC+Protocol Buffers (高性能RPC)               |
| 网关       | Kong(API网关，路由转发+限流)                    |
| 存储       | MySQL(TiDB备用)+Redis+Elasticsearch（评论检索） |
| 前端       | Vue3+TS (Admin) + Uniapp (H5)                   |
| DevOps     | Jenkins+GitLab CI （自动化部署）                |

