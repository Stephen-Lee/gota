package middlewares

import (
	"time"
)

type TimerMiddleware struct {
}

func (timer *TimerMiddleware) Call(env *Env) MiddlewareResponder {
	start_time := time.Now()
	App.Call(Env)
	end_time := time.Now()
}
