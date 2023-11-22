package main

import (
	"fmt"
	"github.com/HadesTso/new.hadestso.top/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
