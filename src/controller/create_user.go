package controller

import (
	"fmt"
	"net/http"
	"peruccii/goapi/src/configurations/rest_err"
	"peruccii/goapi/src/controller/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    
    var userRequest request.UserRequest

    // ShouldBindJSON attempt to read the struct, if an error ( like missing fields )
    // you can handle the erros.
    if err := c.ShouldBindJSON(&userRequest); err != nil {
        restErr := rest_err.NewBadRequestError(
            fmt.Sprintf("there are some incorrect fields, error=%s\n", err.Error()))
            // err.Error <-- will take the method on rest_err that returns a message error    
            

        c.JSON(restErr.Code, restErr) // <-- c.JSON requires an code and err obj
        
        return // <-- return to not continues the code
    }

    c.JSON(http.StatusOK, userRequest)
        
}
