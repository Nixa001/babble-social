package groups

import (
	"backend/models"
	"backend/server/cors"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Post struct {
	ID       int
	Content  string `json:"content"`
	Media    string `json:"media"`
	Date     string `json:"date"`
	User_id  int    `json:"userId"`
	Group_id int    `json:"groupId"`
	Privacy  string `json:"privacy"`
}

type ResponseGroup struct {
	GroupData models.Group `json:"group_data"`
	Posts     []Post       `json:"posts"`
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()

	parsedURL, err := url.Parse(r.URL.String())
	if err != nil {
		fmt.Fprintf(w, "Error parsing URL: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	values := parsedURL.Query()
	id := values.Get("id")
	groupId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid Id")
		return
	}

	if id == "" {
		fmt.Fprintf(w, "No 'id' parameter found in the URL")
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}
	fmt.Println(groupId)

	allPosts, err := getPosts(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allPosts, len(allPosts))

	// jsonData, err := json.Marshal(group)
	// if err != nil {
	// 	fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	groupData, err := getInfoGroup(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseGroup := ResponseGroup{
		GroupData: groupData,
		Posts:     allPosts,
	}
	jsonData, err := json.Marshal(responseGroup)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

func getPosts(db *sql.DB, groupID int) ([]Post, error) {
	var allPosts []Post

	query := "SELECT id, content, media ,date, user_id ,privacy FROM posts WHERE group_id =?"
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute all post query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Content, &post.Media, &post.Date, &post.User_id, &post.Privacy)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group data: %w", err)
		}
		allPosts = append(allPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read all groups results: %w", err)
	}

	return allPosts, nil
}

func getInfoGroup(db *sql.DB, groupId int) (models.Group, error) {
	var group models.Group
	// db.QueryRow("SELECT name, description, id_user_create FROM groups WHERE id = ?", groupId).Scan(&name, &id_user_create)

	query := "SELECT * FROM groups WHERE id = ?"
	rows, err := db.Query(query, groupId)
	if err != nil {
		return models.Group{}, fmt.Errorf("failed to execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			return models.Group{}, fmt.Errorf("failed to scan group data: %w", err)
		}
	}
	return group, nil
}

func getAllMembers(db *sql.DB) ([]models.Group, error) {
	var allGroups []models.Group

	query := "SELECT * FROM groups"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group data: %w", err)
		}
		allGroups = append(allGroups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read all groups results: %w", err)
	}

	return allGroups, nil
}
