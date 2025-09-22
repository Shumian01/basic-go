# Kubernetes入门

![image-20250806015453739](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806015453739.png)



## 入门:基本概念

![image-20250806015859284](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806015859284.png)



## Docker启用K8s

![image-20250806020840119](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806020840119.png)



## 安装kubectl工具

![image-20250806020909128](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806020909128.png)

# 部署k8s





## 先去除对mysql和redis的依赖

![image-20250806022328304](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806022328304.png)



## 部署方案

![image-20250806022340378](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806022340378.png)



## 准备K8s容器镜像

![image-20250806022541398](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806022541398.png)

![image-20250806022634644](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250806022634644.png)

## 编写Deployment

![image-20250807025503588](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807025503588.png)



命令：kubectl apply -f k8s-webook-deployment.yaml

​		kubectl get deployments



![image-20250807025917580](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807025917580.png)



## apiVersion

![image-20250807030226852](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807030226852.png)



## spec

![image-20250807030504980](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807030504980.png)



## selector

![image-20250807030755161](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807030755161.png)

## template

![image-20250807031032587](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807031032587.png)



## image

![image-20250807031224332](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250807031224332.png)



## 编写Service

![image-20250808133941580](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808133941580.png)



## 启动服务

![image-20250808134730417](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808134730417.png)





# K8s部署Mysql

![image-20250808155859192](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808155859192.png)



## Mysql Service Deployment

![image-20250808160512924](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808160512924.png)



![image-20250808161747489](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808161747489.png)



## mysql template

![image-20250808164109680](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808164109680.png)

## pvc

![image-20250808165027951](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808165027951.png)



## PersistmentVolume

![image-20250808195251915](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808195251915.png)



## storageClass

![image-20250808204708411](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808204708411.png)



## accessMode

![image-20250808204736200](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808204736200.png)

![image-20250808210012691](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808210012691.png)



## 管理工具直连Mysql

![image-20250808211153466](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808211153466.png)



# K8s部署Redis

## 部署最简单的单机Redis

![image-20250808220102982](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808220102982.png)



## 允许外部访问Redis

![image-20250808220116115](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808220116115.png)



## port  nodeport tagetport 含义

![image-20250808220237298](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808220237298.png)



## 测试Redis

![image-20250808221010909](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808221010909.png)

```
.\redis-cli -h localhost -p 30003
```



# K8s部署nginx



## 什么是ingress

![image-20250808221159236](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808221159236.png)



## ingress 和ingress controller

![image-20250808221506798](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250808221506798.png)



## 安装helm, ingress-nignx

![image-20250809004232112](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250809004232112.png)



## 编写ingress

![image-20250809005810280](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250809005810280.png)



# 编译标签

![image-20250811021001232](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811021001232.png)



![image-20250811021014076](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811021014076.png)



# 前端

![image-20250811021024779](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811021024779.png)



![image-20250811021144123](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811021144123.png)



# 面试要点

![image-20250811021154898](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250811021154898.png)