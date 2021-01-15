package tests

import (
	"medium-clone/services/posts"
	"testing"
)

func TestCreatePost(t *testing.T) {
	posts.CreatePost("Hello World", "This is the body", "Software", []string{"Go", "Software"})
}
