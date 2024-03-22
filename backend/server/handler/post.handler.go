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
	idUser := "1"
	cors.SetCors(&w)
	log.Println("--------------------------------------------")
	log.Println("          Post Form values                 ")
	log.Println("--------------------------------------------")
	switch r.Method {
	case "POST":
		PostContent := r.FormValue("content")
		log.Println("[INFO] post content: ", PostContent) //debug

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
		log.Println("[INFO] categories: ", categorie) //debug

		Privacy := r.FormValue("privacy")
		log.Println("[INFO] privacy: ", Privacy) //debug

		Viewers := fmt.Sprintf("%s, %s", idUser, r.FormValue("viewers"))
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
	case "GET":
		log.Println("in get method")
		postTab, err := service.PostServ.GetPost(1)
		if err != nil {
			log.Println("problem after get service ", err)
			msg := models.Errormessage{
				Type:       "Internal Servor Error",
				Msg:        "Oops something wrong occured !",
				StatusCode: 500,
				Display:    true,
			}
			utils.Alert(w, msg)
			return
		} else {
			log.Println("Gotten => ", postTab)
			//log.Println("Gotten top => ", postTab[0])
			utils.AlertPostData(w, models.WResponse{
				Type:       "loadPost",
				Data:       postTab,
				StatusCode: 200,
				Display:    false,
				Msg:        "posts retrieved succesfully",
			})
		}
	default:
		msg := models.Errormessage{
			Type:       "Not allowed",
			Msg:        "Method not allowed",
			StatusCode: 401,
			Display:    true,
		}
		utils.Alert(w, msg)
	}
}
