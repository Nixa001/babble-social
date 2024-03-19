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

type CommentRepository struct {
	BaseRepo
}

const GetCommentQuery = `
SELECT DISTINCT
    c.id,
    c.content,
    c.media,
    c.date,
    c.post_id,
    u.avatar,
    u.user_name,
    concat (u.first_name, " ", u.last_name) as full_name
FROM
    comment AS c
    LEFT JOIN users AS u on u.id = c.user_id
WHERE ?
ORDER BY c.id
    DESC;
`

func (c *CommentRepository) init() {
	c.DB = db.DB
	c.TableName = "comment"
}

func (c *CommentRepository) CreateComment(comment models.Comment) (bool, models.Errormessage) {
	/*
		! code here........
		todo: check post's and user's validity
		todo: create media in local
		todo: inssert comment in database
	*/
	//checking content's validity
	if strings.TrimSpace(comment.Content) == "" && strings.TrimSpace(comment.Media) == "" {
		fmt.Printf("⚠ ERROR ⚠ : Couldn't create comment from user %s due to empty content and media ❌\n", user)
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create comment due to empty input",
				StatusCode: models.BRstatus,
				Location:   "home",
				Display:    false,
			}
	}

	if len(comment.Content) > 1500 { //chars number limit exceeded
		fmt.Printf("⚠ ERROR ⚠ : Couldn't create comment from user %s due to invalid input ❌\n", user)
		return true,
			models.Errormessage{
				Type:       models.BRtype,
				Msg:        "Couldn't create comment due to invalid input",
				StatusCode: models.BRstatus,
				Location:   "home",
				Display:    false,
			}
	}

	errIns := c.InsertComment(comment)
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

func (c *CommentRepository) InsertComment(comment models.Comment) error {
	//generating commentID, date and time
	id_comment, errp := uuid.NewV4() //id
	id_image, errimg := uuid.NewV4() //id
	if errp != nil {
		fmt.Println("❌ Create_comment ⚠ ERROR ⚠ : couldn't generate a unique comment id")
		return errp
	}
	if errimg != nil {
		fmt.Println("❌ Create_comment ⚠ ERROR ⚠ : couldn't generate a unique image id")
		return errimg
	}
	date, time := tools.Time() //date and time

	// inserting value in database
	if comment.Content != "" {
		comment.Content = utils.EncodeValue(comment.Content)
	}
	// processing image
	if comment.Media {
		imgData, err := base64.StdEncoding.DecodeString(comment.Media)
		if err != nil {
			log.Println("❌ error while decoding image", err)
			return
		}

		err = ioutil.WriteFile(id_image, imgData) // storing in local
		if err != nil {
			log.Println("❌ error while storing image in local:", err)
			return
		}
		comment.Media = fmt.Sprintf("http://localhost:8000/images/%s", id_image)
		log.Println("✔ image decoded successfully")
	}

	err := p.DB.Insert(p.TableName, comment)
	if err != nil {
		fmt.Println("❌ error while inserting comment")
		return err
	}

	fmt.Printf("✅ comment %s has been added to database successfully\n", id_comment.String())

	return nil
}

func (c *CommentRepository) LoadComment(postID int) ([]models.DataComment, error) {
	var commentTab []models.DataComment
	Condition := fmt.Sprintf("c.post_id = '%v' ", postID)
	rows, err := c.DB.Exec(GetCommentQuery, Condition)
	if err != nil {
		fmt.Println("❌ Error while retrieving comments => ", err)
		return models.DataComment{}, errors.New("error while retrieving comments from the database")
	}
	defer rows.Close()

	//! modify retrieval
	for rows.next() {
		var temp models.DataCommment
		errScan := rows.scan(&temp.ID, &temp.Content, &temp.Media, &temp.Date, &temp.Post_id, &temp.Avatar, &temp.UserName, &temp.FullName)
		if errScan != nil {
			fmt.Println("⚠ GetComment scan err ⚠ :", errScan)
			return models.DataComment{}, errors.New("error while scanning")
		}

		temp.Content = utils.DecodeValue(temp.Content)
		commentTab = append(commentTab, temp)
	}
	return commentTab, nil
}

func (c *CommentRepository) GetOneComment(commentID int) (models.DataComment, error) {
	Condition := fmt.Sprintf("c.id = '%v' ", commentID)
	rows, err := c.DB.Exec(GetCommentQuery, Condition)
	if err != nil {
		fmt.Println("❌ Error while retrieving comments => ", err)
		return models.DataComment{}, errors.New("error while retrieving comments from the database")
	}
	defer rows.Close()

	var data models.DataCommment
	//! modify retrieval
	for rows.next() {
		errScan := rows.scan(&data.ID, &data.Content, &data.Media, &data.Date, &data.Post_id, &data.Avatar, &data.UserName, &data.FullName)
		if errScan != nil {
			fmt.Println("⚠ GetOneComment scan err ⚠ :", errScan)
			return models.DataComment{}, errors.New("error while scanning")
		}

		data.Content = utils.DecodeValue(data.Content)
	}
	return data, nil
}
