## chatim
https://gitee.com/kiritokun/chatim

这是一个go-zero开发的聚合聊天项目

- platform  模拟平台ws服务端

  platform里需要有的 
  
  websocket服务端，不断往客户端发消息，而且会调http接口发消息

  `ws://127.0.0.1:8888/platform/ws/mtflower?token=111&platformType=1`
- rmq 连接平台ws的客户端，是rocketmq的（下行消息）生产者和（上行消息）消费者

  rmq里需要有的

  接口接收平台的http回调消息

  rocketmq生产者

  rocketmq消费者

  在消费者里需要存表

- chatim 连接web前端的服务端，是rocketmq的（上行消息）生产者和（下行消息）消费者

  `ws://localhost:8890/chatim/ws?token=123`

清单

-[x] 1.搭建platform，用于平台向rmq发消息
-[x] 2.搭建rmq，收到消息后发到rocketmq
-[ ] 3.搭建api，从rocketmq读消息
-[ ] 4.api读到消息后向客户端发送

platform
```shell
goctl api go -api .\service\platform\platform.api -dir .\service\platform
```

rmq
```shell
goctl api go -api .\service\rmq\rmq.api -dir .\service\rmq
```

chatim
```shell
goctl api go -api .\service\chatim\chatim.api -dir .\service\chatim
```