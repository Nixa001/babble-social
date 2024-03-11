package seed

import (
	"database/sql"
	"log"
)

func InsertData(db *sql.DB) {
	_, err := db.Exec(`
	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ('Madike', 'Yade', 'dickss', 'Male', 'dickss@gmail.com', '1234', 'private', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ('IBG', 'Gueye', 'ibg', 'Male', 'ibg@gmail.com', '1234', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Vincent', 'Ndour','vindour', 'Male', 'vindour@gmail.com', '1234', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');


	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Nikola', 'Faye', 'nixa','Male', 'nixa@gmail.com', '1234', 'private', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Daniella', 'Gueye', 'daniella', 'Female', 'dani@gmail.com', '1234', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');
	`)
	if err != nil {
		log.Fatal("Insert into users", err.Error())

	}

	//`user followers`

	_, err = db.Exec(`
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 2);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 3);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 4);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (2, 3);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (2, 4);
	`)
	if err != nil {
		log.Fatal("Insert into users_followers", err.Error())
	}

	// insert into groups

	_, err = db.Exec(`
		INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 1', 'first group', 1, 'avatar/group.png', '2024-03-05');
		INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 2', 'second group', 3, 'avatar/group.png', '2024-03-05');
				`)
	if err != nil {
		log.Fatal("Insert into groups", err.Error())
	}

	//post user

	_, err = db.Exec(`
		INSERT INTO posts (post_content, post_media, post_date, user_id, group_id, type)
		VALUES ('This is the content of the first post.', '/media/first_post.jpg', '2024-03-05', 1, 1, 'private');
		INSERT INTO posts (post_content, post_media, post_date, user_id, group_id, type)
		VALUES ('This is the content of the second post.', '/media/second_post.jpg', '2024-03-05', 2, 2, 'public');
		INSERT INTO posts (post_content, post_media, post_date, user_id, group_id, type)
		VALUES ('This is the content of the third post.', '/media/third_post.jpg', '2024-03-05', 2, 1, 'private');
	`)
	if err != nil {
		log.Fatal("Insert into posts", err.Error())
	}

	//insert into belongs
	// _, err = db.Exec(`
	//     INSERT INTO postCategory (post_id, category_id)
	//     VALUES (1, 1);
	//     INSERT INTO postCategory (post_id, category_id)
	//     VALUES (1, 2);
	//     INSERT INTO postCategory (post_id, category_id)
	//     VALUES (2, 1);
	//     INSERT INTO postCategory (post_id, category_id)
	//     VALUES (2, 2);
	// `)
	// if err != nil {
	// 	log.Fatal("Insert into postCategory", err.Error())
	// }

	// insert into postReact
	_, err = db.Exec(`
        INSERT INTO postReact (post_id, user_id, is_like)
        VALUES (1, 1, true);
        INSERT INTO postReact (post_id, user_id, is_like)
        VALUES (1, 2, true);
        INSERT INTO postReact (post_id, user_id, is_like)
        VALUES (2, 1, true);
        INSERT INTO postReact (post_id, user_id, is_like)
        VALUES (2, 2, true);
    `)
	if err != nil {
		log.Fatal("Insert into postReact", err.Error())
	}

	//insert into comment
	_, err = db.Exec(`
        INSERT INTO comment (post_id, user_id, content_comment)
        VALUES (1, 1, 'This is the first comment');
        INSERT INTO comment (post_id, user_id, content_comment)
        VALUES (1, 2, 'This is the second comment');
        INSERT INTO comment (post_id, user_id, content_comment)
        VALUES (2, 1, 'This is the third comment');
        INSERT INTO comment (post_id, user_id, content_comment)
        VALUES (2, 2, 'This is the fourth comment');
    `)
	if err != nil {
		log.Fatal("Insert into comment", err.Error())
	}

	// insert into commentReact
	_, err = db.Exec(`
        INSERT INTO commentReact (comment_id, user_id, is_like)
        VALUES (1, 1, true);
        INSERT INTO commentReact (comment_id, user_id, is_like)
        VALUES (1, 2, true);
        INSERT INTO commentReact (comment_id, user_id, is_like)
        VALUES (2, 1, true);
        INSERT INTO commentReact (comment_id, user_id, is_like)
        VALUES (2, 2, true);
    `)
	if err != nil {
		log.Fatal("Insert into commentReact", err.Error())
	}

	// insert into session

	_, err = db.Exec(`
        INSERT INTO sessions (token, user_id, expiration)
        VALUES ('NULL', 'NULL', 'NULL')  
		`)
	if err != nil {
		log.Fatal("Insert into sessions", err.Error())
	}

	// insert into message

	_, err = db.Exec(`
            INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the first message', '2024-03-05');
            INSERT INTO messages (user_id_sender, group_id_receiver,  message_content, date)
            VALUES (2, 1, 'This is the second message', '2024-03-06');
			INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the third message', '2024-03-06');
			`)

	if err != nil {
		log.Fatal("Insert into messages", err.Error())
	}

	// insert into group_followers

	_, err = db.Exec(`
            INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 1);
            INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 2);
			`)
	if err != nil {
		log.Fatal("Insert into group_followers", err.Error())
	}

	// insert into notifications

	_, err = db.Exec(
		`INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver)
		VALUES ("message", false, 1, 3);
		INSERT INTO notifications (notification_type, status,  user_id_sender, id_group)
		VALUES ("message", false, 1, 2 );
		INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver)
		VALUES ("message", true, 1, 3);
		`)
	if err != nil {
		log.Fatal("Insert into notifications", err.Error())
	}

	// insert into confirm

	_, err = db.Exec(
		`INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (1, 2, NULL);

		INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (2, NULL, 1);
        `)
	if err != nil {
		log.Fatal("Insert into confirm", err.Error())
	}

}
