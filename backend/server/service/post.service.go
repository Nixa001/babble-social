package service

import (
	"backend/models"
	r "backend/server/repositories"
)

type PostService struct {
	PostRepo r.PostRepository
}

func (p *PostService) init() {
	p.PostRepo = *r.PostRepo
}

func (p *PostService) CreatePost(post models.Post) (bool, models.Errormessage) {
	Not_ok, err := p.PostRepo.CreatePost(post)
	if Not_ok {
		return true, err
	}
	return false, models.Errormessage{}
}

func (p *PostService) GetPost(Id int) ([]models.DataPost, error) {
	posts, err := p.PostRepo.LoadPost(Id)
	if posts == nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) FetchPostGroup(Id int) (models.DataPost, error) {
	post, err := p.PostRepo.LoadPostGroup(Id)
	if err != nil {
		return models.DataPost{}, err
	}
	return post, nil
}