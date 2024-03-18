package groups

import (
	"backend/server/cors"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Group struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ID_User_Create int    `json:"id_user_create"`
	Avatar         string `json:"image"`
	Creation_Date  string `json:"creation_date"`
}

var AllGroups []Group

func GetGroups(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	db, err := sql.Open("sqlite3", "./database/social_network.db")
	if err != nil {

		http.Error(w, "Erreur lors de l'ouverture de la base de données", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var groups []Group
	// var groups_joined []Group

	rows, err := db.Query("SELECT * FROM groups")
	if err != nil {
		fmt.Println("ic")
		http.Error(w, "Erreur lors de l'exécution de la requête", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture des données de groupe", http.StatusInternalServerError)
			return
		}
		groups = append(groups, group)
	}
	fmt.Println(groups)

	if err := rows.Err(); err != nil {
		http.Error(w, "Erreur lors de la lecture des résultats", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}
