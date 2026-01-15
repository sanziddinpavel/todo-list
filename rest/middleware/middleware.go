package middleware

import "Todo-list/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}

}
