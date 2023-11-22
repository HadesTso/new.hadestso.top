package main

import "github.com/HadesTso/new.hadestso.top/framework"

func RegisterRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
