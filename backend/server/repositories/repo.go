package repositories

import (
	db "backend/database"
)

type BaseRepo struct {
	DB        *db.Database
	TableName string
}

var (
	// UserRepo is the repository for user
	UserRepo = &UserRepository{}
	// SessionRepo is the repository for session
	SessionRepo = &SessionRepository{}
)

func init() {
	UserRepo.init()
	SessionRepo.init()
}
