package repositories

import "backend/models"

type UserRepository struct {
	BaseRepo
}

func (u *UserRepository) init() {
	u.TableName = "users"
}

func (u *UserRepository) CreateUser(user *models.User) error {
	_, err := u.Db.Exec("INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", user.Firstname, user.Lastname, user.Username, user.Gender, user.Email, user.Password, user.UserType, user.BirthDate, user.Avatar, user.AboutMe)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := u.Db.QueryRow("SELECT id, first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me FROM users WHERE id = $1", id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.Db.QueryRow("SELECT id, first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me FROM users WHERE email = $1", email).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	_, err := u.Db.Exec("UPDATE users SET first_name = $1, last_name = $2, user_name= $3, gender= $4, email= $5, password= $6, user_type= $7, birth_date= $8, avatar= $9, about_me= $10 WHERE id = $11", user.Firstname, user.Lastname, user.Username, user.Gender, user.Email, user.Password, user.UserType, user.BirthDate, user.Avatar, user.AboutMe, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteUser(id int) error {
	_, err := u.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := u.Db.Query("SELECT id, first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
