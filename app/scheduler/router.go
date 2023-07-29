package scheduler

import "github.com/ramabmtr/go-barebone/app/scheduler/handler"

var mapCronFunc = map[string]func(){
	"ping": CronFuncContextWrapper(handler.Ping),
}
