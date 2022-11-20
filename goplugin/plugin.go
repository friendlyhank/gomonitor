package goplugin

import (
	"log"
	"net/http"
	"plugin"
)

var(
	PromHandler func() http.Handle
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

	PromHandler,ok = promHandler.(func() http.Handle)
	if !ok{
		log.Println("!ok:PromHandler")
		return
	}
}
