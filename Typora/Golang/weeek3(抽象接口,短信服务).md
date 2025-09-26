

![image-20250829174046756](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829174046756.png)

# 优化登录性能



## 利用wrk来压测接口

![image-20250811140018652](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811140018652.png)



## 压测前准备

![image-20250828194655738](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250828194655738.png)

## 压测注册接口

命令

```
shumian@xzl:/mnt/f/GoCode/src/basic-go/webook$ wrk -t1 -c2 -d1s -s ./scripts/wrk/signup.lua --latency "http://$HOSTIP:8080/users/signup"

//login
shumian@xzl:/mnt/f/GoCode/src/basic-go/webook$ wrk -t4 -c20 -d10s -s ./scripts/wrk/login.lua --latency "http://$HOSTIP:8
080/users/login"

//profile
shumian@xzl:/mnt/f/GoCode/src/basic-go/webook$ wrk -t4 -c20 -d10s -s ./scripts/wrk/profile.lua --latency "http://$HOSTIP:8
080/users/profile"
```

![image-20250829153300539](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829153300539.png)

![image-20250829153419954](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829153419954.png)

注册接口性能瓶颈:加密，数据库操作

## 压测登录接口

![image-20250829173909209](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829173909209.png)

## 压测Profile接口

![image-20250829173814609](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829173814609.png)

## 扩展练习

![image-20250829173928417](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829173928417.png)



## 性能瓶颈

![image-20250829174153660](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829174153660.png)

## 引入缓存

![image-20250829174335983](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250829174335983.png)

## 业务专属缓存抽象的作用

![image-20250830182338196](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830182338196.png)



## 序列化与反序列化

![image-20250830182416243](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830182416243.png)



## 集成UserCache

![image-20250830182431907](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830182431907.png)



## 代码详解

![image-20250830182448150](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830182448150.png)

## 检测数据不存在的写法

![image-20250830184803811](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830184803811.png)



## 登录要不要利用Redis优化性能

![image-20250830185057843](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830185057843.png)



## Redis数据结构

![image-20250830185258493](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830185258493.png)

## 升职加薪指南

![image-20250830190724194](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830190724194.png)

## 如何在测试中从维护登录态

![image-20250830190914078](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830190914078.png)



## Redis面试题目

![image-20250830191040157](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830191040157.png)



# 短信验证码登录

## 多种登录方式

![image-20250830204334243](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830204334243.png)



## 需求分析与系统设计

![image-20250830210527374](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830210527374.png)

![image-20250830210547389](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830210547389.png)

![image-20250830211207368](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830211207368.png)

## 从功能 非功能角度分析

![image-20250830211444283](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830211444283.png)



## 发送验证码流程

![image-20250830211523835](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830211523835.png)



## 登录流程

![image-20250830211551608](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830211551608.png)



## 深入分析

![image-20250830211745842](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830211745842.png)



## 手机验证码功能

![image-20250830212013580](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830212013580.png)



## 综合考虑

![image-20250830212144207](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830212144207.png)



## 服务划分

![image-20250830213331997](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830213331997.png)



# 短信服务

![image-20250830214411659](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830214411659.png)



## 以特定供应商的短信API开始

![image-20250830214500802](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830214500802.png)



## 腾讯短信API

![image-20250830215124872](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830215124872.png)

## 接口抽象



![image-20250830220413189](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830220413189.png)



## sms包

![image-20250830220619213](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830220619213.png)



## 腾讯实现:定义与初始化



![image-20250830221028858](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250830221028858.png)

## 腾讯实现:发送实现

![image-20250831134544903](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831134544903.png)



## 腾讯实现:运行测试

![image-20250831142316272](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831142316272.png)



# 验证码服务

![image-20250831171204734](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831171204734.png)

## 深入分析验证码安全问题

![image-20250831171247460](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831171247460.png)



## 验证码服务抽象接口

![image-20250831172652424](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831172652424.png)



## 发送验证码

![image-20250831173732293](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250831173732293.png)



## 怎么实现？

### 下面伪代码有什么问题？

![image-20250908184651524](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250908184651524.png)



### Check-do-something(并发场景)

![image-20250908185042371](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250908185042371.png)



### 并发场景分析

![image-20250908185127017](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250908185127017.png)



### 怎么办？

![image-20250908185208117](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250908185208117.png)



### lua脚本整体实现

