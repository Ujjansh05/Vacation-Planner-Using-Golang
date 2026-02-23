package main

import (
	"github.com/gin-gonic/gin"
	"langchaingo/routes"
)

func main(){
	r:= ginDefault()
	routes.GetVactionRouter(r)
	r.run()
}