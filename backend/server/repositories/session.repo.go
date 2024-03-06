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
	_, err := s.Db.Exec("INSERT INTO sessions (id, expiration_date, data) VALUES ($1, $2, $3)", session.Id, session.ExpirationDate, session.Data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionRepository) GetSession(id string) (*models.Session, error) {
	var session models.Session
	err := s.Db.QueryRow("SELECT id, expiration_date, data FROM sessions WHERE id = $1", id).Scan(&session.Id, &session.ExpirationDate, &session.Data)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionRepository) GetSessionByUserId(userId string) (*models.Session, error) {
	var session models.Session
	err := s.Db.QueryRow("SELECT id, expiration_date, data FROM sessions WHERE user_id = $1", userId).Scan(&session.Id, &session.ExpirationDate, &session.Data)
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
	_, err := s.Db.Exec("UPDATE sessions SET expiration_date = $1, data = $2 WHERE id = $3", session.ExpirationDate, session.Data, session.Id)
	if err != nil {
		return err
	}
	return nil
}
