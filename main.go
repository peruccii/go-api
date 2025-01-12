package main

import (
	"log"
	"peruccii/goapi/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // todo: init routes passing router groups
    router := gin.Default()

    routes.InitRoutes(&router.RouterGroup)

    err = router.Run(":8080") // err bc this action can return an error

    if err != nil {
        log.Fatal(err)
    }
}
