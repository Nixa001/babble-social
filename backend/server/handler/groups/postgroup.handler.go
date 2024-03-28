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
	// --------retrieving form values ----------
	idPost := "1"
	cors.SetCors(&w)

	switch r.Method {
	case "POST":
		log.Println("--------------------------------------------")
		log.Println("          Post Form values                 ")
		log.Println("--------------------------------------------")
		PostContent := r.FormValue("content")
		log.Println("[INFO] post content: ", PostContent) //debug

		Sport := r.FormValue("Sport")
		Health := r.FormValue("Health")
		Music := r.FormValue("Music")
		News := r.FormValue("News")
		Others := r.FormValue("Others")
		Techno := r.FormValue("Tech")
		categorie := []string{Health, Sport, News, Techno, Others, Music}
		var sortCat []string
		for _, v := range categorie {
			if v != "" {
				sortCat = append(sortCat, v)
			}
		}
		categorie = sortCat
		log.Println("[INFO] categories: ", categorie) //debug

		Privacy := r.FormValue("privacy")
		log.Println("[INFO] privacy: ", Privacy) //debug

		Viewers := fmt.Sprintf("%s, %s", idPost, r.FormValue("viewers"))
		log.Println("[INFO] viewers: ", Viewers) //debug

		Image, _ := utils.Uploader(w, r, 20, "image", "")
		log.Println("[INFO] imagelink: ", Image) //debug
		PostToCreate := models.Post{
			ToIns: models.InsPost{
				Content:  PostContent,
				Media:    utils.FormatImgLink(Image),
				User_id:  1,
				Group_id: 0,
				Privacy:  Privacy,
			},
			Categories: categorie,
			Viewers:    Viewers,
		}
		log.Println(PostToCreate)
		notOk, err := service.PostServ.CreatePost(PostToCreate)
		if notOk {
			log.Println("problem after create service ", err)
			utils.Alert(w, err)
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
