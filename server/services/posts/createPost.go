package posts

import (
	"fmt"
	"log"
	"medium-clone/database"
)

type Post struct {
	Title    string
	Body     string
	Category string
	Tags     []string
}

func CreatePost(title, body, category string, tags []string) error {

	post := &Post{title, body, category, tags}
	fmt.Printf("%T", post)
	fmt.Println(post)
	_, err := database.DB().Model(post).Insert()

	if err != nil {
		msg := fmt.Sprintf("Post creation failed ERROR: %s", err.Error())
		log.Fatal(msg)
		return err
	}

	return nil
}
