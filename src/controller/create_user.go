package controller

import (
	"net/http"
	"peruccii/goapi/src/configurations/validation"
	"peruccii/goapi/src/controller/model/request"
	"peruccii/goapi/src/model"

	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {
    
    var userRequest request.UserRequest

    // ShouldBindJSON attempt to read the struct, if an error ( like missing fields )
    // you can handle the erros.
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        restErr := validation.ValidateUserError(err)
        c.JSON(restErr.Code, restErr) // <-- c.JSON requires an code and err obj
        
        return // <-- return to not continues the code
    }
    
    domain := model.NewUserDomain(
        userRequest.Email,
        userRequest.Password,
        userRequest.Name,
        userRequest.Age, 
    )

     if err := domain.CreateUser(); err != nil {
         c.JSON(err.Code, err)
         return 
    }

    c.JSON(http.StatusOK, "")
        
}
