package handler

import (
	"backend/server/cors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Impossible de lire le formulaire", http.StatusBadRequest)
		return
	}

	group := Group{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}
	fmt.Println(group)

	imageFile, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Impossible de lire l image", http.StatusBadRequest)
		return
	}
	fmt.Println(imageFile)
	defer imageFile.Close()
	fmt.Println(imageFile)

	fileName := handler.Filename
	ext := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("%s%s", group.Name, ext)
	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Impossible de creer l image", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, imageFile)
	if err != nil {
		http.Error(w, "Impossible de copier l image", http.StatusInternalServerError)
		return
	}


	response := map[string]string{"message": "Groupe cree"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
