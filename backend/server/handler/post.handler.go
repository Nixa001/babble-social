package handler

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
)

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	// --------retrieving form values ----------
	cors.SetCors(&w)
	fmt.Println("r is ", r)
	fmt.Println("w is ", w)

	fmt.Println("--------------------------------------------")
	fmt.Println("          Post Form values                 ")
	fmt.Println("--------------------------------------------")
	if r.Method == "POST" {
		PostContent := r.FormValue("content")
		fmt.Println("[INFO] post content: ", PostContent) //debug

		Sport := r.FormValue("sport")
		Health := r.FormValue("health")
		Music := r.FormValue("music")
		News := r.FormValue("news")
		Others := r.FormValue("others")
		Techno := r.FormValue("techno")
		categorie := []string{Health, Sport, News, Techno, Others, Music}
		var sortCat []string
		for _, v := range categorie {
			if v != "" {
				sortCat = append(sortCat, v)
			}
		}
		categorie = sortCat
		fmt.Println("[INFO] categories: ", categorie) //debug

		Privacy := r.FormValue("privacy")
		fmt.Println("[INFO] privacy: ", Privacy) //debug

		Viewers := r.FormValue("viewers")
		fmt.Println("[INFO] viewers: ", Viewers) //debug

		Image, _ := utils.Uploader(w, r, 20, "image", "")
		fmt.Println("[INFO] imagelink: ", Image) //debug
		PostToCreate := &models.Post{
			ToIns: models.InsPost{
				Content:  PostContent,
				Media:    Image,
				User_id:  1,
				Group_id: 0,
				Privacy:  Privacy,
			},
			Categories: categorie,
			Viewers:    Viewers,
		}
		fmt.Println(PostToCreate)
		notOk, err := service.PostServ.CreatePost(PostToCreate)
		if !notOk {
			log.Println("problem after create service")
			return
		}
	}
}
