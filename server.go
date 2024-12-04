package main

import (
	"pro/internal/dbConfig"
	"pro/internal/auth"
	"pro/internal/controller"
	"pro/internal/global"
)


func main() {

    global.DB = dbConfig.DbConfig()

    global.Gin.GET("/hello", controller.HelloFunc)
    global.Gin.POST("/createUser", controller.CreateUser)
    global.Gin.POST("/registerUser", controller.RegisterUser)
    global.Gin.GET("/loginUser", controller.LoginUser)
    global.Gin.GET("/getUser",auth.AuthMiddleware(), controller.GetUserInfo)

    global.Gin.Run(":8080")
}

// http://localhost:8080/loginUser
// {
//     "username": "newuser2",
//     "password": "password123"
// }
// "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InBhc3N3b3JkMTIzIiwidXNlcm5hbWUiOiJuZXd1c2VyMiJ9.J5a_P65VcjHmzTSrWeDTZipKbVozMBrPUcOIni_-_RU"
