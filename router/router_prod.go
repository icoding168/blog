//go:build prod
// +build prod

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/zxysilent/logs"
)

const AppJsUrl = "/static/js/app.min.js"
const AppCssUrl = "/static/css/app.min.css"

func init() {
	logs.SetLevel(logs.LWARN)
}

// RegDocs 注册文档
// prod[正式] 模式不需要文档
func RegDocs(engine *echo.Echo) {
}

// midLogger 中间件-日志记录
func midLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return next
}

/*  生产模式-编译 取消文档-请求记录
 *  生成文档 swag init
 *  go build -tags=prod  .\main.go
 *
 *  开发模式-编译 添加文档-请求记录
 *  go build .\main.go
 */
