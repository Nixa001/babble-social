package joingroup

import (
	"backend/database"
	"backend/models"
	"backend/server/service"
	"database/sql"
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
		var groupId sql.NullInt64
		err := rows.Scan(&notification.ID, &notification.Type, &notification.Status, &notification.UserIDSender, &notification.UserIDReceveived, &groupId, &notification.Date)
		notification.GroupId = getIntValue(groupId)
		if err != nil {
			fmt.Println("Error scanning row: ", err.Error())
			return nil
		}
		notifications = append(notifications, notification)
	}
	fmt.Println("List notification ======", notifications)
	return notifications
}

func getIntValue(value sql.NullInt64) int {
	if value.Valid {
		return int(value.Int64)
	}
	return 0
}
