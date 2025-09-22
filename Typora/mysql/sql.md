# 1. 初识Mysql

## 1.连接数据库

命令行连接

```
mysql -u root -p xzl201515   --连接数据库

--------------------------
所有的语句使用分号结尾
show databases; 查看所有的数据库
use school --切换	数据库  use 数据库名

show tables;--查看数据库所有的表
describe student;--显示数据库中所有的表的信息

create database westos;--创建一个数据库

exit --退出连接

# 
--单行注释(sql本来的注释)

/*
dawda
*/  多行注释
```



数据库xxx语言 CRUD 增删改查

DDL 定义

DML 操作

DQL 查询

DCL 控制



# 2.SQL

操作数据库>操作数据库中的表>操作数据库表的数据



===mysql关键字不区分大小写===

## 2.1操作数据库

1 创建数据库

```mysql
CREATE DATABASE [IF NOT EXISTS] westos
```

2 删除数据库

```mysql
DROP DATABASE IF EXISTS westos
```

3 使用数据库

```
USE `school`   //如果表面名或字段名是一个特殊字符 需要带  ``
```

4 查看数据库

```
SHOW DATABASES   --查看所有的数据库
```

## 2.2 数据库的列类型

> 数值

* tinyint 十分小的数据    一个字节
* smallint  较小了数据   两个字节
* int	标准的整数       四个字节
* bitint                               八个字节
* float                                四个字节
* double                            八个字节
* decimal   字符串形式的浮点数

> 字符串

* char   字符串固定大小 0~255
* varchar  可变字符串 0~65535   常用的string
* tinytext   微型文本   2^8-1
* text         文本串       2^16-1  保存大文本

> 时间日期

* date YYYY-MM-DD 日期格式
* time HH:mm:ss 时间格式
* datetime  YYYY-MM-DD   HH:mm:ss 最常用的时间格式
* timestamp  时间戳   1970.1.1到现在的毫米数

> null

* 没有值 未知
* ==注意 不要使用null进行运算 结果为null==



## 2.3数据库的字段属性(重点)

* Unsigned:
	* 无符号的整数
	* 声明了该列不能为负数
* zerofill
	* 0填充的
	* 不足的位数 使用0来填充
* 自增：
	* 通常理解为自增，自动在上一条的记录的基础上+1 默认
	* 通常用来设计唯一的主键 ~index 必须是整数类型
	* 可以自定义设计主键的初始值和步长
* 非空 Null not null
	* 假设设置为 not null 如果不给他赋值 就会报错
	* Null 如果不填写值 默认就是null
* 默认
	* 设置默认的值
	* sex 默认值为男,如果不指定该列的值 则为默认的

==拓展:==

每一个表都要有这5个字段

id 主键

`version` 乐观锁

is_delete 伪删除

gmt_create 创建时间

gmt_updata 修改时间

## 表操作

添加字段 

```
ALTER TABLE 表名 ADD 字段名 类型 注释 约束
```

修改字段

修改数据类型

```
ALTER TABLE 表名 MODIFY 字段名 新数据类型(长度)
```

修改字段名和字段类型

```
ALTER TABLE 表名 CHANGE 旧字段名 新字段名 类型 注释 
```



删除字段

```
ALTER TABLE 表名 DROP 字段名
```



修改表名

```
ALTER TABLE 表名 RENAME TO 新表名;
```



表删除

```
DROP TABLE [IF EXISTS] 表名;
```



## 1. DDL-数据库操作

![4f9e2bd6-d440-4b7f-a2e6-5882d0b481ae](image/4f9e2bd6-d440-4b7f-a2e6-5882d0b481ae.png)

## 2. DDL-表操作

![e39b3151-b0a9-4f0e-9e05-330b103544dc](image/e39b3151-b0a9-4f0e-9e05-330b103544dc.png)



## DML 数据库操作语言

![820a7d91-4716-4a47-806c-b090c4a048a8](image/820a7d91-4716-4a47-806c-b090c4a048a8.png)

### DML-添加数据

![image-20250820152041071](C:/Users/肖子龙/AppData/Roaming/Typora/typora-user-images/image-20250820152041071.png)

![67d1b2c5-b4ea-458e-9c1e-40723a50d5cb](image/67d1b2c5-b4ea-458e-9c1e-40723a50d5cb.png)



### DML-修改数据

![c390f86f-a9ba-4fc1-b029-d7c177d50633](image/c390f86f-a9ba-4fc1-b029-d7c177d50633.png)

### DML-删除数据

![87566475-3bf9-455b-92b9-b6081ba1f5dd](image/87566475-3bf9-455b-92b9-b6081ba1f5dd.png)



## DQL-数据库查询

### DQL-语法

