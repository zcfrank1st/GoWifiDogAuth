package service

import (
    "github.com/gin-gonic/gin"
    "util"
    "fmt"
    "define"
)

type UserInfo struct {
    Mobile string `json:"mobile"`
    Code   string `json:"code"`
}

var infos = make(map[string]string)

func Ping(context *gin.Context) {
    context.String(200, "pong")
}

func LoginPage(context *gin.Context) {
    context.HTML(200, "login.html", gin.H{
        "title": "Dejionline Gateway Service",
    })
}

// *** 自定义操作
func RequestCode(context *gin.Context) {
    var user UserInfo
    if context.BindJSON(&user) == nil {
        code := "123456"
        infos[user.Mobile] = code
        // todo send sms
        context.String(200, "")
    } else {
        context.String(500,"")
    }

}

func CheckIn(context *gin.Context) {
    var user UserInfo
    if context.BindJSON(&user) == nil {
        if v, ok := infos[user.Mobile]; ok {
            if user.Code == v {
                // todo send user mobile info to register
                delete(infos, user.Mobile)
                context.Redirect(307, fmt.Sprintf("http://%s:%s/wifidog/auth?token=%s", define.GatewayIP, define.GatewayPort, util.UUID()))
            } else {
                context.String(500, "invalid mobile or code!")
            }
        }
    }

}
// ***

func Auth(context *gin.Context) {
    stage := context.Query("stage")
    switch stage {
    case "login", "counter":
        context.String(200, "Auth: 1")
    default:
        context.String(200, "Auth: 0")
    }
}

func Portal(context *gin.Context) {
    // todo use origin request url to redirect
    context.Redirect(301, "https://www.baidu.com")
}
