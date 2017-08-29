package main

import (
    "github.com/gin-gonic/gin"
    "service"
)

func main() {
    serv := gin.Default()
    serv.LoadHTMLGlob("**/templates/*")

    serv.GET("/status/", func(c *gin.Context) {
        c.String(200, "OK")
    })

    g := serv.Group("/wgo")
    {
        g.GET("/ping/", service.Ping)
        g.GET("/login/", service.LoginPage)
        g.POST("/code/", service.RequestCode) // login check
        g.POST("/checkin/", service.CheckIn) // login check
        g.GET("/auth/", service.Auth)
        g.GET("/portal/", service.Portal)
    }

    serv.Run()
}
