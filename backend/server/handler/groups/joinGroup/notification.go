package joingroup

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

func InsertNotification(idGroup int, db *sql.DB) error {
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
	id_user_connected := 1

	checkNotif, _ := CheckNotif(db, idGroup, id_user_connected)
	if checkNotif {

		check, _ := CheckJoinNotification(id_user_created_group, id_user_connected, idGroup, db)
		if check == 0 {

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

			_, err = stm.Exec("Join_group", 0, id_user_connected, id_user_created_group, idGroup)
			if err != nil {
				fmt.Println("Erreur lors de l'execution de la requete inserte dans la base ", err)
				return err
			}
		} else {
			fmt.Println("Vous avez une demande en cours")
			return err
		}
	}
	return nil
}

func CheckJoinNotification(id_user_created_group int, id_user_connected int, idGroup int, db *sql.DB) (int, int) {
	req := `
        SELECT id, status FROM notifications WHERE user_id_sender = $1 AND user_id_receiver = $2 AND id_group = $3
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
		// fmt.Println("Error querying checkJoinNotification: ", err)
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

func InsertFollowNotification(followerId, followedId int, db *sql.DB) error {
	// Insertion de la notification
	fmt.Println("Insertion de la notification")
	checkNotif, err := CheckNotifExist(db, followerId, followedId)
	if err != nil {
		fmt.Println("Erreur checked ", err.Error())
		return err
	}
	// fmt.Println("====", checkNotif)
	if !checkNotif {
		fmt.Println("CheckNotifExist ")
		check, _ := CheckJoinFollowNotification(followerId, followedId, db)
		if check == 0 {

			req := `
			INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, date) VALUES ($1, $2, $3, $4, $5)
		`
			stm, err := db.Prepare(req)

			if err != nil {
				fmt.Println(err)
				return err
			}

			defer stm.Close()
			date := time.Now()
			_, err = stm.Exec("follow", 0, followerId, followedId, date)
			if err != nil {
				fmt.Println("Erreur lors de l'execution de la requete inserte dans la base ", err)
				return err
			}
		} else {
			fmt.Println("Vous avez une demande en cours")
			return err
		}
	}
	return nil
}

func CheckJoinFollowNotification(followerId, followedId int, db *sql.DB) (int, int) {
	fmt.Printf("FollwerId = %v, FoolowedId = %v", followerId, followedId)

	req := `
        SELECT id, status FROM notifications WHERE user_id_sender = $1 AND user_id_receiver = $2
    `

	stm, err := db.Prepare(req)
	if err != nil {
		fmt.Println("Error preparing request checkJoinNotification: ", err)
		return 0, 0
	}

	defer stm.Close()

	var id_notification int
	var state int

	err = stm.QueryRow(followerId, followedId).Scan(&id_notification, &state)
	if err != nil {
		fmt.Println("Error querying checkJoinNotification: ", err)
		return 0, 0
	}

	fmt.Println("Il existe dejà une demande de follow")
	fmt.Println("Id notification ", id_notification)
	return id_notification, state
}

func CheckNotifExist(db *sql.DB, followerId, followedId int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id_sender = $1 AND user_id_receiver = $2 AND notification_type = $3", followerId, followedId, "follow").Scan(&count)
	if err != nil {
		fmt.Println("Erreur requete ", err)
		return false, err
	}
	fmt.Println("Count = ", count)
	return count > 0, nil
}
