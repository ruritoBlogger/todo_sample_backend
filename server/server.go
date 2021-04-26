package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ruritoBlogger/todo_sample_backend/controller"
)

func router() *gin.Engine {
	r := gin.Default()

	p := r.Group("/posts")
	{
		ctrl := post.Controller{}
		p.GET("", ctrl.Index)
		p.POST("", ctrl.Create)
	}

	return r
}

func Start() {
	r := router()
	r.Run()
}
