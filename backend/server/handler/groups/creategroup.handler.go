package groups

import (
	"backend/server/cors"
	utils "backend/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type group struct {
	Name           string
	Description    string
	ID_User_Create int
	Avatar         string
	Creation_Date  string
}

var user_id int = 1

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database/social_network.db")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	cors.SetCors(&w)
	var group group = parseFormData(w, r)
	isExist, err := checkGroupExists(db, group.Name)

	if group.Name == "" {
		fmt.Println("Name input required")
		return
	}

	if !isExist {
		insertGroupCreated(db, group)
	} else {
		fmt.Println("group already exist")

		return
	}

	response := map[string]string{"message": "Groupe cree"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertGroupCreated(db *sql.DB, group group) {
	stmt, err := db.Prepare("INSERT INTO groups(name, description, id_user_create, avatar, creation_date) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête d'insertion:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(group.Name, group.Description, group.ID_User_Create, group.Avatar, group.Creation_Date)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion des données:", err)
		return
	}
}
func parseFormData(w http.ResponseWriter, r *http.Request) group {
	var avatarGroup = "profil_group_avatar.jpg"
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Impossible de lire le form", http.StatusBadRequest)
		return group{}
	}

	group := group{
		Name:           r.FormValue("name"),
		Description:    r.FormValue("description"),
		ID_User_Create: user_id,
		Creation_Date:  utils.GetCurrentDateTime(),
		Avatar:         avatarGroup,
	}

	var imageFile multipart.File
	var handler *multipart.FileHeader
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Image par defaut")
			return group
		}
		defer file.Close()
		imageFile = file
		handler = fileHeader
	}

	if imageFile != nil {
		uploadsDir := "./uploads"
		ext := filepath.Ext(handler.Filename)
		newFileName := fmt.Sprintf("profil_group_%s%s", group.Name, ext)
		avatarGroup = newFileName
		if utils.IsValidImageType(newFileName) {
			uploadsPath := filepath.Join(uploadsDir, newFileName)
			newFile, err := os.Create(uploadsPath)
			if err != nil {
				http.Error(w, "Impossible de creer l image", http.StatusInternalServerError)
				return group
			}
			defer newFile.Close()

			_, err = io.Copy(newFile, imageFile)
			if err != nil {
				http.Error(w, "Impossible de copier l image", http.StatusInternalServerError)
				return group
			}
			group.Avatar = avatarGroup
		}
	}
	return group
}

func checkGroupExists(db *sql.DB, groupName string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM groups WHERE name = ?)", groupName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
