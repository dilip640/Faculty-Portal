package storage

import log "github.com/sirupsen/logrus"

// InsertPost add new post
func InsertPost(postName string) error {
	sqlStatement := `INSERT INTO post (post_name) VALUES ($1);`
	_, err := db.Exec(sqlStatement, postName)
	return err
}

// DeleteDepartment remove post
func DeletePost(postID string) error {
	sqlStatement := `DELETE FROM post WHERE id = $1;`
	_, err := db.Exec(sqlStatement, postID)
	return err
}

// GetPost returns one post
func GetPost(postID string) (Post, error) {
	post := Post{}
	sqlStatement := `SELECT id, post_name FROM post
						WHERE id = $1`
	err := db.QueryRow(sqlStatement, postID).Scan(
		&post.PostID, &post.Name)

	return post, err
}

// GetAllPosts returns all depts
func GetAllPosts() ([]*Post, error) {
	posts := make([]*Post, 0)

	rows, err := db.Query(
		`SELECT id, post_name FROM post`)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			postID   string
			postName string
		)

		if err := rows.Scan(&postID, &postName); err == nil {
			posts = append(posts, &Post{postID, postName})
		} else {
			log.Error(err)
		}
	}

	return posts, nil
}

// Post struct
type Post struct {
	PostID string
	Name   string
}
