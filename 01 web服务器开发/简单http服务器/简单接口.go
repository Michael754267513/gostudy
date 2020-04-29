package demo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//func main() {
//	s := g.Server()
//	s.BindHandler("/", func(r *ghttp.Request) {
//		r.Response.Write("哈喽世界！")
//	})
//	s.SetPort(80, 8200, 8300) // 监听多端口
//	s.Run()
//}
func main() {
	s := g.Server()
	s.Domain("127.0.0.1").BindHandler("/hello1", hello1)
	s.Domain("127.0.0.1").BindHandler("/hello2", hell2)
	s.SetPort(80)
	s.Run()
}
func hello1(r *ghttp.Request) {
	r.Response.Write("hello1")
}

func hell2(r *ghttp.Request) {
	r.Response.Write("hello2")
}
