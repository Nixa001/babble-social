package groups

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	utils "backend/utils"
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
	GroupData   models.Group      `json:"group_data"`
	Posts       []models.DataPost `json:"posts"`
	Members     []models.User     `json:"members"`
	Followers   []models.User     `json:"followers"`
	Event       []models.Event    `json:"events"`
	EventJoined []models.Event    `json:"events_joined"`
}

// const userId int = 1

func GetGroup(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()

	groupId, err := GetGroupIDFromRequest(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		log.Println("Invalid Token ", err)
		utils.Alert(w, models.Errormessage{
			Type:       "Get group",
			Msg:        "Invalid Token",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	userId := session.User_id
	allPosts, err := service.PostServ.FetchPostGroup(groupId)
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
	if err != nil {

		fmt.Fprintf(w, "Error on getting event")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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
func GetEventJoinedID(db *sql.DB, userID, groupID int) ([]int, error) {
	var events []int

	query := "SELECT event_id FROM event_joined WHERE user_id =? AND group_id =?"
	rows, err := db.Query(query, userID, groupID)
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

	return events, nil
}

func GetEvent(db *sql.DB, groupID, userID int) ([]models.Event, []models.Event, error) {
	var events []models.Event
	var eventsJoined []models.Event
	events_joined_id, err := GetEventJoinedID(db, userID, groupID)
	if err != nil {
		fmt.Println("error on GetEventJoined", err)
	}

	query := `
				SELECT *
				FROM event
				WHERE event.group_id = ? AND event.id NOT IN (
				SELECT event_notjoined.event_id
				FROM event_notjoined );
				`
	rows, err := db.Query(query, groupID)
	if err != nil {
		return events, nil, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.GroupID, &event.UserID, &event.Description, &event.Date, &event.Is_joined)
		if err != nil {
			return events, nil, fmt.Errorf("err scan group data: %w", err)
		}
		event.Date = formatDateTimeFr(event.Date)
		if contains(events_joined_id, event.ID) {
			eventsJoined = append(eventsJoined, event)
		} else {
			events = append(events, event)
		}
		// fmt.Println("event ", events)
	}

	return events, eventsJoined, nil
}

func contains(arr []int, value int) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
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
