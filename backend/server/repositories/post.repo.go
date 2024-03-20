package repositories

import (
	db "backend/database"
	"backend/models"
	"backend/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/gofrs/uuid"
)

type PostRepository struct {
	BaseRepo
}

const (
	GetPostQuery = `
SELECT 
    p.id AS post_id,
    p.content AS post_content,
    p.media AS post_media,
    p.date AS post_date,
    p.user_id AS post_user_id,
	u.avatar as avatar,
    u.user_name as username,
    concat (u.first_name, " ", u.last_name) as full_name,
    COUNT(c.id) AS comment_count,
	GROUP_CONCAT(DISTINCT cat.category) AS categories
FROM 
    posts AS p
	users AS u
LEFT JOIN 
    comments AS c ON p.id = c.post_id,
	categories AS cat ON p.id = cat.post_id
WHERE 
    (
        p.privacy = 'public'
        OR (
            p.privacy = 'private' AND (
                p.user_id = ? -- Post creator
                OR EXISTS (
                    SELECT 1
                    FROM users_followers
                    WHERE user_id_followed = p.user_id
                    AND user_id_follower = ?
                )
                OR EXISTS (
                    SELECT 1
                    FROM users_followers
                    WHERE user_id_follower = p.user_id
                    AND user_id_followed = ?
                )
            )
        )
        OR (
            p.privacy = 'almost' AND EXISTS (
                SELECT 1
                FROM viewers
                WHERE user_id = ?
                AND post_id = p.id
            )
        )
    )
	GROUP BY p.id, p.content, p.media, p.date, p.user_id
`
	GetOnePostQuery = `
SELECT
    p.id,
    p.content,
    p.media,
    p.date,
    p.user_id,
    u.avatar,
    u.user_name,
    concat (u.first_name, " ", u.last_name),
    COUNT(DISTINCT c.id),
    GROUP_CONCAT(DISTINCT cat.category),
    CASE
    WHEN p.privacy ="public" THEN "public"
    WHEN p.privacy ="private" THEN (
         SELECT GROUP_CONCAT(user_id_follower) 
            FROM users_followers 
            WHERE user_id_followed = p.user_id 
    )
    WHEN p.privacy ="almost" THEN (
        SELECT GROUP_CONCAT(user_id)
            FROM viewers
            WHERE post_id = p.id
        )
        ELSE NULL
    END AS isPublic
FROM
    posts AS p
LEFT JOIN comment AS c ON p.id = c.post_id,
    categories AS cat ON p.id = cat.post_id,
    users AS u ON u.id =  p.user_id
WHERE p.id = ?;
`
)

func (p *PostRepository) init() {
	p.DB = db.DB
	p.TableName = "posts"
}

func (p *PostRepository) CreatePost(post models.Post) (bool, models.Errormessage) {

	//todo: checking Id_user validity

	//checking Title's validity
	if strings.TrimSpace(post.Content) == "" && strings.TrimSpace(post.Media) == "" {
		fmt.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to empty content and media ❌\n", user)
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to empty input",
				StatusCode: models.BRstatus,
				Location:   "home",
				Display:    false,
			}
	}

	//checking category's validity
	if len(post.Categories) < 1 { //user did not select a categorie
		fmt.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to missing category❌\n", user)
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to missing category",
				StatusCode: models.BRstatus,
				Location:   "home",
				Display:    false,
			}
	}

	if len(post.Content) > 1500 { //found only spaces,newlines in the input or chars number limit exceeded
		fmt.Printf("⚠ ERROR ⚠ : Couldn't create post from user %s due to invalid input ❌\n", user)
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create post due to invalid input",
				StatusCode: models.BRstatus,
				Location:   "home",
				Display:    false,
			}
	}

	errIns := p.InsertPost(post)
	if errIns != nil {
		return true,
			models.Errormessage{
				Type:       models.ISEtype,
				Msg:        models.ISEmsg,
				StatusCode: models.ISEstatus,
				Location:   "home",
				Display:    false,
			}
	}
	return false, models.Errormessage{}
}

