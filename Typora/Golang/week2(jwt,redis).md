#  Session和JWT



## 主要内容

* 多实例部署的Session问题
* 刷新Session的过期时间
* JWT
* 初步保护系统





## 已有的实现

![image-20250725155251880](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725155251880.png)



## Gin Session 存储的实现

![image-20250725155700128](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725155700128.png)





![image-20250725155921819](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725155921819.png)



![image-20250725172254269](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725172254269.png)



## 启动Redis

![image-20250725172504079](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725172504079.png)



## 使用基于Redis的实现

![image-20250725201250303](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725201250303.png)



## 自由切换的好处

![image-20250725202335455](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725202335455.png)



# Session参数与刷新

## Gin Session参数

![image-20250725204017641](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250725204017641.png)



## 刷新登录状态

![image-20250726152729671](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250726152729671.png)



## 如何刷新

![image-20250727151153534](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250727151153534.png)



## 在Middleware中刷新

![image-20250727151514513](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250727151514513.png)



# JWT简介

![image-20250727160636538](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250727160636538.png)



![image-20250727161252162](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250727161252162.png)



## JWT使用

![image-20250727161514606](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250727161514606.png)



## 接入JWT的步骤总结

![image-20250802154824878](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250802154824878.png)



## 前端携带JWT

![image-20250802154846095](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250802154846095.png)



## JWT登录校验

![image-20250803142328520](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803142328520.png)





## JWT优缺点

![image-20250802163415426](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250802163415426.png)



## 混用jwt session

![image-20250802164012389](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250802164012389.png)





 # 系统保护



## 系统漏洞

![image-20250803142358007](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803142358007.png)



## 怎么办? 限流

![image-20250803142520431](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803142520431.png)



## 限流

![image-20250803142548835](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803142548835.png)



## 限流对象

![image-20250803142908070](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803142908070.png)



## 限流阈值

![image-20250803145832650](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803145832650.png)



## 使用Gin的限流插件

![image-20250803160350325](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803160350325.png)



## 为啥用redis限流

![image-20250803161235112](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803161235112.png)



# 增强登录安全

## 安全问题

![image-20250803164558642](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250803164558642.png)



## 怎么解决?

![image-20250805022434863](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250805022434863.png)



## 登录的其他信息

![image-20250805022656968](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250805022656968.png)



## 利用User-Agent增强安全性

![image-20250805024629700](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250805024629700.png)



## 保护公司前端接口







# 面试要点

![image-20250805155213633](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250805155213633.png)
