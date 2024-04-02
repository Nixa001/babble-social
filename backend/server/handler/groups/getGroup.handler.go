package groups

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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
	GroupData   models.Group         `json:"group_data"`
	Posts       []models.DataPost    `json:"posts"`
	Members     []models.User        `json:"members"`
	Followers   []models.User        `json:"followers"`
	Event       []models.Event       `json:"events"`
	EventJoined []models.EventJoined `json:"events_joined"`
}

const userId int = 1

func GetGroup(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()

	groupId, err := GetGroupIDFromRequest(w, r)

	allPosts, err := service.PostServ.GetPost(groupId)
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
	events, eventJoined, err := GetEvent(db, groupId, userId)
	// fmt.Println("eventJoined: ", eventJoined)

	responseGroup := ResponseGroup{
		GroupData:   groupData,
		Posts:       allPosts,
		Members:     members,
		Followers:   followers,
		Event:       events,
		EventJoined: eventJoined,
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

func GetGroupIDFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "No id in URL")
		http.Error(w, "no ID", http.StatusBadRequest)
		return 0, fmt.Errorf("no group ID")
	}

	groupID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid Id")
		return 0, fmt.Errorf("invalid group ID")
	}

	return groupID, nil
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
		err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
		if err != nil {
			return user, fmt.Errorf("err scan group data: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return user, fmt.Errorf("err read all groups results: %w", err)
	}
	return user, nil
}
func GetEventJoined(db *sql.DB, userID int) ([]int, error) {
	var events []int

	query := "SELECT event_id FROM event_joined WHERE user_id =?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return events, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event int
		err := rows.Scan(&event)
		if err != nil {
			return events, fmt.Errorf("err scan group data: %w", err)
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return events, fmt.Errorf("err read all groups results: %w", err)
	}

	return events, nil
}

func GetEvent(db *sql.DB, groupID, userID int) ([]models.Event, []models.EventJoined, error) {
	var events []models.Event
	var eventsJoined []models.Event
	events_joined_id, err := GetEventJoined(db, userID)
	if err != nil {
		fmt.Println("error on GetEventJoined", err)
	}
	
	query := "SELECT * FROM event WHERE group_id =?"
	rows, err := db.Query(query, groupID)
	if err != nil {
		return events, nil, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.GroupID, &event.UserID, &event.Description, &event.Date)
		if err != nil {
			return events, nil, fmt.Errorf("err scan group data: %w", err)
		}
		event.Date = formatDateTimeFr(event.Date)
		if contains(events_joined_id, event.ID) {
			eventsJoined = append(eventsJoined, event)
		} else {
			events = append(events, event)
		}
	}

	if err := rows.Err(); err != nil {
		return events, nil, fmt.Errorf("err read all groups results: %w", err)
	}

	// query = "SELECT * FROM event_joined WHERE group_id =?"
	query = "SELECT event_joined.id, event_joined.event_id, event_joined.user_id, event_joined.group_id, event.description, event.event_date FROM event_joined INNER JOIN event ON event_joined.event_id = event.id WHERE event_joined.group_id = ?"

	rows, err = db.Query(query, groupID)
	if err != nil {
		return events, nil, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	var dataEventJoint []models.EventJoined
	for rows.Next() {
		var event models.EventJoined
		err := rows.Scan(&event.ID, &event.Event_id, &event.User_id, &event.Group_id, &event.Description, &event.Date)
		if err != nil {
			fmt.Println("error scanning event ", err)
			log.Fatal(err.Error())
		}
		dataEventJoint = append(dataEventJoint, event)
		
		// fmt.Println("dataEventJoint ", event)
	}

	return events, dataEventJoint, nil
}

func contains(arr []int, value int) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
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
		// comment.User, err = GetUserData(db, comment.User_id)
		// if err != nil {
		// 	return comments, fmt.Errorf("%w", err)
		// }
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
		post.FullName = user.First_name + " " + user.Last_name
		post.Username = user.User_name
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

func formatDateTimeFr(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	weekday := t.Weekday().String()[:3]
	day := t.Day()
	month := (t.Month().String()[:3])
	year := t.Year()
	hour := t.Hour()

	return fmt.Sprintf("%s %d %s %d at %dH", weekday, day, month, year, hour)
}
