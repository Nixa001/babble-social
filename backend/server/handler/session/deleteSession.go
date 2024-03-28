package session

import (
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "Cookie de session non trouvé", http.StatusNotFound)
		return
	}
	// fmt.Printf("Suppression de la session avec l'ID : %s\n", sessionCookie.Value)
	stm := `DELETE FROM sessions WHERE id = ?`
	req, err := seed.DB.Prepare(stm)
	if err != nil {
		http.Error(w, "Erreur lors de la préparation de la requête de suppression", http.StatusInternalServerError)
		return
	}
	defer req.Close()
	_, err = req.Exec(sessionCookie.Value)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression de la session de la base de données", http.StatusInternalServerError)
		return
	}
	var reponse = map[string]interface{}{
		"logout":  true,
		"message": "connect successfully",
	}
	responsJson, err := json.Marshal(reponse)
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Write(responsJson)
}
