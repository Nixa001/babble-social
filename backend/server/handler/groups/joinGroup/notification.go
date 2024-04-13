package joingroup

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func RecupeIdAdminGroup(idGroup int, db *sql.DB) (int, error) {
	req := `
		SELECT id_user_create FROM groups WHERE id = ?
	`

	stm, err := db.Prepare(req)
	if err != nil {
		return 0, err
	}

	defer stm.Close()

	var id_user_greate_group int

	err = stm.QueryRow(idGroup).Scan(&id_user_greate_group)
	if err != nil {
		return 0, err
	}

	fmt.Println("id user = ", id_user_greate_group)

	return id_user_greate_group, nil
}

func InsertNotification(idGroup int, notification_type string, user_id_sender int, db *sql.DB) error {
	id_user_created_group, err := RecupeIdAdminGroup(idGroup, db)
	if err != nil {
		log.Fatal("Erreur lors de la recuperation de l'id de l'admin group ", err)
	}
	// a determiner au niveau de la session

	fmt.Println("type = ", notification_type)
	fmt.Println("idGroup = ", idGroup)
	fmt.Println("user_id_sender =", user_id_sender)
	checkNotif, _ := CheckNotifAndType(db, idGroup, user_id_sender, notification_type)
	fmt.Println("checkNotif = ", checkNotif)
	if !checkNotif {

		req := `
			INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, id_group, date) VALUES ($1, $2, $3, $4, $5, $6)
		`
		stm, err := db.Prepare(req)

		if err != nil {
			return err
		}

		defer stm.Close()

		if err != nil {
			fmt.Println("Erreur lors de la recuperation de l'id_user_created_group", err)
			return err
		}

		date := time.Now()
		formattedDateTime := date.Format("2006-01-02 15:04:05")
		_, err = stm.Exec(notification_type, 0, user_id_sender, id_user_created_group, idGroup, formattedDateTime)
		if err != nil {
			fmt.Println("Erreur lors de l'execution de la requete inserte dans la base ", err)
			return err
		}

		// fmt.Println("Notification insertion success")
	} else {
		fmt.Println("Vous avez une demande en cours")
	}
	return nil
}

func CheckJoinNotification(id_user_created_group int, id_user_connected int, idGroup int, db *sql.DB) (int, int) {
	req := `
        SELECT id, status FROM notifications WHERE user_id_sender = ? AND user_id_receiver = ? AND id_group = ?
    `

	stm, err := db.Prepare(req)
	if err != nil {
		fmt.Println("Error preparing request checkJoinNotification: ", err)
		return 0, 0
	}

	defer stm.Close()

	var id_notification int
	var state int

	err = stm.QueryRow(id_user_connected, id_user_created_group, idGroup).Scan(&id_notification, &state)
	if err != nil {
		fmt.Println("Error querying checkJoinNotification: ", err)
		return 0, 0
	}

	fmt.Println("Il exsit dejat une demande de rejoindre ce groupe")
	fmt.Println("Id notification ", id_notification)
	return id_notification, state
}

func CheckNotifAndType(db *sql.DB, groupID int, userID int, notificationType string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notifications WHERE notification_type = $1 AND user_id_sender = $2 AND id_group = $3", notificationType, userID, groupID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AcceptOrNo(db *sql.DB, user_id_sender, user_id_receiver, id_group int, val string) bool {
	stm := `
		UPDATE notifications SET status = ? WHERE user_id_sender = ? AND user_id_receiver = ? AND id_group = ? 
	`
	// Prépare la requête à partir de la connexion à la base de données
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête: ", err)
		return false
	}
	defer query.Close() // Assurez-vous de fermer la requête préparée après son utilisation

	// Exécute la requête préparée avec les paramètres spécifiques
	_, err = query.Exec(val, user_id_sender, user_id_receiver, id_group) // '1' est la nouvelle valeur de 'status'
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête: ", err)
		return false
	}

	return true
}

func InsertGroupFollowers(db *sql.DB, user_id, group_id int) {
	stm := `
		INSERT INTO group_followers VALUES (?, ?)
	`
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors du prepation de la requete ", err)
		return
	}
	defer query.Close()

	_, err = query.Exec(user_id, group_id)
	if err != nil {
		fmt.Println("Erreur lors de l'execution de la requete ", err)
	}
}
