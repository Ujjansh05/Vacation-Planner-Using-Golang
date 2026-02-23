package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r:= ginDefault()
	routes.GetVactionRouter(r)
	r.run()
}