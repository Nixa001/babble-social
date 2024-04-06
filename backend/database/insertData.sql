-- database: social_networki.db
INSERT INTO users ("first_name", "last_name", "user_name", "gender", "email", "password", "user_type", "birth_date", "avatar", "about_me")
	VALUES ('Madike', 'Yade', 'dickss', 'Male', 'dickss@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'private', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users ("first_name", "last_name", "user_name", "gender", "email", "password", "user_type", "birth_date", "avatar", "about_me")
	VALUES ('IBG', 'Gueye', 'ibg', 'Male', 'ibg@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users ("first_name", "last_name", "user_name", "gender", "email", "password", "user_type", "birth_date", "avatar", "about_me")
	VALUES ( 'Vincent', 'Ndour','vindour', 'Male', 'vindour@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');


	INSERT INTO users ("first_name", "last_name", "user_name", "gender", "email", "password", "user_type", "birth_date", "avatar", "about_me")
	VALUES ( 'Nikola', 'Faye', 'nixa','Male', 'nixa@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'private', '2000-01-01', '/avatars/john.jpg', 'about me...');

	INSERT INTO users ("first_name", "last_name", "user_name", "gender", "email", "password", "user_type", "birth_date", "avatar", "about_me")
	VALUES ( 'Daniella', 'Gueye', 'daniella', 'Female', 'dani@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', '/avatars/john.jpg', 'about me...');
	---------------------------------------------------------------------------
	---------------------------------------------------------------------------
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

        INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 1', 'first group', 1, 'avatar/group.png', '2024-03-05');
		INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 2', 'second group', 3, 'avatar/group.png', '2024-03-05');
		----------------------------------------
        INSERT INTO posts (content, media, date, user_id, group_id, privacy)
		VALUES ('This is the content of the first post.', '/media/first_post.jpg', '2024-03-05', 1, 1, 'private');
		INSERT INTO posts (content, media, date, user_id, group_id, privacy)
		VALUES ('This is the content of the second post.', '/media/second_post.jpg', '2024-03-05',1 , 1, 'public');
		INSERT INTO posts (content, media, date, user_id, group_id, privacy)
		VALUES ('This is the content of the third post.', '/media/third_post.jpg', '2024-03-05', 1, 1, 'almost');
	---------------------------------------------------
    INSERT INTO categories (post_id, category)
		VALUES (1, "other" );
		INSERT INTO categories (post_id, category)
		VALUES (1, "sport" );
		INSERT INTO categories (post_id, category)
		VALUES (2, "music" );
		INSERT INTO categories (post_id, category)
		VALUES (2, "technologie" );
		INSERT INTO categories (post_id, category)
		VALUES (3, "other" );
		INSERT INTO categories (post_id, category)
		VALUES (3, "news" );


    INSERT INTO viewers (post_id, user_id)
		VALUES (1, 1);

		INSERT INTO viewers (post_id, user_id)
		VALUES (1, 2);

		INSERT INTO viewers (post_id, user_id)
		VALUES (2, 3);

		INSERT INTO viewers (post_id, user_id)
		VALUES (3, 4);

		INSERT INTO viewers (post_id, user_id)
		VALUES (3, 5);

          INSERT INTO comment (post_id, user_id, media, content, date)
        VALUES (1, 1, '/media/first_post.jpg', 'This is the first comment', '2024-03-05');
          INSERT INTO comment (post_id, user_id, media, content, date)
        VALUES (1, 1, '/media/first_post.jpg', 'This is the first-sec comment', '2024-03-06');
        INSERT INTO comment (post_id, user_id, media, content, date)
        VALUES (1, 2,'/media/first_post.jpg', 'This is the second comment', '2024-03-05');
        INSERT INTO comment (post_id, user_id, media, content, date)
        VALUES (2, 1, '/media/first_post.jpg','This is the third comment', '2024-03-05');
        INSERT INTO comment (post_id, user_id, media, content, date)
        VALUES (2, 2,'/media/first_post.jpg', 'This is the fourth comment', '2024-03-05');

         INSERT INTO commentReacts (comment_id, post_id, user_id, reaction)
        VALUES (1, 1, 1, true);
        INSERT INTO commentReacts (comment_id, post_id, user_id, reaction)
        VALUES (1, 2,1, true);
        INSERT INTO commentReacts (comment_id, post_id, user_id, reaction)
        VALUES (2, 1,1, true);
        INSERT INTO commentReacts (comment_id, post_id, user_id, reaction)
        VALUES (2, 2,1, true);


          INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the first message', '2024-03-05');
            INSERT INTO messages (user_id_sender, group_id_receiver,  message_content, date)
            VALUES (2, 1, 'This is the second message', '2024-03-06');
			INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the third message', '2024-03-06');


              INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 1);
            INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 2);


            INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver)
		VALUES ("message", false, 1, 3);
		INSERT INTO notifications (notification_type, status,  user_id_sender, id_group)
		VALUES ("message", false, 1, 2 );
		INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver)
		VALUES ("message", true, 1, 3);

        INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (1, 2, NULL);

		INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (2, NULL, 1);
