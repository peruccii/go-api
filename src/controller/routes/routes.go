package routes

import (
	"peruccii/goapi/src/controller"

	"github.com/gin-gonic/gin"
)

// i will initinialize the routes at main.go 
func InitRoutes(r *gin.RouterGroup) { // <-- RouterGroup give me access to the verbs
    
    r.GET("/getUserById/:userId",           controller.FindUserById)
    r.GET("/getUserByEmail/:userEmail",     controller.FindUserByEmail)
    r.POST("/createUser",                   controller.CreateUser)
    r.PUT("/updateUser/:userId",            controller.UpdateUserById)
    r.DELETE("/deleteUser/:userId",         controller.DeleteUserById)

}
