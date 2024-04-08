package events

import (
	"database/sql"
	"fmt"
)

func JoinEvent(id_user int, id_group int, id_event int, db *sql.DB) error {
	stm := `
		INSERT INTO event_joined (event_id, user_id, group_id) VALUES (?,?,?)
	`
	req, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Error preparing request: ", err)
		return err
	}
	defer req.Close()
	_, err = req.Exec(id_event, id_user, id_group)
	if err != nil {
		fmt.Println("Error executing event: ", err)
		return err
	}
	fmt.Println("Event joined successfully")
	return nil
}
