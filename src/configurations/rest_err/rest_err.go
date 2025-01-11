package rest_err

import (
	"net/http"

)

// like a object but i cant put methods on it
type RestErr struct {
    Message string  `json:"message"`
    Err     string  `json:"error"`
    Code    int     `json:"code"`
    Causes  []Causes`json:"causes,omitempty"` // dont appear if causes be empyt
}

type Causes struct {
    Field   string  `json:"fields"`
    Message string  `json:"message"`
}

// constructor ??                                               // return
func NewRestErr(message, err string, code int, causes []Causes) *RestErr { 
    return &RestErr{        // instance of RestErr 
        Message: message,
        Code: code,
        Causes: causes,
        Err: err,
    }
}

// *RestErr <-- i am pointing to RestErr memory address, and not copying it.
// &RestErr <-- so i have to create and return a instance of RestErr

func NewBadRequestError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err: "bad request",
        Code: http.StatusBadRequest,
    }
}
