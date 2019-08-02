### Subsystem "A" for Testing SSO System by Golang

### Golang 编写的测试系统，用于测试SSO单点登录系统

###### sso system address: https://github.com/WiJoWill/SSOweb_login-gin

[主系统请点击此处](https://github.com/WiJoWill/SSOweb_login-gin)

----

#### 项目简介/Instruction

此项目是作为SSO单点登录系统的子系统，用于测试SSO功能。

----

#### 功能特性/Function

— 简单的访问主系统请求验证

— 存储Token

— 显示登录的用户信息

----

#### 环境依赖/Dependencies

golang编程语言

gin框架(https://github.com/gin-gonic/gin)

（以及一些杂七杂八的，源码的import里都有）

----

#### 目录结构描述/Catalog

##### a_controller

- a_controller.go

  - 客户端验证用户信息：是否登录，是否携带token，token是否过期

  - ```go
    func Validate() gin.HandlerFunc{}
    ```

  - 获取sub_token密钥，在本系统内设置相关信息：token，用户ip地址等。并且加载a.html

  - ```go
    func AGet (C * gin.Context) {}
    ```

  - 向主系统发送请求，获取登录用户的个人信息

  - ```go
    func APost (C * gin.Context){}
    ```

- a_controller.go
  - 「未启用」

##### a_model

- a_model.go

  - 「未启用」

- a_redis_model.go

  - 连接Redis

  - ```go
    func ConnectRedis(){}
    ```

  - 存储访问用户的token到本系统

  - ```go
    func SetTokeninASystem(token string, user *controller.UserClaims){}
    ```

  - 验证token是否存在

  - ```go
    func CheckToken(token string) bool{}
    ```

  - 在本地redis库（新建一个db）存储ip与对应token

  - ```go
    func SetTokenIPinASystem (ip string, token string){}
    ```

  - 在本地redis库查询IP和对应token

  - ```go
    func CheckIPAndTokeninASystem(ip string,token string) bool{}
    ```

##### a_router

- router.go

  - 启用路由

  - ```go
    func InitRouter() *gin.Engine {}
    ```

##### a_view

- a.html
  - 「请看注释」

##### static

- a_user_info.js
  - 「请看注释」

----

#### 内容更新/Update

- Aug.2, 2019: 上传代码
- Aug.2, 2019: 添加了获取用户信息时，需要进行一次IP地址验证，具体请看用户主系统

----

#### 主系统/Main System

SSO单点登录系统用户主系统模块：https://github.com/WiJoWill/SSOweb_login-gin