![9d411dbe-68f5-459f-8f7e-5fb9d9d659f5](image/9d411dbe-68f5-459f-8f7e-5fb9d9d659f5.png)

### DQL-基本查询

![cfe1fc49-c053-45a1-a54b-adff136cfbd3](image/cfe1fc49-c053-45a1-a54b-adff136cfbd3.png)

### DQL-条件查询

![56eddea0-1716-42a9-9400-0d3601c7a08b](image/56eddea0-1716-42a9-9400-0d3601c7a08b.png)



### DQL-聚合函数

![dd8ea31d-6dbc-4bfd-9619-38fc934f4f62](image/dd8ea31d-6dbc-4bfd-9619-38fc934f4f62.png)

==null值不参加聚合函数计算==



### DQL-分组查询

![3b95f9cb-3cab-4545-ad69-4ed93735e36b](image/3b95f9cb-3cab-4545-ad69-4ed93735e36b.png)

### DQL-排序查询

![bb2050e4-563c-428b-95ef-29207865fa6e](image/bb2050e4-563c-428b-95ef-29207865fa6e.png)

### DQL-分页查询

![a12f0d97-0453-484e-ad4b-bf04df6726f2](image/a12f0d97-0453-484e-ad4b-bf04df6726f2.png)

### DQL-执行顺序

![e0ebc51d-4f23-45d9-a2f7-8d19bbd5b154](image/e0ebc51d-4f23-45d9-a2f7-8d19bbd5b154.png)



## DCL 数据控制语言

用来管理数据库用户，控制数据库访问权限

### DCL-管理用户

![e6681f8a-4083-4ec7-9c51-b80dcdd936a6](image/e6681f8a-4083-4ec7-9c51-b80dcdd936a6.png)

### DCL-权限控制

![de57d740-5f31-41dd-a513-cbdea73a8a28](image/de57d740-5f31-41dd-a513-cbdea73a8a28.png)

![0a8ee1f8-9b5e-4d2d-949e-d810ca6ad7e7](image/0a8ee1f8-9b5e-4d2d-949e-d810ca6ad7e7.png)

### DCL小结

![a362a103-842d-4d7f-b366-d1ff67fa2d21](image/a362a103-842d-4d7f-b366-d1ff67fa2d21.png)

# 3. 函数

## 字符串函数

![ace95050-7c71-4967-a9e3-8740611bad6a](image/ace95050-7c71-4967-a9e3-8740611bad6a.png)

## 数值函数

![fe4b55e5-4166-4d50-90ce-23d00da3d5ba](image/fe4b55e5-4166-4d50-90ce-23d00da3d5ba.png)

## 日期函数

![919497c2-411c-4b13-acb3-430580a71c0c](image/919497c2-411c-4b13-acb3-430580a71c0c.png)

## 流程函数

![126ab987-0b05-48f4-bf1c-75a223b4ec61](image/126ab987-0b05-48f4-bf1c-75a223b4ec61.png)

# 约束

![35fa9692-3358-42b5-ba98-0a2ca707245f](image/35fa9692-3358-42b5-ba98-0a2ca707245f.png)

## 外键约束

![15e43763-40ff-42be-b2f3-72f155d4fdb9](image/15e43763-40ff-42be-b2f3-72f155d4fdb9.png)

### 添加删除外键

![369ad105-da20-4617-9ffd-2e8573ce3951](image/369ad105-da20-4617-9ffd-2e8573ce3951.png)

### 删除/更新行为

![bc262b02-d409-46d6-a095-151d4e74d3d1](image/bc262b02-d409-46d6-a095-151d4e74d3d1.png)



# 多表查询

## 多表关系

### 一对多

![b611892f-5df7-4b87-be37-fadd62b29e7d](image/b611892f-5df7-4b87-be37-fadd62b29e7d.png)

### 多对多

![5e61fa3e-07b1-487f-b902-5c034eb76867](image/5e61fa3e-07b1-487f-b902-5c034eb76867.png)

### 一对一

![872d9b67-95b3-4976-a0ad-53fe802c38e9](image/872d9b67-95b3-4976-a0ad-53fe802c38e9.png)

## 多表查询

### 概述

![46ac11ab-5234-46ac-8834-c3e1e47dc4fd](image/46ac11ab-5234-46ac-8834-c3e1e47dc4fd.png)

### 分类

![c2dfad9d-8499-4634-8bc5-c26c1403619c](image/c2dfad9d-8499-4634-8bc5-c26c1403619c.png)



## 内连接

![c3e21d41-4bcb-4427-8048-72f9750389fd](image/c3e21d41-4bcb-4427-8048-72f9750389fd.png)

## 外连接

![ccc7d3f9-4fb3-4f6e-bcf8-ef249b9a838c](image/ccc7d3f9-4fb3-4f6e-bcf8-ef249b9a838c.png)