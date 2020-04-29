package demo

import (
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 鉴权
func MiddlewareAuth(r *ghttp.Request) {
	token := r.Get("token")
	if token == "123456" {
		r.Response.Writeln("auth")
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// 跨域访问
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.Writeln("cors")
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 错误处理
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("哎哟我去，服务器居然开小差了，请稍后再试吧！")
	}
}

func main() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS, MiddlewareAuth, MiddlewareErrorHandler)
		group.ALL("/user/list", func(r *ghttp.Request) {
			r.Response.Writeln("list")
		})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.ALL("/error", func(r *ghttp.Request) {
				panic("GG") // panic 异常
			})
		})
	})
	s.SetPort(8199)
	s.Run()
}