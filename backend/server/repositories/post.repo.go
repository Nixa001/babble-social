package repositories

import (
	db "backend/database"
	"backend/models"
)

type PostRepository struct {
	BaseRepo
}

func (p *PostRepository) init() {
	p.DB = db.DB
	p.TableName = "posts"
}

func (p *PostRepository) CreatePost(post models.Post) error {
	/*
	! code here........
	todo: check privacy case
	todo: if privacy = "almost private", insert in post table then insert in viewer table
	todo: else simply insert posts 
	*/
	return nil
}

func (p *PostRepository) InsertReaction(post models.PostReact) error {
	/*
	! code here........
	todo: check post's validity
	todo: insert into database 
	*/
	return nil
}

func (p *PostRepository) LoadPost(lastPostID int) (models.DataPost, error) {
	/*
	! code here........
	todo: retrieve from database each post's data (reaction, comment and comment reaction)
	*/
	return nil, nil
}
