package service

import (
	"peruccii/goapi/src/configurations/rest_err"
	"peruccii/goapi/src/model"
)

type userDomainInterface struct {}

type UserDomainService interface {
    CreateUser(model.UserDomainInterface)   *rest_err.RestErr // <-- create and may return an error
    UpdateUser(string,   model.UserDomainInterface)  *rest_err.RestErr // <-- pass user id and your content
    FindUser(string)    (*model.UserDomainInterface, *rest_err.RestErr) // <-- return User or error
    DeleteUser(string)  *rest_err.RestErr
}