![image-20250908195428080](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250908195428080.png)



## SendCode实现

![image-20250909114307794](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250909114307794.png)



## 验证验证码

![image-20250909114418970](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250909114418970.png)



## Verify实现

![image-20250909184023421](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250909184023421.png)

## 深入讨论:业务逻辑放在哪里合适?

![image-20250909184519286](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250909184519286.png)



## 深入讨论:验证码发送渠道抽象

![image-20250909193145556](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250909193145556.png)



# 用户验证码登录

## 验证码登录接口

![image-20250910101100732](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250910101100732.png)



## 另外一种思路

![image-20250915105142116](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250915105142116.png)	

## 发送验证码核心逻辑

![image-20250915105314745](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250915105314745.png) 



## 验证码登录逻辑

![image-20250917084603613](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250917084603613.png)

 
## FindOrCreate 并发问题
![[Pasted image 20250917100526.png]]


## FindOrCreate逻辑
![[Pasted image 20250917101450.png]]
## 唯一索引问题
![[Pasted image 20250917101819.png]]
## 为什么在UserHandler聚合
![[Pasted image 20250917101936.png]]

## 使用本地sms实习短信登录服务
![[Pasted image 20250917102201.png]]

## 提供sql包扩展支持
![[Pasted image 20250917102702.png]]

## 如何实现一个验证码登录功能
![[Pasted image 20250917102940.png]]


# 面向接口编程 依赖注入
![[Pasted image 20250917103637.png]]

## 依赖注入
![[Pasted image 20250917112307.png]]

##  两种写法对比
![[Pasted image 20250917112321.png]]
## 引入依赖注入中间件
![[Pasted image 20250917112600.png]]
## wire简单示例
![[Pasted image 20250917161108.png]]
## wire生成代码
![[Pasted image 20250917161150.png]]

## ioc控制反转
![[Pasted image 20250917161238.png]]

# wire改造代码

## 组装过程
![[Pasted image 20250917162403.png]]

## 预期main函数
![[Pasted image 20250917190926.png]]

## 组装Repository
![[Pasted image 20250917190948.png]]

## 组装Headler
![[Pasted image 20250917191023.png]]

## InitSmsService
![[Pasted image 20250917191048.png]]

## 组成gin.Engine
![[Pasted image 20250917191156.png]]
## GinMiddlewares

![[Pasted image 20250917191539.png]]

## 缺点那么多，为啥用wire
![[Pasted image 20250917191606.png]]

## 已有代码缺点

![[Pasted image 20250917191810.png]]

# 面向接口编程
## 什么是面向接口编程 
![[Pasted image 20250917191947.png]]

## 非面向接口编程
![[Pasted image 20250917193445.png]]
## 面向接口
![[Pasted image 20250917193652.png]]
## 为什么要面向接口编程
![[Pasted image 20250917193726.png]]
## 改造代码
![[Pasted image 20250917194220.png]]


![[Pasted image 20250917211645.png]]
![[Pasted image 20250917211816.png]]![[Pasted image 20250917211819.png]]![[Pasted image 20250917211856.png]]
![[Pasted image 20250918142247.png]]
# 超前设计 最小化实现
![[Pasted image 20250918142451.png]]
![[Pasted image 20250918143046.png]]


# 单元测试

![[Pasted image 20250918204445.png]]
## 已有代码测试困难
![[Pasted image 20250918204634.png]]


## 怎么办？
![[Pasted image 20250918204723.png]]

## 单元测试
![[Pasted image 20250918205750.png]]

## Go中编写单元测试
![[Pasted image 20250918205956.png]]

## IDE直接允许单元测试
![[Pasted image 20250918210429.png]]
## IDE直接允许包下面的全部测试![[Pasted image 20250918210457.png]]

## IDE运行结果
![[Pasted image 20250918210806.png]]

## 命令行运行单元测试
![[Pasted image 20250918210933.png]]

## Table Driven模式
![[Pasted image 20250920135447.png]]

## 运行 Table Driven下的单个测试
![[Pasted image 20250920141248.png]]

## 测试用例定义
![[Pasted image 20250920141321.png]]

## 运行测试用例
![[Pasted image 20250920143052.png]]

## 设计测试用例
![[Pasted image 20250920143318.png]]

## 不是所有场景都很好测试
![[Pasted image 20250920143724.png]]


