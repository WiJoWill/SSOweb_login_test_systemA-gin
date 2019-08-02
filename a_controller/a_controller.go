package a_controller

/* 目标功能
1、 识别cookie是否登录，如果登录了就进入网页
2、 cookie未登录，导向controller，获取对应的信息
3、 如果登录了，则将token储存
3.1 申请token验证，如果验证成功则允许进入网页
4、 如果未登录，未获取合适的token信息，则跳转至登录页面 0717 登录页面登录后怎么回到这个系统呢！！！！
	是不是可以让用户自动返回？要不让用户去一个叫做homepage之类的地方吧...
5、

 */

/* 使用session的目标功能
1、 使用session查看是否登录
 */
import (
	"fmt"
	_ "github.com/fwhezfwhez/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web_login/controller"
	"web_login_test_systemA/a_model"
)

var user_claims *controller.UserClaims


//加一个请求验证功能 调用sso里面的validate函数（应该要去写在token controller或者login controller里面 然后这样子就可以不用session了）
//子系统的Validate功能与sso不一样，是解析url里是否带有sub_token
func Validate() gin.HandlerFunc{
	return func(c *gin.Context) {
		//token := c.Query("sub_token")
		token, err := c.Cookie("token")
		if token == "" {
				c.Abort()
				c.Redirect(302, "http://127.0.0.1:8081/login?redirectURL=http://127.0.0.1:9091")
				return
			//}
		}
		log.Print("get token: ", token)
x
		/*
		//检验ip地址是否与登录时的ip地址一致
		user_ip := c.ClientIP();
		fmt.Println(user_ip)
		if a_model.CheckIPAndTokeninASystem(user_ip, token) == false{
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "A系统内，IP地址与登录地点不符，请重新登录",
			})
			c.Abort()
			c.Redirect(302,"http://127.0.0.1:8081/login")
			return
		}

		 */

		j := controller.NewJWT()
		//解析token包含的信息
		claims, err := j.ParseToken(token)
		user_claims = claims;
		if err != nil {
			if err == controller.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				c.Redirect(302, "http://127.0.0.1:8081/login?redirectURL=http://127.0.0.1:9091")
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)

	}
}

//通过sub_token解析用户信息，sub_token类似于service ticket，被认证有效后给予用户信息
func AGet(c * gin.Context) {
	claims := c.MustGet("claims").(*controller.UserClaims)
	if claims != nil {
		sub_token := c.Query("sub_token")
		a_model.ConnectRedis()
		a_model.SetTokeninASystem(sub_token, user_claims)

		c.HTML(http.StatusOK, "a.html", gin.H{"title": "a测试页"})
	}
}

func APost (c * gin.Context){
	token, err := c.Cookie("token")
	if err != nil{
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"code": 0, "message":"获取密钥过程发生错误，请检查登录信息"})
	}
	//这里缺少验证ip的部分 稍后添加

	//
	data := controller.RequestUserInfo(token)
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{"code": 0, "UserInfo": data})
}