## chatim
https://gitee.com/kiritokun/chatim

这是一个go-zero开发的聚合聊天项目

- platform  模拟平台ws服务端
- rmq 连接平台ws的客户端，是rocketmq的（下行消息）生产者和（上行消息）消费者
- chatim 连接web前端的服务端，是rocketmq的（上行消息）生产者和（下行消息）消费者

platform
```shell
goctl api go -api .\service\platform\platform.api -dir .\service\platform\
```