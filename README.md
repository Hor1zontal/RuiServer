# Rui
```
├── config  //加载配置文件
├── db      //数据库层（直接与数据库连接并进行增删改查）
├── exception   //异常处理（直接将错误抛给客户端）
├── initialize  //初始化
├── router      //路由层（调用服务层，与前端进行数据的交互，处理数据）
│─────├── handler //处理请求的方法
│─────├── helper   //帮助函数（用来方便处理发送给客户端的数据）
├── service //服务层（做逻辑处理，并调用数据库层）
└── utils（通用方法）
```

前端请求过来
router层--->service层--->db层