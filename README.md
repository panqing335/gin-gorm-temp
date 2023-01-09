#### 目录结构类似Hyperf框架，封装了一些工具类，使代码写起来更流畅，结构更清晰
#### 非常适合转golang的phper

### 技术框架
```
1. Web框架gin
2. ORM框架grom
3. 配置管理viper
4. 数据库Mysql
5. redis
6. logrus
7. crom
8. grpc
9. ..
```

### 目录结构
```
├── README.md
├── app                  
│   └── common
│   └── constants
│   └── controller                      controller类，对外的接口。所有http请求的入口
│   └── entity                          请求实体类
│       └── dto                         变更类
│       └── qo                          查询类
│   └── grpcService                     grpc生成的 pb.go 目录
│   └── middleware                      中间件
│       └── authMiddleware.go           jwt验证中间件
│       └── recoverMiddleware.go        全局错误捕获中间件
│       └── corsMiddleware.go           cors中间件
│   └── model                           模型类目录
│   └── router                          路由
│   └── service                         业务逻辑目录
│   └── task                            cron目录
│   └── utils                           工具类目录
├── bin
│   └── main.go                         main函数入口，程序启动
├── config
│   └── config.go   
│   └── config.yml                      全局配置文件   
├── protos                              proto目录
├── go.mod
│   └── go.sum                        
├── grpcGen.sh
└── runtime
   └── logs                             日志文件

```

### 入口文件
```
temp/bin/main.go

func main() {
	Config.Setup() // 初始化配置
	Redis.Setup()  // 初始化redis
	Mysql.Setup()  // 初始化DB
	Grpc.Setup()   // 初始化Grpc
	Cron.Setup()   // 初始化定时任务
	Router.Setup() // 路由，开启服务
}
```

### 配置文件
```
temp/config/config.yml

server:
  port: 9001                // httpserver port

grpc:
  enabled: true             // 是否开启grpc
  port: 9002                // grpc端口

datasource:
  driverName: mysql         // mysql配置
  host: 121.43.42.249
  port: 3309
  database: temp
  username: root
  password: 123456
  charset: "utf8mb4"
  loc: Local
  tablePrefix: ""

redis:                      // redis配置
  host: 192.168.0.88:6379
  db: 0
  password: lyh.123456

log:                        // log配置
  filePath: ./runtime/logs/
  fileName: system.log

cron:               
  enabled: false            // 定时任务是否开启

modelCache:                 // 模型缓存配置
  key: '{mc:%s}:%s:%s'
  ttl: 1500
  emptyTTL: 15
```

### 工具类
```
app/utils/response.go                   // response封装
app/utils/paramsBindEntity/bind.go      // requset参数绑定
app/utils/serResult.go                  // 类似rust的异常处理，告别err != nil{..} 

使用示例：
app/controller/login/loginController:

func Login(c *gin.Context) {
	var loginDto dto.LoginDto 
	loginDto = paramsBindEntity.Bind(c, loginDto)
	resultToken := tempUserService.Login(loginDto)

	util.Success(c, resultToken)
}

app/service/tempUserService:

func Login(loginDto dto.LoginDto) ResultToken {
	tempUser := model.TempUser{Username: loginDto.Username}
	user := util.NewResult(tempUser.FindByUsername()).UnwrapOr(errorCode.LOGIN_ERROR, "根据用户名查询数据失败")
	util.NewResult(util.CheckPassword(loginDto.Password, user.PasswordSalt, user.Password)).UnwrapOr(errorCode.LOGIN_ERROR, "账户或密码错误，请重试")

	token, _ := util.GenerateToken(util.HmacUser{
		Id:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
	})

	return ResultToken{Token: token}
}

```
```
app/utils/session.go                    // 通过auth中间件获取jwt携带的值
app/utils/genModel.go                   // 生成model文件
app/utils/Iterator.go                   // 泛型迭代器
app/utils/paginator/bind.go             // 分页器
app/utils/logger.go                     // 日志封装
...
```

## 感 谢 Star！！！