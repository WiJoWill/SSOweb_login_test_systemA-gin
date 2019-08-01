package main

import (
	"web_login/database"
	"web_login_test_systemA/a_router"
)

func main(){
	database.InitMysql()
	router := a_router.InitRouter()
	router.Run(":9091")
}

