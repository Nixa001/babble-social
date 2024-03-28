package session

import (
	"backend/models"
	"backend/server/cors"
	"backend/utils/seed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var session models.Session

func CreateSession(w http.ResponseWriter, r *http.Request, id_user int) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var token = uuid.NewString()
	session.Expiration = time.Now().Add(3 * time.Hour)
	// token = "1234"
	// dataJson, err := json.Marshal(data)

	stm := `
		INSERT INTO sessions (token, user_id, expiration) VALUES(?, ?, ?)
	`
	req, err := seed.DB.Prepare(stm)
	if err != nil {
		fmt.Println("error preparing session", err)
		log.Fatal(err.Error())
	}
	defer req.Close()
	_, err = req.Exec(token, id_user, session.Expiration)
	if err != nil {
		fmt.Println("error executing session", err)
		log.Fatal(err.Error())
	}
	cookie := http.Cookie{
		Name:     "dickss",
		Value:    token,
		Path:     "/",
		Expires:  session.Expiration,
		HttpOnly: true,
	}
	fmt.Println("token:", token)
	fmt.Println("expiration:", session.Expiration)
	http.SetCookie(w, &cookie)
}
