package middlewares

type MiddlewareChain struct {
}

func (ms *MiddlewareChain) Regist(m Middleware) {

}

func (ms *MiddlewareChain) ExecMiddlewareChain(env Env) MiddlewareResponder {

}
