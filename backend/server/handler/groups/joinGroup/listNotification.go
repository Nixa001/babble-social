package joingroup

import (
	"backend/database"
	"backend/models"
	"backend/server/service"
	"fmt"
	"net/http"
)

func ListNotification(r *http.Request) []models.Notification {
	id_user, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		fmt.Println("Error verifying ", err.Error())
		return nil
	}
	db := database.NewDatabase()
	fmt.Println("-------------------------", id_user.User_id)
	stm := `
		SELECT * FROM notifications WHERE user_id_receiver = ? AND status = ?
	`
	req, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Error preparing notifications: ", err.Error())
		return nil
	}
	defer req.Close()
	rows, err := req.Query(id_user.User_id, 0)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return nil
	}
	defer rows.Close()
	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.ID, &notification.Type, &notification.Status, &notification.UserIDSender, &notification.UserIDReceveived, &notification.GroupId, &notification.Date)
		if err != nil {
			fmt.Println("Error scanning row: ", err.Error())
			return nil
		}
		notifications = append(notifications, notification)
	}
	fmt.Println("List notification ======", notifications)
	return notifications
}
