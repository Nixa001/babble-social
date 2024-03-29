package groups

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
)

func PostGroupHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	switch r.Method {
	case "POST":
		content := r.FormValue("content")
		groupID, err := GetGroupIDFromRequest(w, r)
		if err != nil {
			return
		}
		
		categories := getCategoriesFromForm(r)
		privacy := r.FormValue("privacy")
		viewers := fmt.Sprintf("%s, %s", "1", r.FormValue("viewers"))
		imageLink, err := utils.Uploader(w, r, 20, "image", "")
		if err != nil {
			fmt.Println("Upload img")
			return
		}

		post := models.Post{
			ToIns: models.InsPost{
				Content:  content,
				Media:    utils.FormatImgLink(imageLink),
				User_id:  1,
				Group_id: groupID,
				Privacy:  privacy,
			},
			Categories: categories,
			Viewers:    viewers,
		}

		notOk, err1 := service.PostServ.CreatePost(post)
		if notOk {
			log.Println("problem after create service ", err)
			utils.Alert(w, err1)
			return
		} else {
			msg := models.Errormessage{
				Type:       "success",
				Msg:        "post created successfully",
				StatusCode: 200,
				Display:    false,
			}
			utils.Alert(w, msg)
		}
	}
}

func getCategoriesFromForm(r *http.Request) []string {
	categories := []string{"Health", "Sport", "News", "Techno", "Others", "Music"}
	var selectedCategories []string
	for _, cat := range categories {
		if r.FormValue(cat) != "" {
			selectedCategories = append(selectedCategories, cat)
		}
	}
	return selectedCategories
}
