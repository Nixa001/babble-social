package main

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {

	// creation de la table user
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			user_name TEXT NOT NULL,
			gender TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			user_type VARCHAR(25) NOT NULL,
			birth_date VARCHAR(12),
			avatar VARCHAR(256),
			about_me TEXT
			)
			 `)
	if err != nil {
		log.Fatal("User table", err.Error())
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users_followers (
		user_id_followed INTEGER NOT NULL,
		user_id_follower INTEGER NOT NULL,
        FOREIGN KEY(user_id_followed) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY(user_id_follower) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
			)
			 `)
	if err != nil {
		log.Fatal("Users_followers table", err.Error())
	}

	// Creation de la table catégorie
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			category_name TEXT NOT NULL
		)
		`)

	if err != nil {
		log.Fatal("Categories table", err.Error())
	}

	var idCategory string
	err = db.QueryRow("SELECT id FROM categories WHERE categoryName = 'other'").Scan(&idCategory)
	if err != nil {
		_, err = db.Exec(`
			INSERT INTO categories (categoryName) VALUES ('technologie');
			INSERT INTO categories (categoryName) VALUES ('sport');
			INSERT INTO categories (categoryName) VALUES ('other');
			INSERT INTO categories (categoryName) VALUES ('musique');
			INSERT INTO categories (categoryName) VALUES ('sante');
			INSERT INTO categories (categoryName) VALUES ('news');
			`)
	}

	// Inserer donnee de la table catégorie

	if err != nil {
		log.Fatal("Categories insert", err.Error())
	}

	// Creation de la table users
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			post_title TEXT NOT NULL,
			post_content TEXT NOT NULL,
			post_media TEXT NOT NULL,
			post_date TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			group_id INTEGER NOT NULL,
			type TEXT NOT NULL NOT NULL,
			FOREIGN KEY("user_id")
			REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
		 )
		`)
	if err != nil {
		log.Fatal("Posts table", err.Error())
	}

	// Création tavle belong
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS postCategory (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			category_id INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE ON UPDATE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("PostCategorie table", err.Error())
	}

	// Création de la table likes-post

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS likePost (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			post_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			is_like INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES "posts"("id") ON DELETE CASCADE,
			FOREIGN KEY("user_i") REFERENCES "users"("id") ON DELETE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("LikePost table", err.Error())
	}

	// Créate de la table commente
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comment (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			content_comment TEXT NOT NULL,
			post_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("Comment table", err.Error())
	}

	// Creation de le table likes_comment
	_, err = db.Exec(
		`
		CREATE TABLE IF NOT EXISTS likesComments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			comment_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			is_like INTEGER NOT NULL,
			FOREIGN KEY("comment_id") REFERENCES "comment"("id") ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
		)
	`)
	if err != nil {
		log.Fatal("LikesComent table", err.Error())
	}
	// Creation de la table session
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS sessions (
    	id TEXT PRIMARY KEY NOT NULL,
    	expiration DATETIME NOT NULL,
    	data VARCHAR(255) NOT NULL,
		)
		`)
	if err != nil {
		log.Fatal("Sessions table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS messages (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			message_content TEXT NOT NULL,
			user_id_sender INTEGER NOT NULL,
			user_id_receiver INTERGER DEFAULT "NULL",
			group_id_receiver INTERGER DEFAULT "NULL",
			date TEXT,
			FOREIGN KEY (user_id_sender) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY (user_id_receiver) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Message table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS groups (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			name varchar(256)NOT NULL,
			description text NOT NULL,
			id_user_create integer NOT NULL,
			avatar VARCHAR(256),
			creation_date VARCHAR(256) NOT NULL,
			FOREIGN KEY(id_user_create) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Groups table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS group_followers (
			user_id INTEGER NOT NULL,
			group_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE
			)`)
	if err != nil {
		log.Fatal("Notifications table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS notifications (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			notification_type VARCHAR(25) NOT NULL,
			status BOOL NOT NULL,
			user_id_receiver INTEGER NOT NULL,
			user_id_sender INTEGER NOT NULL
			FOREIGN KEY(user_id_sender) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(user_id_receiver) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Notifications table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS confirm (
			user_id_asker INTEGER NOT NULL,
			user_id_asked INTEGER DEFAULT "NULL",
			group_id INTEGER DEFAULT "NULL"
            FOREIGN KEY(user_id_asker) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(user_id_asked) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE

		)`)
	if err != nil {
		log.Fatal("Confirm table", err.Error())
	}

}
