package post

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruritoBlogger/todo_sample_backend/service"
)

type Controller struct{}

func (pc Controller) Index(c *gin.Context) {
	var s post.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(404)
		log.Fatalln("Failed to get data", err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) Create(c *gin.Context) {
	var s post.Service
	p, err := s.CreatePost(c)

	if err != nil {
		c.AbortWithStatus(404)
		log.Fatalln("Failed to create data", err)
	} else {
		c.JSON(201, p)
	}
}
