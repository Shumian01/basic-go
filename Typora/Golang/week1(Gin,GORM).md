# Gin入门

![image-20250702175427781](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702175427781.png)

## 最简Gin应用

![image-20250702175649038](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702175649038.png)

监听成功

![image-20250702180744866](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702180744866.png)

## gin.Engine

![image-20250702181034135](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702181034135.png)



## gin.Context

![image-20250702184502991](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702184502991.png)



## http方法包括:

![image-20250702184717775](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702184717775.png)

## 路由注册

![image-20250702185124360](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702185124360.png)



## Gin路由

Gin支持很多路由:

* 静态路由:完全匹配的路由,也就是前面我们注册的hello路由
* 参数路由:在路径上带上了参数的路由
* 通配符路由:任意匹配的路由
* ![image-20250702185441409](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702185441409.png)



## Gin路由:参数路由

参数路由的关键是获得参数,Gin提供了Param方法

## Gin路由:查询参数

查询参数和参数路由不太一样,查询参数是指在URL后面附着的参数,要用Query方法来拿到查询参数的值

![image-20250702193930420](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702193930420.png)



## Gin路由:通配符路由

![image-20250702194738919](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702194738919.png)



## Gin其他输入输出方法

 ![image-20250702194838145](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702194838145.png)



## 我改用什么方法？什么路由

![image-20250702194944827](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702194944827.png)



# 定义用户基本接口

![image-20250702195554326](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702195554326.png)



## Handler的用途

![image-20250702195826257](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702195826257.png)



## 集中注册vs分散注册

![image-20250702205138297](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702205138297.png)



## 用分组路由来简化注册

![image-20250702205712429](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702205712429.png)



## 目录结构

![image-20250702211755553](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250702211755553.png)



## 注册页面

![image-20250711155802458](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250711155802458.png)



## 接受请求数据

![image-20250711185344266](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250711185344266.png)



## 接受请求数据:Bing 方法

![image-20250711185511578](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250711185511578.png)



## 正则表达式校验

![image-20250714160658385](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250714160658385.png)



regxp.Match()



# 跨域问题

![image-20250715020659432](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715020659432.png)

![image-20250715021408594](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715021408594.png)

## preflight请求的特征

![image-20250715021459172](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715021459172.png)



## 使用middleware来解决CORS

![image-20250715022017843](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715022017843.png)



## middleware是啥

![image-20250715163420690](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715163420690.png)



## middleware在Gin中的定义和用法

![image-20250715163852594](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715163852594.png)



## middle例子

![image-20250715164138978](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250715164138978.png)



## CORs middleware

![image-20250718172148031](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718172148031.png)



## CORS middleware效果

![image-20250718172447582](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718172447582.png)



## 跨域问题要点

![image-20250718172627887](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718172627887.png)



## 设计并实现一个Gin插件库

![image-20250718172914541](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718172914541.png)	



# Gin面试要点

## Gin面试题

* 什么是Gin的middleware?能用来解决什么问题
* 什么是跨域问题,怎么解决？
* 跨域问题需要设置哪些头部?





![image-20250718174046978](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718174046978.png)

# 用户的基本功能与GORM入门



## GORM入门

![image-20250718174124314](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718174124314.png)



## GORM入门:增删改查

![image-20250718174446922](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718174446922.png)



## GORM学习难点

![image-20250718181540512](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718181540512.png)



## Product定义

![image-20250718181818955](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718181818955.png)



## 模型定义





## 用户注册:存储用户基本信息

![image-20250718182658182](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250718182658182.png)



## Docker Compose基本命令

![image-20250719163510986](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719163510986.png)



## 数据库相关代码放哪里？

![image-20250719163751015](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719163751015.png)



## 引入Service-Repostiory-DAO 三层结构

![image-20250719163909399](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719163909399.png)



## 如何理解

![image-20250719171532940](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719171532940.png)



## 调用流程

![image-20250719201337977](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719201337977.png)



## 改造代码

![image-20250719201434274](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250719201434274.png)



## dao中的User模型

![image-20250720162320745](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720162320745.png)



## User模型

![image-20250720162659445](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720162659445.png)



## 如何建表?

![image-20250720162912976](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720162912976.png)



## 初始化结构体

![image-20250720183737075](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720183737075.png)



## main函数

![image-20250720183758457](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720183758457.png)



# 密码加密



## 密码怎么加密?

![image-20250720185356847](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720185356847.png)



##  加密的位置

![image-20250720185729160](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250720185729160.png)



## 如何加密

![image-20250722000717416](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722000717416.png)



## 使用BCrypt加密

![image-20250722001054273](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722001054273.png)



## 如何获得邮件冲突的错误

![image-20250722003306767](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722003306767.png)



## 传导错误与检测

![image-20250722010509884](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722010509884.png)



# 登录功能

![image-20250722142407216](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722142407216.png)



## 登录接口实现

![image-20250722144537429](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722144537429.png)



## 登录校验

![image-20250722155220965](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722155220965.png)



## 无状态的HTTP协议

![image-20250722155254181](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722155254181.png)



## Cookie

![image-20250722155535289](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722155535289.png)

![image-20250722160027152](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722160027152.png)



## *Session*

![image-20250722160623849](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722160623849.png)	

![image-20250722160856112](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722160856112.png)



## 如何让客户端携带 sess_id

![image-20250722161017449](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722161017449.png)



## 几个网站的Cookie

![image-20250722161305009](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722161305009.png)



## 使用Gin的Session插件来实现登录功能

![image-20250722161345342](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722161345342.png)



## 登录校验实现

![image-20250722171935796](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250722171935796.png)





## 登录实现顺序

1. 先在main函数接入![image-20250723145813857](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723145813857.png)



2. 在web里面设置,不要忘记save

	![image-20250723145844050](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723145844050.png)

3. 注册middleware

	![image-20250723145927962](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723145927962.png)

![image-20250723145940242](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723145940242.png)

![image-20250723160606502](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723160606502.png)



## 面试要点

![image-20250723160623114](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250723160623114.png)



## 增强扩展GORM功能

![image-20250724150536604](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250724150536604.png) 



## 作业

![image-20250724150836496](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250724150836496.png)



