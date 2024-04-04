package joingroup

import (
	"database/sql"
	"fmt"
	"log"
)

func recupeIdAdminGroup(idGroup int, db *sql.DB) (int, error) {
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
	fmt.Println("Mangi inserer")
	// Insertion de la notification
	/*
		// Req pour supprimer les donnees de testes
		query := "DELETE FROM notifications WHERE id = ?"
		result, err := db.Exec(query, 2)
		if err != nil {
			// Gérer l'erreur ici, par exemple :
			fmt.Println("Erreur lors de l'exécution de la requête DELETE:", err)
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			// Gérer l'erreur ici, par exemple :
			fmt.Println("Erreur lors de la récupération du nombre de lignes affectées:", err)
			return err
		}

		fmt.Println("Nombre de lignes supprimées:", rowsAffected)
	*/

	id_user_created_group, err := recupeIdAdminGroup(idGroup, db)
	if err != nil {
		log.Fatal("Erreur lors de la recuperation de l'id de l'admin group ", err)
	}
	// a determiner au niveau de la session

	checkNotif, _ := CheckNotifAndType(db, idGroup, user_id_sender, notification_type)
	fmt.Println("Checking ", checkNotif)
	fmt.Println("GroupId = ", idGroup)
	// log.Fatal("err")
	if !checkNotif {

		req := `
			INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, id_group) VALUES ($1, $2, $3, $4, $5)
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

		fmt.Printf("Notification type: %v , userId : %v, sender: %v, idGroup: %v\n", notification_type, user_id_sender, id_user_created_group, idGroup)
		fmt.Println("Notification ", notification_type)
		_, err = stm.Exec(notification_type, 0, user_id_sender, id_user_created_group, idGroup)
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

func CheckNotif(db *sql.DB, groupID, userID int) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT id FROM notifications WHERE user_id_sender = $1 AND id_group = $2)", userID, groupID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func CheckNotifType(db *sql.DB, groupID, userID int, notification_type string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT notification_type FROM notifications WHERE user_id_sender = $1 AND id_group = $2)", userID, groupID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func CheckNotifAndType(db *sql.DB, groupID int, userID int, notificationType string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notifications WHERE notification_type = $1 AND user_id_sender = $2 AND id_group = $3", notificationType, userID, groupID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
