package service

import (
	r "backend/server/repositories"
)

type PostService struct {
	PostRepo r.PostRepository
	SessRepo r.SessionRepository
}

func (p *AuthService) init() {
	p.UserRepo = *r.PostRepo
	p.SessRepo = *r.SessionRepo
}
