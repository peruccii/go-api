package controller

import (
	"net/http"
	"peruccii/goapi/src/configurations/validation"
	"peruccii/goapi/src/controller/model/request"

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

    c.JSON(http.StatusOK, userRequest)
        
}