## 测试Headler
![[Pasted image 20250920144611.png]]
### http接口测试
![[Pasted image 20250920152803.png]]

#### 构造http请求
![[Pasted image 20250920152840.png]]
#### 获得http响应
![[Pasted image 20250920153450.png]]
#### 用httpest来记录响应
![[Pasted image 20250920153734.png]]
#### 怎么解决UserHandler的初始化问题
![[Pasted image 20250920154439.png]]

#### mock工具入门
![[Pasted image 20250920154527.png]]
#### 为UserService和CodeService生成mock实现

![[Pasted image 20250920162044.png]]

```
mockgen -source=./webook/internal/service/user.go -package=svcmocks -destination=./webook/internal/service/mocks/user.mock.go

```

#### 使用mock
![[Pasted image 20250922103957.png]]
#### 测试
![[Pasted image 20250922105523.png]]

####  设计测试用例:注册成功

![[Pasted image 20250922142107.png]]

#### 设计测试用例:注册成功 mock
![[Pasted image 20250922142155.png]]

#### 设计测试用例:注册成功的http请求与响应
![[Pasted image 20250922142300.png]]

#### 测试数据校验逻辑

![[Pasted image 20250922142352.png]]

#### 测试bind方法出错的用例
![[Pasted image 20250922142452.png]]
#### 步骤总结
![[Pasted image 20250922143039.png]]
### 测试service

#### 测试login
![[Pasted image 20250922162244.png]]

![[Pasted image 20250922163228.png]]
![[Pasted image 20250922163241.png]]

![[Pasted image 20250922163254.png]]

### 测试Repository

### 测试cache
![[Pasted image 20250923140911.png]]
### sql mock入门
![[Pasted image 20250923172446.png]]
### 测试DAO
![[Pasted image 20250923171305.png]]

#### 定义测试用例
![[Pasted image 20250923172626.png]]
####  运行测试代码
![[Pasted image 20250923173613.png]]
#### 插入成功用例
![[Pasted image 20250923173641.png]]

![[Pasted image 20250923220840.png]]

# 集成测试
![[Pasted image 20250923234737.png]]
![[Pasted image 20250923235045.png]]

## 集成测试:以发验证码为例
![[Pasted image 20250923235452.png]]

### 定义测试用例
![[Pasted image 20250924233253.png]]

### 运行测试用例
![[Pasted image 20250924233311.png]]
### 总结
![[Pasted image 20250925000238.png]]

# 第三方服务调用治理
![[Pasted image 20250925081311.png]]
## 整体思路
![[Pasted image 20250925081603.png]]
## 客户端限流
![[Pasted image 20250925082119.png]]
##  第一种做法:整体限流
![[Pasted image 20250925083253.png]]
## 限流器抽象
![[Pasted image 20250925085304.png]]
## 在已有代码里面集成限流器
![[Pasted image 20250925101845.png]]
## 进一步改进
![[Pasted image 20250925103547.png]]

## 利用装饰器来改进
![[Pasted image 20250925103727.png]]## 如何理解装饰器模式
![[Pasted image 20250925103925.png]]
## 装饰器模式实现限流的短信服务
![[Pasted image 20250925104050.png]]

## 单元测试装饰器
![[Pasted image 20250925105026.png]]
## 装饰器另外一种实现
![[Pasted image 20250925111616.png]]
## 开闭原则 非侵入式 装饰器
![[Pasted image 20250926103704.png]]

## 自动切换不同服务商
### 服务商出问题
![[Pasted image 20250926104831.png]]
### 怎么知道服务商出现问题
![[Pasted image 20250926104844.png]]
### 第一种策略failover
![[Pasted image 20250926105143.png]]
#### failover 第一种实现
![[Pasted image 20250926105742.png]]
#### failover 第一种实现的测试
![[Pasted image 20250926110215.png]]
![[Pasted image 20250926110239.png]]

#### failover第二种实现

![[Pasted image 20250926110628.png]]
![[Pasted image 20250926111634.png]]
![[Pasted image 20250926111739.png]]
### 第二种 策略
![[Pasted image 20250926112346.png]]
#### 基于超时响应的判定
![[Pasted image 20250926115421.png]]
#### 具体实现
![[Pasted image 20250926115504.png]]
#### 单元测试
![[Pasted image 20250926121742.png]]
![[Pasted image 20250926121834.png]]