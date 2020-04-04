package BunnyGo

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Bunny struct {
	Host        string
	Port        int
	controllers map[string]interface{}
}

func (bunny *Bunny) Init() {
	bunny.controllers = make(map[string]interface{})
}

func (bunny *Bunny) router(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	urlArray := strings.Split(path, "/")
	if urlArray[1] == "" {
		urlArray[1] = "Index"
	}
	mod := strings.ToLower(urlArray[1])
	action := "index"
	if len(urlArray) > 2 && urlArray[2] != "" {
		action = urlArray[2]
	}
	actionName := "Ac" + strings.Title(action)
	controller := bunny.controllers[mod]
	if controller != nil {
		classType := reflect.TypeOf(controller)
		_, ok := classType.MethodByName(actionName)
		if !ok {
			_, _ = fmt.Fprintf(writer, "Action %s Not Exist", action)
		} else {
			classVal := reflect.ValueOf(controller)
			actionFunc := classVal.MethodByName(actionName)
			var args []reflect.Value
			res := actionFunc.Call(args)
			_, _ = fmt.Fprintf(writer, res[0].String())
		}
	} else {
		_, _ = fmt.Fprintf(writer, "Mod %s Not Exist", mod)
	}
}

func (bunny *Bunny) Controller(controller interface{}) {
	classType := reflect.TypeOf(controller)
	clsName := classType.Name()
	clsName = strings.ToLower(strings.Replace(clsName, "Controller", "", 1))
	bunny.controllers[clsName] = controller
}

func (bunny *Bunny) Run() {
	http.HandleFunc("/", bunny.router)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Error", err)
	}
}
