/* A ce niveau on aura besoin des deux table Users et UsersFollowers
   Pour etre dans la liste des online users d'un Users connecter par example
   ==> Soit le users te follow ou tu le follow (Cette liste sera afficher dans la partie des users avec qui il me chatter)
       Maintenant restera a determiner si le user est en ligne ou pas (Utilisation des sessions)

	   NB : Si lordre des users par rapport aux historique des messages va etre prise en compte faudra
	   joindre le table de messages aussi
*/

package seed

import (
	"backend/models"
	"database/sql"
	"fmt"
	"sort"
	"strings"
)

// Messages entre deux users

func SelectMsgBetweenUsers(db *sql.DB, CurrentUser, UserID int) ([]models.Chat, error) {
	query := `SELECT * FROM messages WHERE (user_id_sender = ? AND user_id_receiver = ? ) OR (user_id_sender = ? AND user_id_receiver = ?);`
	rows, err := db.Query(query, CurrentUser, UserID, UserID, CurrentUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chat []models.Chat

	for rows.Next() {
		var m models.Chat
		var groupIDReceiver *int
		// var userIDReceiver *int
		err := rows.Scan(&m.ID, &m.UserSender, &m.UserReceiver, &m.MessageContent, &m.GroupIDReceiver, &m.Date)
		if err != nil {
			return nil, err
		}
		m.GroupIDReceiver = groupIDReceiver
		row, _ := db.Query("SELECT first_name FROM users WHERE id = ?", m.UserSender)
		defer row.Close()

		for row.Next() {
			if err := row.Scan(&m.FirstName); err != nil {
				fmt.Println("error scanning first_name")
			}
		}
		chat = append(chat, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chat, nil
}

// SelectFollowersAndFollowing récupère tous les utilisateurs suivis et les suivants d'un utilisateur spécifique.
func SelectFollowersAndFollowing(db *sql.DB, userID int) ([]models.User, error) {
	// SQL query to retrieve followed and followers.
	query := `
		SELECT u.* FROM users u
		JOIN users_followers uf ON u.id = uf.user_id_followed OR u.id = uf.user_id_follower
		WHERE uf.user_id_follower = ? OR uf.user_id_followed = ?
	`

	rows, err := db.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use a map to keep track of unique users.
	userMap := make(map[int]models.User)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
		if err != nil {
			return nil, err
		}
		// Add the user to the map if it's not already present.
		if _, exists := userMap[user.ID]; !exists {
			userMap[user.ID] = user
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert the map to a slice.
	var users []models.User
	for _, user := range userMap {
		users = append(users, user)
	}

	return users, nil
}

// ListeUsers retourne un tableau de tableaux d'utilisateurs en se basant sur les deux fonctions
// SelectMsgBetweenUsers et SelectFollowersAndFollowing. Les utilisateurs sont triés par ordre des
// derniers messages s'il y en a, sinon par ordre alphabétique des noms des utilisateurs.
func ListeUsers(db *sql.DB, userID int) ([][]models.User, error) {
	// Récupérer tous les utilisateurs suivis et suivants.
	followersAndFollowing, err := SelectFollowersAndFollowing(db, userID)
	if err != nil {
		return nil, err
	}

	// Trier les utilisateurs par ordre alphabétique des noms de famille.
	sort.Slice(followersAndFollowing, func(i, j int) bool {
		return followersAndFollowing[i].Firstname < followersAndFollowing[j].Lastname
	})

	// Récupérer les messages entre l'utilisateur actuel et chaque utilisateur suivi/suivi.
	var messages []models.Chat
	for _, user := range followersAndFollowing {
		userMessages, err := SelectMsgBetweenUsers(db, userID, user.ID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, userMessages...)
	}

	// Si des messages ont été trouvés, trier les utilisateurs par ordre des derniers messages.
	if len(messages) > 0 {
		sort.Slice(followersAndFollowing, func(i, j int) bool {
			// Trouver le dernier message de chaque utilisateur.
			lastMessageI := findLastMessage(messages, followersAndFollowing[i].ID)
			lastMessageJ := findLastMessage(messages, followersAndFollowing[j].ID)

			// Trier par date du dernier message.
			return strings.Compare(lastMessageI.Date, lastMessageJ.Date) > 0
		})
	}

	// Grouper les utilisateurs par tableau de tableaux.
	var groupedUsers [][]models.User
	groupSize := 10 // Taille du groupe, ajustez selon vos besoins.
	for i := 0; i < len(followersAndFollowing); i += groupSize {
		end := i + groupSize

		if end > len(followersAndFollowing) {
			end = len(followersAndFollowing)
		}

		groupedUsers = append(groupedUsers, followersAndFollowing[i:end])
	}
	return groupedUsers, nil
}

// findLastMessage trouve le dernier message d'un utilisateur dans un tableau de messages.
func findLastMessage(messages []models.Chat, userID int) models.Chat {
	var lastMessage models.Chat
	for _, message := range messages {
		if message.UserSender == userID || message.UserReceiver == userID {
			if strings.Compare(lastMessage.Date, message.Date) < 0 {
				lastMessage = message
			}
		}
	}
	return lastMessage
}

// InsertMessage insère un nouveau message dans la base de données.
func InsertMessage(db *sql.DB, userSender, userReceiver int, messageContent, date string) error {
	// Préparation de la requête SQL pour insérer un nouveau message.
	query := `INSERT INTO messages (user_id_sender, user_id_receiver, message_content, date) VALUES (?, ?, ?, ?);`
	// Exécution de la requête avec les valeurs fournies.
	_, err := db.Exec(query, userSender, userReceiver, messageContent, date)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion du message:", err)
		return err
	}

	return nil
}
