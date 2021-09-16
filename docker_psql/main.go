package main

import (
	"fmt"
	"os"

	"nerdcademy.dev/psql/model"
)



func main() {

	initErr := model.Init()

	if initErr != nil {
		fmt.Printf("%v\n", initErr.Error())
		os.Exit(1)
	}

	defer model.Close()

	post, err := model.GetPost(1)

	if err != nil {
		fmt.Printf("Error getting post: %v\n", err.Error())
	} else {
		fmt.Printf("Post ID: %d\nTitle: %s\nContent: %s\n", post.ID, post.Title, post.Content)
	}
}