package model

import (
	"crypto/md5"
	"encoding/hex"
	"peruccii/goapi/src/configurations/rest_err"
)

// begin to lowerCase bc its private
type UserDomain struct {
    Email       string
    Password    string
    Name        string
    Age         string
}

func (ud *UserDomain) EncryptPassword() {
    hash := md5.New()
    defer hash.Reset()
    hash.Write([]byte(ud.Password))
    ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
    CreateUser()         *rest_err.RestErr // <-- create and may return an error
    UpdateUser(string) *rest_err.RestErr // <-- pass user id and your content
    FindUser(string)               (*UserDomain, *rest_err.RestErr) // <-- return User or error
    DeleteUser(string)             *rest_err.RestErr
}
