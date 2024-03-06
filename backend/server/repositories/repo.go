package repositories

import "database/sql"

type BaseRepo struct {
	Db        *sql.DB
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
