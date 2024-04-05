package repositories

import (
	db "backend/database"
	opt "backend/database/operators"
	q "backend/database/query"
	"backend/models"
)

type FollowRepository struct {
	BaseRepo
}

func (f *FollowRepository) init() {
	f.DB = db.DB
	f.TableName = "users_followers"
}

func (f *FollowRepository) Getfollower(id int) ([]models.UserFollower, error) {
	rows, err := f.DB.GetAllFrom(f.TableName, q.WhereOption{"user_id_followed": opt.Equals(id)}, "", nil)
	if err != nil {
		return nil, err
	}
	var IdUsers []models.UserFollower
	for rows.Next() {
		var iduser models.UserFollower
		err := rows.Scan(&iduser.User_id_followed, &iduser.User_id_follower)
		if err != nil {
			return nil, err
		}
		IdUsers = append(IdUsers, iduser)
	}
	return IdUsers, nil
}

func (f *FollowRepository) Getfollowing(id int) ([]models.UserFollower, error) {
	rows, err := f.DB.GetAllFrom(f.TableName, q.WhereOption{"user_id_follower": opt.Equals(id)}, "", nil)
	if err != nil {
		return nil, err
	}
	var IdUsers []models.UserFollower
	for rows.Next() {
		var iduser models.UserFollower
		err := rows.Scan(&iduser.User_id_followed, &iduser.User_id_follower)
		if err != nil {
			return nil, err
		}
		IdUsers = append(IdUsers, iduser)
	}
	return IdUsers, nil
}

func (f *FollowRepository) FollowUser(id int, idToFollow int) error {
	err := f.DB.Insert(f.TableName, models.UserFollower{User_id_followed: idToFollow, User_id_follower: id})
	if err != nil {
		return err
	}
	return nil
}
