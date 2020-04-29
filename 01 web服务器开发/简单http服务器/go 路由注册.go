package demo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 函数方法注册
func hello3(r *ghttp.Request) {
	r.Response.Write("hello1")
}
// 对象注册
type Ojreg struct {}

func (b *Ojreg) Index (r *ghttp.Request)  {
	r.Response.Write("index")
}
func (b *Ojreg) Test2 (r *ghttp.Request)  {
	r.Response.Write("show")
}

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
	s.BindHandler("GET:/a/:name/hello", func(r *ghttp.Request) {
		//r.Response.WriteJson(r.Router)
		r.Response.Writeln(r.Get("name"))
	})
	s.BindHandler("POST:/b/*any/hello1", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router)
	})
	s.BindHandler("/c/{filed}/hello2", func(r *ghttp.Request) {
		r.Response.Writefln(r.Router.Uri)
	})
	// 函数方法注册
	s.BindHandler("/func/bind",hello3)
	// 对象注册
	b := new(Ojreg)
	s.BindObject("/object",b)
	s.SetPort(80) // 监听多端口
	s.Run()
}
