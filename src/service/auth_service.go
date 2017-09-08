package service

import (
    "github.com/gin-gonic/gin"
    "util"
    "define"
    "log"
    "github.com/valyala/fasthttp"
    "fmt"
)

type UserInfo struct {
    Mobile string `json:"mobile"`
    Code   string `json:"code"`
}

var infos = make(map[string]string)

func Ping(context *gin.Context) {
    context.Data(200, "text/plain", []byte("Pong"))
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

                code, body, err := fasthttp.Get(nil, "http://" + define.GatewayIP + ":" + define.GatewayPort + "/wifidog/auth?token=" + util.UUID())
                if err != nil {
                    log.Println("err -->", err)
                    context.String(500, "request wifidog gateway error")
                } else {
                    respString := string(body)
                    context.String(code, util.ParseRedirectUrl(respString))
                }
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
    case "login", "counters":
        context.Data(200, "text/plain", []byte("Auth: 1"))
    default:
        context.Data(200, "text/plain", []byte("Auth: 0"))
    }
}

func Portal(context *gin.Context) {
    // todo use origin request url to redirect
    context.HTML(200, "portal.html", gin.H{})
}