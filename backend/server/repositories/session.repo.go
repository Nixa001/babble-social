package repositories

import "backend/models"

type SessionRepository struct {
	BaseRepo
}

func (s *SessionRepository) init() {
	// s.DB =
	s.TableName = "sessions"
}

func (s *SessionRepository) CreateSession(session *models.Session) error {
	_, err := s.Db.Exec("INSERT INTO sessions (token, user_id, expiration) VALUES ($1, $2, $3)", session.Token, session.UserId, session.ExpirationDate)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) GetSession(token string) (*models.Session, error) {
	var session models.Session
	err := s.Db.QueryRow("SELECT token, user_id, expiration FROM sessions WHERE token = $1", token).Scan(&session.Token, &session.UserId, &session.ExpirationDate)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionRepository) GetSessionByUserId(userId string) (*models.Session, error) {
	var session models.Session
	err := s.Db.QueryRow("SELECT token, user_id, expiration FROM sessions WHERE user_id = $1", userId).Scan(&session.Token, &session.UserId, &session.ExpirationDate)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionRepository) DeleteSession(id string) error {
	_, err := s.Db.Exec("DELETE FROM sessions WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) UpdateSession(session *models.Session) error {
	_, err := s.Db.Exec("UPDATE sessions SET expiration= $1, user_id = $2 WHERE token = $3", session.ExpirationDate, session.UserId, session.Token)
	if err != nil {
		return err
	}
	return nil
}
