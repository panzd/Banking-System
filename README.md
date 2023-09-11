# Banking-System By Golang

Golang中基于REST API，实现一个简易的银行系统，实现以下功能:

- 用户的增删改查
- 账户金额的存入与提出
- 基于角色的访问控制RBAC

##

运用到的组件:

- gorilla/mux
- gomock
- jwt-go
- 运用了领域设计模型

##

步骤概述

- [ ]  在Postman中得到8000端口发送的JSON消息
- [ ]  Go module 编译
- [ ]  创建请求复用器
- [ ]  对客户Requst请求的条件进行约束
- [ ]  对请求方法进行约束
- [ ]  部署Service服务端口和repository存储
- [ ]  尝试返回硬编码数据
- [ ]  Mock Adapter -> Database Adapter 正式连接到MySQL
- [ ]  创建ID接口
- [ ]  创建错误信息返回函数
- [ ]  创建写响应函数
- [ ]  重构JSON解码函数
- [ ]  设计应用程序错误Exception
- [ ]  根据状态查询用户
- [ ]  引入结构化日志Logger
- [ ]  添加DTO层
- [ ]  添加本地化变量
- [ ]  创建用户账户
- [ ]  验证账户
- [ ]  进行系统测试
