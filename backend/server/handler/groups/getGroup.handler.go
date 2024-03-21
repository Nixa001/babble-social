package groups

import (
	"backend/models"
	"backend/server/cors"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Post struct {
	ID       int              `json:"id"`
	Content  string           `json:"content"`
	Media    string           `json:"media"`
	Date     string           `json:"date"`
	User_id  int              `json:"userId"`
	FullName string           `json:"fullname"`
	Username string           `json:"username"`
	User     models.User      `json:"user"`
	Group_id int              `json:"groupId"`
	Privacy  string           `json:"privacy"`
	Comments []models.Comment `json:"comments"`
	Likes    int              `json:"likes"`
	Dislikes int              `json:"dislikes"`
}

type ResponseGroup struct {
	GroupData models.Group  `json:"group_data"`
	Posts     []Post        `json:"posts"`
	Members   []models.User `json:"members"`
	Followers []models.User `json:"followers"`
}

const userId int = 1

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
	// fmt.Println(groupId)

	allPosts, err := getPosts(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	members, err := getMembers(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	groupData, err := getInfoGroup(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}
	followers, err := getFollowers(db, userId)
	if err != nil {
		fmt.Println(err)
		return
	}
	responseGroup := ResponseGroup{
		GroupData: groupData,
		Posts:     allPosts,
		Members:   members,
		Followers: followers,
	}
	jsonData, err := json.Marshal(responseGroup)
	if err != nil {
		fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
func GetUserData(db *sql.DB, userID int) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id =?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return user, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
		if err != nil {
			return user, fmt.Errorf("err scan group data: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return user, fmt.Errorf("err read all groups results: %w", err)
	}

	// fmt.Println(user)
	return user, nil
}

func getLikesDislikes(db *sql.DB, postID int) (int, int) {
	var likes, dislikes int
	err := db.QueryRow("SELECT COUNT(*) FROM postReact WHERE post_id = ? AND reaction = 1", postID).Scan(&likes)
	if err != nil {
		log.Printf("Error fetching likes: %v\n", err)
	}

	err = db.QueryRow("SELECT COUNT(*) FROM postReact WHERE post_id = ? AND reaction = 0", postID).Scan(&dislikes)
	if err != nil {
		log.Printf("Error fetching dislikes: %v\n", err)
	}

	return likes, dislikes
}

func getComments(db *sql.DB, postID int) ([]models.Comment, error) {
	var comments []models.Comment

	query := "SELECT * FROM comment WHERE post_id = ?"
	rows, err := db.Query(query, postID)
	if err != nil {
		return comments, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Date, &comment.Media, &comment.Post_id, &comment.User_id)
		if err != nil {
			return comments, fmt.Errorf("%w", err)
		}
		comment.User, err = GetUserData(db, comment.User_id)
		if err != nil {
			return comments, fmt.Errorf("%w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func getPosts(db *sql.DB, groupID int) ([]Post, error) {
	var allPosts []Post

	query := "SELECT id, content, media, date, user_id, privacy FROM posts WHERE group_id =? ORDER BY id DESC"
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("err exec all post query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Content, &post.Media, &post.Date, &post.User_id, &post.Privacy)
		if err != nil {
			return nil, fmt.Errorf("err group data: %w", err)
		}
		user, err := GetUserData(db, post.User_id)
		if err != nil {
			return nil, fmt.Errorf("err GetUserData data: %w", err)
		}
		post.Comments, err = getComments(db, post.ID)
		if err != nil {
			return nil, fmt.Errorf("err GetUserData data: %w", err)
		}

		// fmt.Println(post.Comments)
		post.FullName = user.Firstname + " " + user.Lastname
		post.Username = user.Username
		post.User = user

		post.Likes, post.Dislikes = getLikesDislikes(db, post.ID)
		allPosts = append(allPosts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("err read all groups results: %w", err)
	}

	return allPosts, nil
}

func getInfoGroup(db *sql.DB, groupId int) (models.Group, error) {
	var group models.Group
	// db.QueryRow("SELECT name, description, id_user_create FROM groups WHERE id = ?", groupId).Scan(&name, &id_user_create)

	query := "SELECT * FROM groups WHERE id = ?"
	rows, err := db.Query(query, groupId)
	if err != nil {
		return models.Group{}, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			return models.Group{}, fmt.Errorf("err scan group data: %w", err)
		}
	}

	group.User, err = GetUserData(db, group.ID_User_Create)
	if err != nil {
		return models.Group{}, fmt.Errorf("err get userData Group: %w", err)
	}
	return group, nil
}

func getMembers(db *sql.DB, groupID int) ([]models.User, error) {
	var members []models.User

	query := "SELECT user_id FROM group_followers WHERE group_id = ?"
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		user, err := GetUserData(db, id)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		members = append(members, user)
	}

	return members, nil
}

func getFollowers(db *sql.DB, userID int) ([]models.User, error) {
	var followers []models.User

	query := "SELECT user_id_follower FROM users_followers WHERE user_id_followed = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		user, err := GetUserData(db, id)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		followers = append(followers, user)
	}

	return followers, nil
}
