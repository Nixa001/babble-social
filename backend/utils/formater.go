package utils

import (
	"backend/models"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Time() (string, string) {
	current := time.Now()
	state := "am"

	// check wether the time is past or after morning
	if current.Hour() >= 12 {
		state = "pm"
	}

	date := time.Now().Format("Jan 2, 2006")
	hour := time.Now().Format("03:04" + " " + state)
	return date, hour
}

func FormatCategory(categories []string, postID string) []models.Category {
	var tab []models.Category
	for _, v := range categories {
		tab = append(tab, models.Category{Post_id: postID, Category: v})
	}
	return tab
}

func FormatViewers(viewers string, postID string) []models.Viewers {
	var tab []models.Viewers
	for _, v := range strings.Split(viewers, ",") {
		id, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error while converting viewer id")
			return nil
		}
		tab = append(tab, models.Viewers{Post_id: postID, User_id: id})
	}
	return tab
}
