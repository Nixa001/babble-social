package groups

import (
	"backend/models"
	"backend/server/cors"
	utils "backend/utils"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

var group models.Group

var UserId int = 1

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	var db = seed.CreateDB()
	defer db.Close()
	var group models.Group = parseFormData(w, r)

	if group.Name == "" {
		http.Error(w, "Le champ 'Name' est obligatoire", http.StatusBadRequest)
		return
	}

	isExist, err := checkGroupExists(db, group.Name)
	if err != nil {
		fmt.Println("Erreur lors de la vérification du groupe:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	if isExist {
		response := map[string]string{"message": "Le groupe existe déjà"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	insertGroupCreated(db, group)

	response := map[string]string{"message": "Groupe créé"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func insertGroupCreated(db *sql.DB, group models.Group) {
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
func parseFormData(w http.ResponseWriter, r *http.Request) models.Group {
	var avatarGroup = "profil_group_avatar.jpg"
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Impossible de lire le form", http.StatusBadRequest)
		return models.Group{}
	}

	group := models.Group{
		Name:           r.FormValue("name"),
		Description:    r.FormValue("description"),
		ID_User_Create: UserId,
		Creation_Date:  utils.GetCurrentDateTime(),
		Avatar:         avatarGroup,
	}
	// fmt.Println(group)

	// var imageFile multipart.File
	Image, _ := utils.Uploader(w, r, 20, "image", "")
	group.Avatar = utils.FormatImgLink(Image)

	// var handler *multipart.FileHeader
	// if r.MultipartForm != nil && r.MultipartForm.File != nil {
	// 	file, fileHeader, err := r.FormFile("image")
	// 	if err != nil {
	// 		fmt.Println("Image par defaut")
	// 		return group
	// 	}
	// 	defer file.Close()
	// 	imageFile = file
	// 	handler = fileHeader
	// }

	// if imageFile != nil {
	// uploadsDir := "./uploads"
	// ext := filepath.Ext(handler.Filename)
	// newFileName := fmt.Sprintf("profil_group_%s%s", group.Name, ext)
	// avatarGroup = newFileName
	// if utils.IsValidImageType(newFileName) {
	// 	uploadsPath := filepath.Join(uploadsDir, newFileName)
	// 	newFile, err := os.Create(uploadsPath)
	// 	if err != nil {
	// 		http.Error(w, "Impossible de creer l image", http.StatusInternalServerError)
	// 		return group
	// 	}
	// 	defer newFile.Close()

	// 	_, err = io.Copy(newFile, imageFile)
	// 	if err != nil {
	// 		http.Error(w, "Impossible de copier l image", http.StatusInternalServerError)
	// 		return group
	// 	}
	// 	group.Avatar = avatarGroup
	// }
	// }
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
