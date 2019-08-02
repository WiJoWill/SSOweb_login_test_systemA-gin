package a_model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"web_login/controller"
)

var c redis.Conn
var err error

func ConnectRedis() {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis error", err.Error())
	}else{
		fmt.Println("Succeed in Connecting to Redis")
	}
}

//将子系统的token存入子系统的redis，key为token，value为user信息
func SetTokeninASystem(token string, user *controller.UserClaims){
	if c == nil{
		return;
	}
	c.Do("SELECT", 1)
	user_id := user.ID
	user_username := user.Username
	v, err := c.Do("SETEX", token, 300, "user_id is "+user_id+" ; "+"username is "+user_username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func CheckToken(token string) bool{
	if c == nil{
		return false
	}

	c.Do("SELECT", 1)
	existStatus, err := c.Do("EXISTS", token)
	fmt.Println(existStatus)
	if err != nil{
		return false
	}
	return true
}

//在redis库中储存对应密钥的ip地址
func SetTokenIPinASystem (ip string, token string){
	if c == nil{
		return;
	}
	c.Do("SELECT", 3)
	fmt.Println(ip)

	v, err0 := c.Do("SET", ip, token)
	if err0 != nil {
		fmt.Println(err0)
		return
	}

	fmt.Println(v)
	set_ip, err1 := c.Do("GET", ip)

	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("ip并没有在A系统存储成功")
	}else{
		fmt.Printf("存储了ip：")
		fmt.Println(set_ip)
	}
}

func CheckIPAndTokeninASystem(ip string,token string) bool{
	if c == nil{
		return false
	}
	c.Do("SELECT", 3)
	_, err0 := c.Do("EXISTS", ip)

	if err0 != nil {
		fmt.Printf("web_login_test_systemA - a_redis_model func CheckIPAndToken 出错err0:" )
		fmt.Println(err0)
		return false
	}

	Set_token, err1 := redis.String(c.Do("GET", ip))
	if err1 != nil {
		fmt.Printf("web_login_test_systemA - a_redis_model func CheckIPAndToken 出错err1:")
		fmt.Println(err1)
		return false
	}

	if Set_token != token{
		return false
	}
	return true
}
