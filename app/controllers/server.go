package controllers

import (
	"github.com/revel/revel"
)

var status = 200

type Server struct {
	*revel.Controller
}

func (c Server) Online() revel.Result {
	status = 200
	return c.RenderJSON(status)
}

func (c Server) Offline() revel.Result {
	status = 404
	return c.RenderJSON(status)
}

func (c Server) Status() revel.Result {
	return c.RenderJSON(status)
}
