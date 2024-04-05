package service

import (
	"backend/models"
	r "backend/server/repositories"
)

type FollowService struct {
	FollowRepo r.FollowRepository
	UserRepo   r.UserRepository
}

func (f *FollowService) init() {
	f.FollowRepo = *r.FollowRepo
}

func (f *FollowService) GetFollowersByID(id int) ([]models.User, error) {
	IdUsers, err := f.FollowRepo.Getfollower(id)
	if err != nil {
		return nil, err
	}
	followers := make([]models.User, 0)
	for _, idUser := range IdUsers {
		user, err := f.UserRepo.GetUserById(idUser.User_id_follower)
		if err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}

func (f *FollowService) GetFollowingsByID(id int) ([]models.User, error) {
	IdUsers, err := f.FollowRepo.Getfollowing(id)
	if err != nil {
		return nil, err
	}
	followings := make([]models.User, 0)
	for _, idUser := range IdUsers {
		user, err := f.UserRepo.GetUserById(idUser.User_id_followed)
		if err != nil {
			return nil, err
		}
		followings = append(followings, user)
	}
	return followings, nil
}

func (f *FollowService) FollowUser(id int, idToFollow int) error {
	err := f.FollowRepo.FollowUser(id, idToFollow)
	if err != nil {
		return err
	}
	return nil
}
