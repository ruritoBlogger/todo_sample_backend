package post

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ruritoBlogger/todo_sample_backend/db"
	"github.com/ruritoBlogger/todo_sample_backend/entity"
)

type Service struct{}

func (s Service) GetAll() ([]entity.Post, error) {
	db := db.GetDB()
	var posts []entity.Post

	_, err := db.Select(&posts, "select * from posts order by post_id")
	count, err := db.SelectInt("select count(*) from posts")
	log.Println(posts)
	log.Println(count)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s Service) CreatePost(c *gin.Context) (entity.Post, error) {
	db := db.GetDB()
	var post entity.Post

	if err := c.BindJSON(&post); err != nil {
		return post, err
	}

	if err := db.Insert(&post); err != nil {
		return post, err
	}

	return post, nil
}
