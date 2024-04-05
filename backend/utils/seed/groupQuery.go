package seed

import (
	"backend/models"
	"database/sql"
	"fmt"
	"log"
)

// InsertMessage insère un nouveau message dans la base de données.
func InsertGroupMessage(db *sql.DB, userSender, groupReceiver int, messageContent, date string) error {
	// Préparation de la requête SQL pour insérer un nouveau message.
	query := `INSERT INTO messages (user_id_sender, group_id_receiver, message_content, date) VALUES (?, ?, ?, ?);`
	// Exécution de la requête avec les valeurs fournies.
	_, err := db.Exec(query, userSender, groupReceiver, messageContent, date)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion du message group:", err)
		return err
	}

	return nil
}

// getGroupMessage récupère les messages pour un groupe spécifique
func GetGroupMessage(db *sql.DB, groupIDReceiver int) ([]models.Chat, error) {
	// Préparation de la requête SQL
	query := `SELECT * FROM messages WHERE group_id_receiver = ?`
	rows, err := db.Query(query, groupIDReceiver)
	if err != nil {
		log.Printf("Erreur lors de la récupération des messages du groupe: %v", err)
		return nil, err
	}
	defer rows.Close()

	// Préparation du slice pour stocker les messages
	var messages []models.Chat

	// Parcours des résultats et remplissage du slice
	for rows.Next() {
		var message models.Chat
		var userReceiver, groupIDReceiver *int
		err := rows.Scan(&message.ID, &message.UserSender, &userReceiver, &message.MessageContent, &groupIDReceiver, &message.Date)
		if err != nil {
			log.Printf("Erreur lors de la lecture des messages du groupe: %v", err)
			return nil, err
		}
		message.UserReceiver = userReceiver
		message.GroupIDReceiver = groupIDReceiver

		// Récupérer le first_name de l'utilisateur qui a envoyé le message
		queryUser := `SELECT first_name FROM users WHERE id = ?`
		userRow, err := db.Query(queryUser, message.UserSender)
		if err != nil {
			log.Println("Erreur lors de la récupération du first_name:", err)
			return nil, err
		}
		defer userRow.Close()
		if userRow.Next() {
			err = userRow.Scan(&message.FirstName)
			if err != nil {
				log.Println("Erreur lors de la récupération du first_name:", err)
				return nil, err
			}
		}

		messages = append(messages, message)
	}

	// Vérification d'une éventuelle erreur lors de la récupération des résultats
	if err = rows.Err(); err != nil {
		log.Printf("Erreur lors de la récupération des messages du groupe: %v", err)
		return nil, err
	}

	return messages, nil
}

func GetLastGroupMessage(db *sql.DB, msg string) (models.Chat, error) {
	query := `SELECT * FROM messages WHERE message_content = ?`
	row, err := db.Query(query, msg)
	if err != nil {
		return models.Chat{}, err
	}
	var message models.Chat
	for row.Next() {
		var IDReceiver *int
		err = row.Scan(&message.ID, &message.UserSender, &IDReceiver, &message.MessageContent, &message.GroupIDReceiver, &message.Date)
		if err != nil {
			log.Println("Erreur lors de recuperer du message:", err)
			return models.Chat{}, err
		}
		// Récupérer le first_name de l'utilisateur qui a envoyé le message
		queryUser := `SELECT first_name FROM users WHERE id = ?`
		userRow, err := db.Query(queryUser, message.UserSender)
		if err != nil {
			log.Println("Erreur lors de la récupération du first_name:", err)
			return models.Chat{}, err
		}
		defer userRow.Close()
		if userRow.Next() {
			err = userRow.Scan(&message.FirstName)
			if err != nil {
				log.Println("Erreur lors de la récupération du first_name:", err)
				return models.Chat{}, err
			}
		}
	}
	return message, err
}