func (P *PostRepository) InsertPost(post models.Post) error {
	id_post, errp := uuid.NewV4()
	id_image, errImg := uuid.NewV4()
	if errp != nil {
		fmt.Println("❌ Create_post ⚠ ERROR ⚠ : couldn't generate a unique post id")
		return errp
	}
	if errImg != nil {
		fmt.Println("❌ Create_post ⚠ ERROR ⚠ : couldn't generate a unique image id")
		return errp
	}
	date, time := utils.Time() //date and time

	// inserting value in database
	//-- formatting value's special chars
	if post.ToIns.Content != "" {
		post.ToIns.Content = utils.EncodeValue(post.ToIns.Content)
	}
	//-- processing image
	if post.ToIns.Media {
		imgData, err := base64.StdEncoding.DecodeString(post.ToIns.Media)
		if err != nil {
			log.Println("❌ error while decoding image", err)
			return
		}

		err = ioutil.WriteFile(id_image, imgData) // storing in local
		if err != nil {
			log.Println("❌ error while storing image in local:", err)
			return
		}
		post.Media = fmt.Sprintf("http://localhost:8000/images/%s", id_image)
		log.Println("✔ image decoded successfully")
	}

	err := p.DB.Insert(p.TableName, post.ToIns)
	if err != nil {
		fmt.Println("❌ error while inserting post")
		return err
	}
	//inserting categories
	for i := range post.Categories {
		errCat := p.DB.Insert("categories", post.Categories[i])
		if errcat != nil {
			fmt.Println("❌ error while inserting categories")
			return err
		}
	}
	if post.ToIns.Privacy == "almost" {
		for i := range post.Viewers {
			errView := p.DB.Insert("viewers", post.Viewers[i])
			if errView != nil {
				fmt.Println("❌ error while inserting viewers")
				return err
			}
		}
	}
	fmt.Println("✅ post has been created successfully")
	return nil
}

func (p *PostRepository) LoadPost(IdUser int) ([]models.DataPost, error) {
	var postTab []models.DataPost

	rows, err := p.DB.Exec(GetPostQuery, IdUser, IdUser, IdUser, IdUser)
	if err != nil {
		fmt.Println("❌ Error while retrieving posts => ", err)
		return nil, errors.New("error while retrieving posts from the database")
	}
	defer rows.Close()

	for rows.next() {
		var temp models.DataPost
		errScan := rows.scan(&temp.ID, &temp.Content, &temp.Media, &temp.Date, &temp.User_id, &temp.Avatar, &temp.UserName, &temp.FullName, &temp.Comment, &temp.Categories)
		if errScan != nil {
			fmt.Println("⚠ GetPost scan err ⚠ :", errScan)
			return nil, errors.New("error while scanning")
		}

		temp.Content = utils.DecodeValue(temp.Content)
		postTab = append(postTab, temp)
	}
	return postTab, nil
}

func (p *PostRepository) GetOnePost(postID int) (models.DataPost, error) {
	rows, err := p.DB.Exec(GetOnePostQuery, postID)
	if err != nil {
		fmt.Println("❌ Error while retrieving in OnePost => ", err)
		return models.DataPost{}, errors.New("error while retrieving onepost from the database")
	}
	defer rows.Close()

	var data models.DataPost
	//! modify retrieval
	for rows.next() {
		errScan := rows.scan(&data.ID, &data.Content, &data.Media, &data.Date, &data.Avatar, &data.UserName, &data.FullName, &data.Comments, &data.Categories, &data.Viewers)
		if errScan != nil {
			fmt.Println("⚠ GetOnePost scan err ⚠ :", errScan)
			return models.DataPost{}, errors.New("error while scanning")
		}
		data.Content = utils.DecodeValue(data.Content)
	}
	return data, nil
}
