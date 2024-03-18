package repositories

import (
	db "backend/database"
	"backend/models"
)

type CommentRepository struct {
	BaseRepo
}

func (c *CommentRepository) init() {
	c.DB = db.DB
	c.TableName = "comment"
}

func (c *CommentRepository) CreateComment(Comment models.Comment) error {
	/*
		! code here........
		todo: check post's validity
		todo: create media in local
		todo: inssert comment in database
	*/
	return nil
}

func (c *CommentRepository) InsertReaction(reaction models.CommentReact) error {
	/*
		! code here........
		todo: check post's validity
		todo: check comment's validity
		todo: insert into database
	*/
	return nil
}

func (c *CommentRepository) LoadComment(lastComment models.Comment) (models.DataComment, error) {
	/*
		! code here........
		todo: retrieve from database each comment's data (reaction + rows value)
	*/
	return nil, nil
}
