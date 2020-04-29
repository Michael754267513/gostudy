package demo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/input", func(r *ghttp.Request) {
		r.Response.Writeln(r.Get("amount"))
	})
	s.BindHandler("/query", func(r *ghttp.Request) {
		r.Response.Writeln(r.GetQuery("amounts"))
	})
	s.SetPort(8199)
	s.Run()
}