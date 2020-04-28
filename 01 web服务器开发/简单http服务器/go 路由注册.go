package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writefln("哈喽世界！")
	})
	s.BindHandler("/name", func(r *ghttp.Request) {
		r.Response.Writefln("michael！")
	})
	/*
		以上示例中展示了gf框架支持的三种模糊匹配路由规则，:name、*any、{field}
		分别表示命名匹配规则、模糊匹配规则及字段匹配规则。不同的规则中使用/符号来划分层级，路由检索采用深度优先算法，层级越深的规则优先级也会越高。
	*/
	s.BindHandler("/:name/hello", func(r *ghttp.Request) {
		r.Response.Writefln(r.Router.Uri)
	})
	s.BindHandler("/*any/hello1", func(r *ghttp.Request) {
		r.Response.Writefln(r.Router.Uri)
	})
	s.BindHandler("/{filed}/hello2", func(r *ghttp.Request) {
		r.Response.Writefln(r.Router.Uri)
	})
	s.SetPort(80) // 监听多端口
	s.Run()
}
