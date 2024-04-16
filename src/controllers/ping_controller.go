//Used for cloud providers (they require an endpoint (GET) that they can hit and you respond OK
//We use for check app for OK running

package controllers

import "net/http"

const (
	pong = "pong"
)

var (
	PingContoller pingContollerInterface = &pingContoller{}
)

type pingContollerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingContoller struct{}

func (c *pingContoller) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
