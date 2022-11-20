package promplugin

import (
	"log"
	"net/http"
	"plugin"
)

var(
	PromHandler func() http.Handler
)

func InitPlugin(soPath string){

	var(
		ok bool
	)

	plugin, err := plugin.Open(soPath)
	if err != nil {
		log.Println(err)
		return
	}

	promHandler,err :=plugin.Lookup("PromHandler")
	if err != nil{
		log.Println(err)
		return
	}

	PromHandler,ok = promHandler.(func() http.Handler)
	if !ok{
		log.Println("!ok:PromHandler")
		return
	}
}
